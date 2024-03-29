// Copyright 2024 Authors of koffloader-io
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/koffloader-io/koffloader/pkg/k8s/apis/koffloader.koffloader.io/v1beta1"
	scheme "github.com/koffloader-io/koffloader/pkg/k8s/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// KClustersGetter has a method to return a KClusterInterface.
// A group's client should implement this interface.
type KClustersGetter interface {
	KClusters() KClusterInterface
}

// KClusterInterface has methods to work with KCluster resources.
type KClusterInterface interface {
	Create(ctx context.Context, kCluster *v1beta1.KCluster, opts v1.CreateOptions) (*v1beta1.KCluster, error)
	Update(ctx context.Context, kCluster *v1beta1.KCluster, opts v1.UpdateOptions) (*v1beta1.KCluster, error)
	UpdateStatus(ctx context.Context, kCluster *v1beta1.KCluster, opts v1.UpdateOptions) (*v1beta1.KCluster, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.KCluster, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.KClusterList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.KCluster, err error)
	KClusterExpansion
}

// kClusters implements KClusterInterface
type kClusters struct {
	client rest.Interface
}

// newKClusters returns a KClusters
func newKClusters(c *KoffloaderV1beta1Client) *kClusters {
	return &kClusters{
		client: c.RESTClient(),
	}
}

// Get takes name of the kCluster, and returns the corresponding kCluster object, and an error if there is any.
func (c *kClusters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.KCluster, err error) {
	result = &v1beta1.KCluster{}
	err = c.client.Get().
		Resource("kclusters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of KClusters that match those selectors.
func (c *kClusters) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.KClusterList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.KClusterList{}
	err = c.client.Get().
		Resource("kclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested kClusters.
func (c *kClusters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("kclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a kCluster and creates it.  Returns the server's representation of the kCluster, and an error, if there is any.
func (c *kClusters) Create(ctx context.Context, kCluster *v1beta1.KCluster, opts v1.CreateOptions) (result *v1beta1.KCluster, err error) {
	result = &v1beta1.KCluster{}
	err = c.client.Post().
		Resource("kclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(kCluster).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a kCluster and updates it. Returns the server's representation of the kCluster, and an error, if there is any.
func (c *kClusters) Update(ctx context.Context, kCluster *v1beta1.KCluster, opts v1.UpdateOptions) (result *v1beta1.KCluster, err error) {
	result = &v1beta1.KCluster{}
	err = c.client.Put().
		Resource("kclusters").
		Name(kCluster.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(kCluster).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *kClusters) UpdateStatus(ctx context.Context, kCluster *v1beta1.KCluster, opts v1.UpdateOptions) (result *v1beta1.KCluster, err error) {
	result = &v1beta1.KCluster{}
	err = c.client.Put().
		Resource("kclusters").
		Name(kCluster.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(kCluster).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the kCluster and deletes it. Returns an error if one occurs.
func (c *kClusters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("kclusters").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *kClusters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("kclusters").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched kCluster.
func (c *kClusters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.KCluster, err error) {
	result = &v1beta1.KCluster{}
	err = c.client.Patch(pt).
		Resource("kclusters").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
