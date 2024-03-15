// Copyright 2024 Authors of koffloader-io
// SPDX-License-Identifier: Apache-2.0

package serviceExportPolicyManager

import (
	"context"
	crd "github.com/koffloader-io/koffloader/pkg/k8s/apis/koffloader.koffloader.io/v1beta1"
	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/api/errors"
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

func (sepr *serviceExportPolicyReconciler) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	obj := &crd.ServiceExportPolicy{}
	if err := sepr.client.Get(ctx, req.NamespacedName, obj); err != nil {
		sepr.logger.Sugar().Errorf("failed get ServiceExportPolicy,reason=%v", err)
		if errors.IsNotFound(err) {
			return reconcile.Result{}, client.IgnoreNotFound(err)
		}
	}
	sepr.logger.Sugar().Infof("reconcile handle get %v ", obj)
	return reconcile.Result{}, nil
}

func (sepr *serviceExportPolicyReconciler) SetupWithManager(mgr ctrl.Manager) {
	if err := ctrl.NewControllerManagedBy(mgr).
		For(&crd.ServiceExportPolicy{}).
		Complete(sepr); err != nil {
		sepr.logger.Sugar().Fatalf("failed to builder ServiceExportPolicy reconcile, error=%v", err)
	}
}
