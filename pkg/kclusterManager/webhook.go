// Copyright 2024 Authors of koffloader-io
// SPDX-License-Identifier: Apache-2.0

package kclusterManager

import (
	"context"

	"go.uber.org/zap"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"

	crd "github.com/koffloader-io/koffloader/pkg/k8s/apis/koffloader.koffloader.io/v1beta1"
)

// --------------------

type kClusterWebHook struct {
	logger *zap.Logger
}

var _ webhook.CustomValidator = (*kClusterWebHook)(nil)

// mutating webhook
func (kw *kClusterWebHook) Default(ctx context.Context, obj runtime.Object) error {
	logger := kw.logger.Named("mutating wehbook")

	r, ok := obj.(*crd.KCluster)
	if !ok {
		kw := "failed to get obj"
		logger.Error(kw)
		return apierrors.NewBadRequest(kw)
	}
	logger.Sugar().Infof("obj: %+v", r)

	return nil

}

func (kw *kClusterWebHook) ValidateCreate(ctx context.Context, obj runtime.Object) error {
	logger := kw.logger.Named("validating create webhook")

	r, ok := obj.(*crd.KCluster)
	if !ok {
		kw := "failed to get obj"
		logger.Error(kw)
		return apierrors.NewBadRequest(kw)
	}
	logger.Sugar().Infof("obj: %+v", r)

	return nil
}

func (kw *kClusterWebHook) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) error {
	logger := kw.logger.Named("validating update webhook")

	old, ok := oldObj.(*crd.KCluster)
	if !ok {
		kw := "failed to get oldObj"
		logger.Error(kw)
		return apierrors.NewBadRequest(kw)
	}
	new, ok := newObj.(*crd.KCluster)
	if !ok {
		kw := "failed to get newObj"
		logger.Error(kw)
		return apierrors.NewBadRequest(kw)
	}
	logger.Sugar().Infof("oldObj: %+v", old)
	logger.Sugar().Infof("newObj: %+v", new)

	return nil
}

// ValidateDelete implements webhook.CustomValidator so a webhook will be registered for the type
func (kw *kClusterWebHook) ValidateDelete(ctx context.Context, obj runtime.Object) error {
	logger := kw.logger.Named("validating delete webhook")

	r, ok := obj.(*crd.KCluster)
	if !ok {
		kw := "failed to get obj"
		logger.Error(kw)
		return apierrors.NewBadRequest(kw)
	}
	logger.Sugar().Infof("obj: %+v", r)

	return nil
}

// --------------------

func (kw *kClusterWebHook) SetupWebhookWithManager(mgr ctrl.Manager) {
	// the mutating route path : "/mutate-" + strings.ReplaceAll(gvk.Group, ".", "-") + "-" + gvk.Version + "-" + strings.ToLower(gvk.Kind)
	// the validate route path : "/validate-" + strings.ReplaceAll(gvk.Group, ".", "-") + "-" + gvk.Version + "-" + strings.ToLower(gvk.Kind)
	if e := ctrl.NewWebhookManagedBy(mgr).
		For(&crd.KCluster{}).
		WithDefaulter(kw).
		WithValidator(kw).
		Complete(); e != nil {
		kw.logger.Sugar().Fatalf("failed to NewWebhookManagedBy, reason=%v", e)
	}
}
