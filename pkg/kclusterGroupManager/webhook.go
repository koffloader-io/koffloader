// Copyright 2024 Authors of koffloader-io
// SPDX-License-Identifier: Apache-2.0

package kclusterGroupManager

import (
	"context"

	"go.uber.org/zap"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	crd "github.com/koffloader-io/koffloader/pkg/k8s/apis/koffloader.koffloader.io/v1beta1"
)

// --------------------

type kClusterGroupWebHook struct {
	logger *zap.Logger
	client client.Client
}

var _ webhook.CustomValidator = (*kClusterGroupWebHook)(nil)

// mutating webhook
func (kcgw *kClusterGroupWebHook) Default(ctx context.Context, obj runtime.Object) error {
	logger := kcgw.logger.Named("mutating wehbook")

	r, ok := obj.(*crd.KClusterGroup)
	if !ok {
		err := "failed to get obj"
		logger.Error(err)
		return apierrors.NewBadRequest(err)
	}
	logger.Sugar().Infof("obj: %+v", r)

	return nil

}

func (kcgw *kClusterGroupWebHook) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	logger := kcgw.logger.Named("validating create webhook")

	r, ok := obj.(*crd.KClusterGroup)
	if !ok {
		err := "failed to get obj"
		logger.Error(err)
		return admission.Warnings{}, apierrors.NewBadRequest(err)
	}
	logger.Sugar().Infof("obj: %+v", r)

	return admission.Warnings{}, nil
}

func (kcgw *kClusterGroupWebHook) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) (admission.Warnings, error) {
	logger := kcgw.logger.Named("validating update webhook")

	old, ok := oldObj.(*crd.KClusterGroup)
	if !ok {
		err := "failed to get oldObj"
		logger.Error(err)
		return admission.Warnings{}, apierrors.NewBadRequest(err)
	}
	newkcg, ok := newObj.(*crd.KClusterGroup)
	if !ok {
		err := "failed to get newObj"
		logger.Error(err)
		return admission.Warnings{}, apierrors.NewBadRequest(err)
	}
	logger.Sugar().Infof("oldObj: %+v", old)
	logger.Sugar().Infof("newObj: %+v", newkcg)

	return admission.Warnings{}, nil
}

// ValidateDelete implements webhook.CustomValidator so a webhook will be registered for the type
func (kcgw *kClusterGroupWebHook) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	logger := kcgw.logger.Named("validating delete webhook")

	r, ok := obj.(*crd.KClusterGroup)
	if !ok {
		err := "failed to get obj"
		logger.Error(err)
		return admission.Warnings{}, apierrors.NewBadRequest(err)
	}
	logger.Sugar().Infof("obj: %+v", r)

	return admission.Warnings{}, nil
}

// --------------------

func (kcgw *kClusterGroupWebHook) SetupWebhookWithManager(mgr ctrl.Manager) {
	// the mutating route path : "/mutate-" + strings.ReplaceAll(gvk.Group, ".", "-") + "-" + gvk.Version + "-" + strings.ToLower(gvk.Kind)
	// the validate route path : "/validate-" + strings.ReplaceAll(gvk.Group, ".", "-") + "-" + gvk.Version + "-" + strings.ToLower(gvk.Kind)
	if e := ctrl.NewWebhookManagedBy(mgr).
		For(&crd.KClusterGroup{}).
		WithDefaulter(kcgw).
		WithValidator(kcgw).
		Complete(); e != nil {
		kcgw.logger.Sugar().Fatalf("failed to NewWebhookManagedBy, reason=%v", e)
	}
}
