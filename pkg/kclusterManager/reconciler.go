// Copyright 2024 Authors of koffloader-io
// SPDX-License-Identifier: Apache-2.0

package kclusterManager

import (
	"context"
	"fmt"
	"github.com/koffloader-io/koffloader/pkg/ciliumManager"
	crd "github.com/koffloader-io/koffloader/pkg/k8s/apis/koffloader.koffloader.io/v1beta1"
	"github.com/koffloader-io/koffloader/pkg/multiClusterManager"
	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/api/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type kClusterReconciler struct {
	logger        *zap.Logger
	client        client.Client
	multiClient   multiClusterManager.MultiClusterManager
	ciliumManager ciliumManager.CiliumManager
}

var _ reconcile.Reconciler = (*kClusterReconciler)(nil)

func RunKClusterController(l *zap.Logger, mgr ctrl.Manager) {
	r := &kClusterReconciler{
		logger:        l.Named("kClusterReconciler"),
		client:        mgr.GetClient(),
		multiClient:   multiClusterManager.GlobalMultiClusterManager,
		ciliumManager: ciliumManager.GlobalCiliumManager,
	}
	r.SetupWithManager(mgr)

	wh := &kClusterWebHook{
		logger: l.Named("KClusterWebHook"),
	}
	wh.SetupWebhookWithManager(mgr)
}

func (kcr *kClusterReconciler) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	kc := &crd.KCluster{}
	if err := kcr.client.Get(ctx, req.NamespacedName, kc); err != nil {
		kcr.logger.Sugar().Errorf("failed get kcluster,reason=%v", err)
		if errors.IsNotFound(err) {
			return reconcile.Result{}, client.IgnoreNotFound(err)
		}
	}
	kcr.logger.Sugar().Infof("reconcile handle get %v ", kc)

	status := kc.Status.DeepCopy()
	if status.ConnectStatus != "" {
		return reconcile.Result{}, nil
	}
	status.ConnectStatus = crd.KClusterStatusProcess

	kc.Status = *status

	if err := kcr.client.Status().Update(ctx, kc); err != nil {
		e := fmt.Errorf("failed update kcluster status,reason=%v", err)
		kcr.logger.Sugar().Error(e)
		return reconcile.Result{}, e
	}
	err := kcr.multiClient.SaveKubeConfigToLocal(*kc)
	if err != nil {
		kcr.logger.Sugar().Error(err)
		return reconcile.Result{}, err
	}

	err = kcr.ciliumManager.EnableClustermesh(
		kcr.multiClient.GetClusterContextNameByKCluster(*kc),
		kc.Spec.ClusterConnector.Cilium.Namespace,
		kc.Spec.ClusterConnector.Cilium.ClustermeshServiceType)
	if err != nil {
		kcr.logger.Sugar().Error(err)
		status.ConnectStatus = crd.KClusterStatusClustermeshFailed
	} else {
		status.ConnectStatus = crd.KClusterStatusClustermesh
	}

	kc.Status = *status
	if err := kcr.client.Status().Update(ctx, kc); err != nil {
		e := fmt.Errorf("failed update kcluster status,reason=%v", err)
		kcr.logger.Sugar().Error(e)
		return reconcile.Result{}, e
	}

	return reconcile.Result{}, nil
}

func (kcr *kClusterReconciler) SetupWithManager(mgr ctrl.Manager) {
	if err := ctrl.NewControllerManagedBy(mgr).
		For(&crd.KCluster{}).
		Complete(kcr); err != nil {
		kcr.logger.Sugar().Fatalf("failed to builder kcluster reconcile, error=%v", err)
	}
}
