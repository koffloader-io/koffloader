// Copyright 2024 Authors of koffloader-io
// SPDX-License-Identifier: Apache-2.0

package kclusterManager

import (
	"context"
	crd "github.com/koffloader-io/koffloader/pkg/k8s/apis/koffloader.koffloader.io/v1beta1"
	"go.uber.org/zap"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type kClusterReconciler struct {
	logger *zap.Logger
	client client.Client
}

var _ reconcile.Reconciler = (*kClusterReconciler)(nil)

func RunKClusterController(l *zap.Logger, mgr ctrl.Manager) {
	r := &kClusterReconciler{
		logger: l.Named("kClusterReconciler"),
		client: mgr.GetClient(),
	}
	r.SetupWithManager(mgr)

	wh := &kClusterWebHook{
		logger: l.Named("KClusterWebHook"),
	}
	wh.SetupWebhookWithManager(mgr)
}

func (kcr *kClusterReconciler) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	obj := &crd.KCluster{}
	if err := kcr.client.Get(ctx, req.NamespacedName, obj); err != nil {
		kcr.logger.Sugar().Errorf("failed get kcluster,reason=%v", err)
	}
	kcr.logger.Sugar().Infof("get reconcile")
	return reconcile.Result{}, nil
}

func (kcr *kClusterReconciler) SetupWithManager(mgr ctrl.Manager) {
	if err := ctrl.NewControllerManagedBy(mgr).
		For(&crd.KCluster{}).
		Complete(kcr); err != nil {
		kcr.logger.Sugar().Fatalf("failed to builder kcluster reconcile, error=%v", err)
	}
}
