// Copyright 2024 Authors of koffloader-io
// SPDX-License-Identifier: Apache-2.0

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/koffloader-io/koffloader/pkg/k8s/apis/koffloader.koffloader.io/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// KclusterLister helps list Kclusters.
// All objects returned here must be treated as read-only.
type KclusterLister interface {
	// List lists all Kclusters in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Kcluster, err error)
	// Get retrieves the Kcluster from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.Kcluster, error)
	KclusterListerExpansion
}

// kclusterLister implements the KclusterLister interface.
type kclusterLister struct {
	indexer cache.Indexer
}

// NewKclusterLister returns a new KclusterLister.
func NewKclusterLister(indexer cache.Indexer) KclusterLister {
	return &kclusterLister{indexer: indexer}
}

// List lists all Kclusters in the indexer.
func (s *kclusterLister) List(selector labels.Selector) (ret []*v1.Kcluster, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Kcluster))
	})
	return ret, err
}

// Get retrieves the Kcluster from the index for a given name.
func (s *kclusterLister) Get(name string) (*v1.Kcluster, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("kcluster"), name)
	}
	return obj.(*v1.Kcluster), nil
}