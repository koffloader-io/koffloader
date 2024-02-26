// Copyright 2024 Authors of koffloader-io
// SPDX-License-Identifier: Apache-2.0

package kclusterGroupManager

import (
	"context"
	"fmt"
	"github.com/koffloader-io/koffloader/pkg/ciliumManager"
	crd "github.com/koffloader-io/koffloader/pkg/k8s/apis/koffloader.koffloader.io/v1beta1"
	"github.com/koffloader-io/koffloader/pkg/multiClusterManager"
	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type kClusterGroupReconciler struct {
	logger        *zap.Logger
	client        client.Client
	multilClient  multiClusterManager.MultiClusterManager
	ciliumManager ciliumManager.CiliumManager
}

var _ reconcile.Reconciler = (*kClusterGroupReconciler)(nil)

func RunKClusterGroupController(l *zap.Logger, mgr ctrl.Manager) {
	r := &kClusterGroupReconciler{
		logger:        l.Named("kClusterReconciler"),
		client:        mgr.GetClient(),
		multilClient:  multiClusterManager.GlobalMultiClusterManager,
		ciliumManager: ciliumManager.GlobalCiliumManager,
	}
	r.SetupWithManager(mgr)

	wh := &kClusterGroupWebHook{
		logger: l.Named("KClusterGroupWebHook"),
		client: mgr.GetClient(),
	}
	wh.SetupWebhookWithManager(mgr)

}

func (kcgr *kClusterGroupReconciler) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	kcg := crd.KClusterGroup{}
	if err := kcgr.client.Get(ctx, req.NamespacedName, &kcg); err != nil {
		kcgr.logger.Sugar().Errorf("failed get kclustergroup,reason=%v", err)
		if errors.IsNotFound(err) {
			return reconcile.Result{}, client.IgnoreNotFound(err)
		}
	}
	kcgr.logger.Sugar().Infof("reconcile handle get %v ", kcg)

	status := kcg.Status.DeepCopy()
	if status.ConnectStatus == crd.KGroupStatusCreat || status.ConnectStatus == crd.KGroupStatusConnect {
		return reconcile.Result{}, nil
	} else if status.ConnectStatus == "" {
		status.ConnectStatus = crd.KGroupStatusCreat
	}
	kcg.Status = *status

	if err := kcgr.client.Status().Update(ctx, &kcg); err != nil {
		e := fmt.Errorf("failed update kclustergroup status,reason=%v", err)
		kcgr.logger.Sugar().Error(e)
		return reconcile.Result{}, e
	}

	listOptions := &client.ListOptions{}
	if kcg.Spec.KClusterSelector != nil {
		selector, err := metav1.LabelSelectorAsSelector(kcg.Spec.KClusterSelector)
		if err != nil {
			err = fmt.Errorf("failed convert v1.labelselector to label.labelselector,reason=%v", err)
			kcgr.logger.Sugar().Error(err)
			return reconcile.Result{}, err
		}
		listOptions.LabelSelector = selector
	}

	listKCluster := &crd.KClusterList{}
	err := kcgr.client.List(ctx, listKCluster, listOptions)
	if err != nil {
		err = fmt.Errorf("failed list kcluster resource,reason=%v", err)
		kcgr.logger.Sugar().Error(err)
		return reconcile.Result{}, err
	}

	if len(listKCluster.Items) == 0 {
		err = fmt.Errorf("list kcluster item length is zero")
		kcgr.logger.Sugar().Error(err)
		return reconcile.Result{}, err
	}

	// connect each cluster
	// 1 2 3
	// 1->2 1->3 2->3
	for i := 0; i < len(listKCluster.Items); i++ {
		status.MatchKCluster = append(status.MatchKCluster, listKCluster.Items[i].Namespace)
		kCluster := listKCluster.Items[i]
		kClusterStatus := kCluster.Status.DeepCopy()

		if kCluster.Status.ConnectStatus != crd.KClusterStatusClustermesh {
			kcgr.logger.Sugar().Errorf("kcluster %s connect status is not clustermesh", kCluster.Name)
			continue
		}

		for j := i + 1; j < len(listKCluster.Items); j++ {
			err := kcgr.ciliumManager.EstablishClusterConnect(
				listKCluster.Items[i].Name,
				listKCluster.Items[j].Name,
				listKCluster.Items[i].Spec.ClusterConnector.Cilium.Namespace,
			)
			if err != nil {
				kcgr.logger.Sugar().Errorf("failed connect cluster %s to cluster %s,reason=%v",
					listKCluster.Items[i].Name, listKCluster.Items[j].Name, err)
				kClusterStatus.ConnectStatus = crd.KClusterStatusEstablishFailed
			} else {
				kClusterStatus.ConnectStatus = crd.KClusterStatusEstablish
			}

			kCluster.Status = *kClusterStatus
			if err := kcgr.client.Status().Update(ctx, &kCluster); err != nil {
				e := fmt.Errorf("failed update kcluster %s status,reason=%v", kCluster.Name, err)
				kcgr.logger.Sugar().Error(e)
			}
		}
	}

	status.ConnectStatus = crd.KGroupStatusConnect
	kcg.Status = *status

	if err := kcgr.client.Status().Update(ctx, &kcg); err != nil {
		e := fmt.Errorf("failed update kclustergroup status,reason=%v", err)
		kcgr.logger.Sugar().Error(e)
		return reconcile.Result{}, e
	}

	return reconcile.Result{}, nil
}

func (kcgr *kClusterGroupReconciler) SetupWithManager(mgr ctrl.Manager) {
	if err := ctrl.NewControllerManagedBy(mgr).
		For(&crd.KClusterGroup{}).
		Complete(kcgr); err != nil {
		kcgr.logger.Sugar().Fatalf("failed to builder KClusterGroup reconcile, error=%v", err)
	}
}
