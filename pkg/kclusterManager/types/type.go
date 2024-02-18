// Copyright 2024 Authors of koffloader-io
// SPDX-License-Identifier: Apache-2.0

package types

type kclusterManager interface {
	RunWebhookServer(webhookPort int, tlsDir string)
	RunController(leaseName, leaseNameSpace string, leaseId string)
}
