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

type KClusterSpec struct {
	// +kubebuilder:validation:Required
	ClusterConnector *ClusterConnectorType `json:"clusterConnector,omitempty"`

	// +kubebuilder:validation:Required
	KubeConfig *Kubeconfig `json:"kubeconfig,omitempty"`
}

type ClusterConnectorType struct {
	// +kubebuilder:validation:Type:=string
	Type string `json:"type,omitempty"`

	// +kubebuilder:validation:Type:=string
	CiliumNamespace string `json:"ciliumNamespace,omitempty"`
}

type Kubeconfig struct {
	// +kubebuilder:validation:Type:=string
	SecretName string `json:"secretName,omitempty"`

	// +kubebuilder:validation:Type:=string
	SecretNamespace string `json:"secretNamespace,omitempty"`
}

type KClusterStatus struct {
	// +kubebuilder:validation:Type:=string
	KClusterGroup string `json:"kClusterGroup"`

	// +kubebuilder:validation:Type:=string
	ClusterConnector string `json:"clusterConnector,omitempty"`
}

// scope(Namespaced or Cluster)
// +kubebuilder:resource:categories={koffloader},path="kclusters",singular="kcluster",scope="Cluster",shortName={kc}
// +kubebuilder:printcolumn:JSONPath=".spec.clusterConnector.type",description="clusterConnector",name="clusterConnector",type=string
// +kubebuilder:printcolumn:JSONPath=".status.kClusterGroup",description="kClusterGroup",name="kClusterGroup",type=string
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +genclient
// +genclient:nonNamespaced

type KCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec   KClusterSpec   `json:"spec,omitempty"`
	Status KClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

type KClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []KCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KCluster{}, &KClusterList{})
}
