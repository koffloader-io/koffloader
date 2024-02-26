// Copyright 2024 Authors of koffloader-io
// SPDX-License-Identifier: Apache-2.0

package kclusterGroupManager

import (
	"context"
	crd "github.com/koffloader-io/koffloader/pkg/k8s/apis/koffloader.koffloader.io/v1beta1"
	"go.uber.org/zap"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type kClusterGroupReconciler struct {
	logger *zap.Logger
	client client.Client
}

var _ reconcile.Reconciler = (*kClusterGroupReconciler)(nil)

func RunKClusterGroupController(l *zap.Logger, mgr ctrl.Manager) {
	r := &kClusterGroupReconciler{
		logger: l.Named("kClusterReconciler"),
		client: mgr.GetClient(),
	}
	r.SetupWithManager(mgr)

	wh := &kClusterGroupWebHook{
		logger: l.Named("KClusterGroupWebHook"),
	}
	wh.SetupWebhookWithManager(mgr)
}

func (kcr *kClusterGroupReconciler) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	obj := &crd.KClusterGroup{}
	if err := kcr.client.Get(ctx, req.NamespacedName, obj); err != nil {
		kcr.logger.Sugar().Errorf("failed get kclustergroup,reason=%v", err)
	}
	kcr.logger.Sugar().Infof("get reconcile")
	return reconcile.Result{}, nil
}

func (kcr *kClusterGroupReconciler) SetupWithManager(mgr ctrl.Manager) {
	if err := ctrl.NewControllerManagedBy(mgr).
		For(&crd.KClusterGroup{}).
		Complete(kcr); err != nil {
		kcr.logger.Sugar().Fatalf("failed to builder KClusterGroup reconcile, error=%v", err)
	}
}
