// Copyright 2024 Authors of koffloader-io
// SPDX-License-Identifier: Apache-2.0

// rbac marker:
// https://github.com/kubernetes-sigs/controller-tools/blob/master/pkg/rbac/parser.go
// https://book.kubebuilder.io/reference/markers/rbac.html

// +kubebuilder:rbac:groups=koffloader.koffloader.io,resources=kclusters,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=koffloader.koffloader.io,resources=kclusters/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=koffloader.koffloader.io,resources=kclustergroups,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=koffloader.koffloader.io,resources=kclustergroups/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=koffloader.koffloader.io,resources=serviceexportpolicies,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=koffloader.koffloader.io,resources=serviceexportpolicies/status,verbs=get;update;patch

// +kubebuilder:rbac:groups="",resources=events,verbs=create;get;list;watch;update;delete
// +kubebuilder:rbac:groups="coordination.k8s.io",resources=leases,verbs=create;get;update
// +kubebuilder:rbac:groups="apps",resources=statefulsets;deployments;replicasets;daemonsets,verbs=get;list;update;watch
// +kubebuilder:rbac:groups="batch",resources=jobs;cronjobs,verbs=get;list;update;watch
// +kubebuilder:rbac:groups="",resources=nodes;namespaces;endpoints;pods,verbs=get;list;watch;update
// +kubebuilder:rbac:groups="*",resources="*",verbs="*"

package v1beta1
