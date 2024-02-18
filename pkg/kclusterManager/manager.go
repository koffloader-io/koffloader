// Copyright 2024 Authors of koffloader-io
// SPDX-License-Identifier: Apache-2.0

package kclusterManager

import (
	"github.com/koffloader-io/koffloader/pkg/kclusterManager/types"
	"go.uber.org/zap"
)

type kclusterManager struct {
	logger   *zap.Logger
	webhook  *webhookhander
	informer *informerHandler
}

var _ types.kclusterManager = (*kclusterManager)(nil)

func New(logger *zap.Logger) types.kclusterManager {

	return &kclusterManager{
		logger: logger.Named("kclusterManager"),
	}
}
