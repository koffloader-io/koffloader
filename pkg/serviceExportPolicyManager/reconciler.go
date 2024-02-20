// Copyright 2024 Authors of koffloader-io
// SPDX-License-Identifier: Apache-2.0

package serviceExportPolicyManager

import (
	"context"
	crd "github.com/koffloader-io/koffloader/pkg/k8s/apis/koffloader.koffloader.io/v1beta1"
	"go.uber.org/zap"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type serviceExportPolicyReconciler struct {
	logger *zap.Logger
	client client.Client
}

var _ reconcile.Reconciler = (*serviceExportPolicyReconciler)(nil)

func RunServiceExportPolicyController(l *zap.Logger, mgr ctrl.Manager) {
	r := &serviceExportPolicyReconciler{
		logger: l.Named("serviceExportPolicyReconciler"),
		client: mgr.GetClient(),
	}
	r.SetupWithManager(mgr)

	wh := &serviceExportPolicyWebHook{
		logger: l.Named("ServiceExportPolicyWebHook"),
	}
	wh.SetupWebhookWithManager(mgr)
}

func (kcr *serviceExportPolicyReconciler) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	obj := &crd.ServiceExportPolicy{}
	if err := kcr.client.Get(ctx, req.NamespacedName, obj); err != nil {
		kcr.logger.Sugar().Errorf("failed get ServiceExportPolicy,reason=%v", err)
	}
	kcr.logger.Sugar().Infof("get reconcile")
	return reconcile.Result{}, nil
}

func (kcr *serviceExportPolicyReconciler) SetupWithManager(mgr ctrl.Manager) {
	if err := ctrl.NewControllerManagedBy(mgr).
		For(&crd.ServiceExportPolicy{}).
		Complete(kcr); err != nil {
		kcr.logger.Sugar().Fatalf("failed to builder ServiceExportPolicy reconcile, error=%v", err)
	}
}
