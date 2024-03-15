// Copyright 2024 Authors of koffloader-io
// SPDX-License-Identifier: Apache-2.0

package serviceExportPolicyManager

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

type serviceExportPolicyWebHook struct {
	logger *zap.Logger
}

var _ webhook.CustomValidator = (*serviceExportPolicyWebHook)(nil)

// mutating webhook
func (sepw *serviceExportPolicyWebHook) Default(ctx context.Context, obj runtime.Object) error {
	logger := sepw.logger.Named("mutating wehbook")

	r, ok := obj.(*crd.ServiceExportPolicy)
	if !ok {
		err := "failed to get obj"
		logger.Error(err)
		return apierrors.NewBadRequest(err)
	}
	logger.Sugar().Infof("obj: %+v", r)

	return nil

}

func (sepw *serviceExportPolicyWebHook) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	logger := sepw.logger.Named("validating create webhook")

	r, ok := obj.(*crd.ServiceExportPolicy)
	if !ok {
		err := "failed to get obj"
		logger.Error(err)
		return admission.Warnings{}, apierrors.NewBadRequest(err)
	}
	logger.Sugar().Infof("obj: %+v", r)

	return admission.Warnings{}, nil
}

func (sepw *serviceExportPolicyWebHook) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) (admission.Warnings, error) {
	logger := sepw.logger.Named("validating update webhook")

	old, ok := oldObj.(*crd.ServiceExportPolicy)
	if !ok {
		err := "failed to get oldObj"
		logger.Error(err)
		return admission.Warnings{}, apierrors.NewBadRequest(err)
	}
	newsep, ok := newObj.(*crd.ServiceExportPolicy)
	if !ok {
		err := "failed to get newObj"
		logger.Error(err)
		return admission.Warnings{}, apierrors.NewBadRequest(err)
	}
	logger.Sugar().Infof("oldObj: %+v", old)
	logger.Sugar().Infof("newObj: %+v", newsep)

	return admission.Warnings{}, nil
}

// ValidateDelete implements webhook.CustomValidator so a webhook will be registered for the type
func (sepw *serviceExportPolicyWebHook) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	logger := sepw.logger.Named("validating delete webhook")

	r, ok := obj.(*crd.ServiceExportPolicy)
	if !ok {
		err := "failed to get obj"
		logger.Error(err)
		return admission.Warnings{}, apierrors.NewBadRequest(err)
	}
	logger.Sugar().Infof("obj: %+v", r)

	return admission.Warnings{}, nil
}

// --------------------

func (sepw *serviceExportPolicyWebHook) SetupWebhookWithManager(mgr ctrl.Manager) {
	// the mutating route path : "/mutate-" + strings.ReplaceAll(gvk.Group, ".", "-") + "-" + gvk.Version + "-" + strings.ToLower(gvk.Kind)
	// the validate route path : "/validate-" + strings.ReplaceAll(gvk.Group, ".", "-") + "-" + gvk.Version + "-" + strings.ToLower(gvk.Kind)
	if e := ctrl.NewWebhookManagedBy(mgr).
		For(&crd.ServiceExportPolicy{}).
		WithDefaulter(sepw).
		WithValidator(sepw).
		Complete(); e != nil {
		sepw.logger.Sugar().Fatalf("failed to NewWebhookManagedBy, reason=%v", e)
	}
}
