// Copyright 2024 Authors of koffloader-io
// SPDX-License-Identifier: Apache-2.0

package multiClusterManager

import (
	"context"
	"fmt"
	ciliumv2 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2"
	ciliumv2alpha1 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2alpha1"
	tetragonv1alpha1 "github.com/cilium/tetragon/pkg/k8s/apis/cilium.io/v1alpha1"
	"github.com/koffloader-io/koffloader/pkg/utils"
	"go.uber.org/zap"
	"golang.org/x/exp/maps"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	"os"
	"sigs.k8s.io/controller-runtime/pkg/client"

	crd "github.com/koffloader-io/koffloader/pkg/k8s/apis/koffloader.koffloader.io/v1beta1"
	consts "github.com/koffloader-io/koffloader/pkg/types"
)

var scheme = runtime.NewScheme()

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
	utilruntime.Must(ciliumv2.AddToScheme(scheme))
	utilruntime.Must(ciliumv2alpha1.AddToScheme(scheme))
	utilruntime.Must(tetragonv1alpha1.AddToScheme(scheme))
}

type MultiClusterManager interface {
	GetClusterClientByKCluster(kc crd.KCluster) (client.Client, error)
	DeleteClusterClientByName(kClusterName string)
	SaveKubeConfigToLocal(kc crd.KCluster) (err error)
	GetClusterContextNameByKCluster(kc crd.KCluster) string
}

type multiClusterManager struct {
	MultiClient       map[string]client.Client
	MultiContextNames map[string]string
	client            client.Client
	logger            *zap.Logger
	ctx               context.Context
}

var GlobalMultiClusterManager MultiClusterManager

func InitMultiClusterManager(c client.Client, logger *zap.Logger) {
	if GlobalMultiClusterManager == nil {
		GlobalMultiClusterManager = &multiClusterManager{
			client:            c,
			ctx:               context.Background(),
			logger:            logger,
			MultiClient:       map[string]client.Client{},
			MultiContextNames: map[string]string{},
		}
	}
}

func (mcm *multiClusterManager) newClientConfigFromKubeConfig(kc crd.KCluster) (err error) {
	if kc.Spec.KubeConfig == nil {
		err = fmt.Errorf("in kcluster %s ,kubeconfig field is empty", kc.Name)
		mcm.logger.Sugar().Error(err)
		return
	}

	kcfgSecret := &v1.Secret{}

	if err = mcm.client.Get(mcm.ctx, types.NamespacedName{
		Name:      kc.Spec.KubeConfig.SecretName,
		Namespace: kc.Spec.KubeConfig.SecretNamespace},
		kcfgSecret); err != nil {
		err = fmt.Errorf("failed get kubeconfig secret,reason=%v", err)
		mcm.logger.Sugar().Error(err)
		return
	}

	data, ok := kcfgSecret.Data["kubeconfig"]
	if !ok {
		err = fmt.Errorf("not found kubeconfig field in secret")
		mcm.logger.Sugar().Error(err)
		return
	}

	clientConfig, err := clientcmd.NewClientConfigFromBytes(data)
	if err != nil {
		err = fmt.Errorf("failed generate client config from bytes,reason=%v", err)
		mcm.logger.Sugar().Error(err)
		return
	}
	cfg, err := clientConfig.ClientConfig()
	if err != nil {
		err = fmt.Errorf("failed generate client rest config,reason=%v", err)
		mcm.logger.Sugar().Error(err)
		return err
	}
	c, err := client.New(cfg, client.Options{Scheme: scheme})
	if err != nil {
		err = fmt.Errorf("failed generate client,reason=%v", err)
		mcm.logger.Sugar().Error(err)
		return err
	}

	mcm.MultiClient[kc.Name] = c

	return nil
}

func (mcm *multiClusterManager) GetClusterClientByKCluster(kc crd.KCluster) (client.Client, error) {
	c, ok := mcm.MultiClient[kc.Name]
	if ok {
		return c, nil
	}

	if err := mcm.newClientConfigFromKubeConfig(kc); err != nil {
		return nil, err
	}
	c, _ = mcm.MultiClient[kc.Name]
	return c, nil
}

func (mcm *multiClusterManager) DeleteClusterClientByName(kClusterName string) {
	delete(mcm.MultiClient, kClusterName)
}

func (mcm *multiClusterManager) SaveKubeConfigToLocal(kc crd.KCluster) (err error) {
	kcfgSecret := &v1.Secret{}

	if err = mcm.client.Get(mcm.ctx, types.NamespacedName{
		Name:      kc.Spec.KubeConfig.SecretName,
		Namespace: kc.Spec.KubeConfig.SecretNamespace},
		kcfgSecret); err != nil {
		err = fmt.Errorf("failed get kubeconfig secret,reason=%v", err)
		mcm.logger.Sugar().Error(err)
		return
	}

	data, ok := kcfgSecret.Data["kubeconfig"]
	if !ok {
		err = fmt.Errorf("not found kubeconfig field in secret")
		mcm.logger.Sugar().Error(err)
		return
	}

	singleConfig, err := clientcmd.Load(data)
	if err != nil {
		err = fmt.Errorf("failed generate client config from bytes,reason=%v", err)
		mcm.logger.Sugar().Error(err)
		return
	}

	ctxName := maps.Keys(singleConfig.Contexts)
	mcm.MultiContextNames[kc.Name] = ctxName[0]

	cfg, err := mcm.getLocalKubeConfig()
	if err != nil {
		mcm.logger.Sugar().Errorf("failed save kubeconfig to local,reason=%v", err)
		return err
	}
	cfg.AuthInfos = utils.MergeMap(cfg.AuthInfos, singleConfig.AuthInfos)
	cfg.Contexts = utils.MergeMap(cfg.Contexts, singleConfig.Contexts)
	cfg.Clusters = utils.MergeMap(cfg.Clusters, singleConfig.Clusters)
	cfg.Extensions = utils.MergeMap(cfg.Extensions, singleConfig.Extensions)

	cfgBytes, err := clientcmd.Write(*cfg)
	if err != nil {
		mcm.logger.Sugar().Errorf("failed kubeconfig convert bytes,reason=%v", err)
		return err
	}

	if err := os.WriteFile(consts.KubeConfigLocalPath, cfgBytes, 0666); err != nil {
		mcm.logger.Sugar().Errorf("failed kubeconfig write to local file,reason=%v", err)
		return err
	}
	return nil
}

func (mcm *multiClusterManager) getLocalKubeConfig() (*clientcmdapi.Config, error) {
	data, err := os.ReadFile(consts.KubeConfigLocalPath)
	if err != nil {
		mcm.logger.Sugar().Errorf("failed get local kebeconfig,reason=%v", err)
		return nil, err
	}
	cfg, err := clientcmd.Load(data)
	if err != nil {
		mcm.logger.Sugar().Errorf("failed load kubeconfig from bytes,reason=%v", err)
		return nil, err
	}
	return cfg, nil
}

func (mcm *multiClusterManager) GetClusterContextNameByKCluster(kc crd.KCluster) string {
	return mcm.MultiContextNames[kc.Name]
}
