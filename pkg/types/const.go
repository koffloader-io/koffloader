// Copyright 2024 Authors of koffloader-io
// SPDX-License-Identifier: Apache-2.0
package types

import (
	"os/user"
	"path"
)

const (
	ControllerElectorLockName = "kdoctor-controller-lease"
	TlsCaCommonName           = "koffloader.io"
)

var (
	u, _                = user.Current()
	KubeConfigLocalPath = path.Join(u.HomeDir, ".kube", "config")
	KubeConfigLocalDir  = path.Join(u.HomeDir, ".kube")
)
