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

const (
	KGroupStatusCreat   = "Creating"
	KGroupStatusConnect = "Connecting"
)

type KClusterGroupSpec struct {
	// +kubebuilder:validation:Required
	ClusterConnectorType *string `json:"clusterConnectorType,omitempty"`

	// +kubebuilder:validation:Required
	KClusterSelector *metav1.LabelSelector `json:"kClusterSelector,omitempty"`
}

type KClusterGroupStatus struct {
	// +kubebuilder:validation:Optional
	MatchKCluster []string `json:"matchKCluster,omitempty"`

	// +kubebuilder:validation:Optional
	ConnectStatus string `json:"connectStatus,omitempty"`

	// +kubebuilder:validation:Optional
	ClusterConnectorType string `json:"clusterConnectorType,omitempty"`
}

// scope(Namespaced or Cluster)
// +kubebuilder:resource:categories={koffloader},path="kclustergroups",singular="kclustergroup",scope="Cluster",shortName={kcg}
// +kubebuilder:printcolumn:JSONPath=".spec.clusterConnectorType",description="clusterConnectorType",name="clusterConnectorType",type=string
// +kubebuilder:printcolumn:JSONPath=".status.connectStatus",description="ConnectStatus",name="ConnectStatus",type=string
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +genclient
// +genclient:nonNamespaced

type KClusterGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec   KClusterGroupSpec   `json:"spec,omitempty"`
	Status KClusterGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

type KClusterGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []KClusterGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KClusterGroup{}, &KClusterGroupList{})
}
