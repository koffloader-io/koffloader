// Copyright 2024 Authors of koffloader-io
// SPDX-License-Identifier: Apache-2.0

package ciliumManager

import (
	"context"
	"github.com/cilium/cilium-cli/clustermesh"
	"github.com/cilium/cilium-cli/k8s"
	"go.uber.org/zap"
	"os"
)

type CiliumManager interface {
	EnableClustermesh(ctxName, namespace, serviceType string) error
	EstablishClusterConnect(ctxNameSrc, ctxNameDest, namespace string) error
}

var GlobalCiliumManager CiliumManager

type ciliumManager struct {
	ctx          context.Context
	logger       *zap.Logger
	ciliumClient map[string]*k8s.Client
}

func InitCiliumManager(logger *zap.Logger) {
	if GlobalCiliumManager == nil {
		GlobalCiliumManager = &ciliumManager{
			ctx:          context.Background(),
			logger:       logger,
			ciliumClient: make(map[string]*k8s.Client),
		}
	}
	return
}

func (cm *ciliumManager) newCiliumClient(contextName string, namespace string) (*k8s.Client, error) {
	c, err := k8s.NewClient(contextName, "", namespace)
	if err != nil {
		cm.logger.Sugar().Errorf("falied new ClustermeshClient,reaseon=%v", err)
		return nil, err
	}
	cm.ciliumClient[contextName] = c
	return c, nil
}

func (cm *ciliumManager) getCiliumClient(ctxName, namespace string) *k8s.Client {
	var c *k8s.Client
	c, ok := cm.ciliumClient[ctxName]
	if !ok {
		c, _ := cm.newCiliumClient(ctxName, namespace)
		return c
	}
	return c
}

func (cm *ciliumManager) EnableClustermesh(ctxName, namespace, serviceType string) error {
	cm.logger.Sugar().Debugf("enable cluster context %s service type %s clustermesh", ctxName, serviceType)
	c := cm.getCiliumClient(ctxName, namespace)
	param := clustermesh.Parameters{}
	param.Namespace = namespace
	param.ServiceType = serviceType
	param.Writer = os.Stdout
	param.CreateCA = true
	if apiserverVersion, err := c.GetRunningCiliumVersion(cm.ctx, namespace); err == nil {
		param.ApiserverVersion = "v" + apiserverVersion
	}

	if err := clustermesh.EnableWithHelm(cm.ctx, c, param); err != nil {
		cm.logger.Sugar().Errorf("failed enable cluster%s ,reason=%v", ctxName, err)
		return err
	}

	return nil
}

func (cm *ciliumManager) EstablishClusterConnect(ctxNameSrc, ctxNameDest, namespace string) error {
	c := cm.getCiliumClient(ctxNameSrc, namespace)
	param := clustermesh.Parameters{}
	param.Writer = os.Stdout
	param.DestinationContext = ctxNameDest

	if err := clustermesh.NewK8sClusterMesh(c, param).Connect(cm.ctx); err != nil {
		cm.logger.Sugar().Errorf("failed establish connect cluster src: %s dest: %v,reason=%v",
			ctxNameSrc, ctxNameDest, err,
		)
		return err
	}
	return nil
}
