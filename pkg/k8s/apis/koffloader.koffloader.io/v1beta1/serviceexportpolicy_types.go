// Copyright 2024 Authors of koffloader-io
// SPDX-License-Identifier: Apache-2.0

// !!!!!! crd marker:
// https://github.com/kubernetes-sigs/controller-tools/blob/master/pkg/crd/markers/crd.go
// https://book.kubebuilder.io/reference/markers/crd.html
// https://github.com/kubernetes-sigs/controller-tools/blob/master/pkg/crd/markers/validation.go
// https://book.kubebuilder.io/reference/markers/crd-validation.html

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ServiceExportPolicySpec struct {
	// +kubebuilder:validation:Required
	CiliumServiceFeature *CiliumServiceFeature `json:"ciliumServiceFeature,omitempty"`

	// +kubebuilder:validation:Required
	ServiceSelector *metav1.LabelSelector `json:"serviceSelector,omitempty"`

	// +kubebuilder:validation:Required
	KClusterSelector *metav1.LabelSelector `json:"kClusterSelector,omitempty"`
}

type CiliumServiceFeature struct {
	// +kubebuilder:validation:Type:=boolean
	Share bool `json:"share,omitempty"`

	// +kubebuilder:validation:Type:=string
	Affinity string `json:"affinity,omitempty"`

	// +kubebuilder:validation:Type:=boolean
	Global bool `json:"global,omitempty"`
}

type ServiceExportPolicyStatus struct {
	// +kubebuilder:validation:Type:=string
	MatchKCluster []string `json:"matchKCluster,omitempty"`

	// +kubebuilder:validation:Optional
	MatchService []ExportService `json:"matchService,omitempty"`
}

type ExportService struct {
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
}

// scope(Namespaced or Cluster)
// +kubebuilder:resource:categories={koffloader},path="serviceexportpolicies",singular="serviceexportpolicy",scope="Cluster",shortName={sep}
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +genclient
// +genclient:nonNamespaced

type ServiceExportPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec   ServiceExportPolicySpec   `json:"spec,omitempty"`
	Status ServiceExportPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

type ServiceExportPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ServiceExportPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ServiceExportPolicy{}, &ServiceExportPolicyList{})
}
