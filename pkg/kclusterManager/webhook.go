// Copyright 2024 Authors of koffloader-io
// SPDX-License-Identifier: Apache-2.0

package kclusterManager

import (
	"context"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

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
		err := "failed to get obj"
		logger.Error(err)
		return apierrors.NewBadRequest(err)
	}
	logger.Sugar().Infof("obj: %+v", r)

	return nil

}

func (kw *kClusterWebHook) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	logger := kw.logger.Named("validating create webhook")

	r, ok := obj.(*crd.KCluster)
	if !ok {
		err := "failed to get obj"
		logger.Error(err)
		return admission.Warnings{}, apierrors.NewBadRequest(err)
	}
	logger.Sugar().Infof("obj: %+v", r)

	return admission.Warnings{}, nil
}

func (kw *kClusterWebHook) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) (admission.Warnings, error) {
	logger := kw.logger.Named("validating update webhook")

	old, ok := oldObj.(*crd.KCluster)
	if !ok {
		err := "failed to get oldObj"
		logger.Error(err)
		return admission.Warnings{}, apierrors.NewBadRequest(err)
	}
	newkc, ok := newObj.(*crd.KCluster)
	if !ok {
		err := "failed to get newObj"
		logger.Error(err)
		return admission.Warnings{}, apierrors.NewBadRequest(err)
	}
	logger.Sugar().Infof("oldObj: %+v", old)
	logger.Sugar().Infof("newObj: %+v", newkc)

	return admission.Warnings{}, nil
}

// ValidateDelete implements webhook.CustomValidator so a webhook will be registered for the type
func (kw *kClusterWebHook) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	logger := kw.logger.Named("validating delete webhook")

	r, ok := obj.(*crd.KCluster)
	if !ok {
		err := "failed to get obj"
		logger.Error(err)
		return admission.Warnings{}, apierrors.NewBadRequest(err)
	}
	logger.Sugar().Infof("obj: %+v", r)

	return admission.Warnings{}, nil
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
