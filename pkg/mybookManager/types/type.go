// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package types

type MybookManager interface {
	RunWebhookServer(webhookPort int, tlsDir string)
	RunController(leaseName, leaseNameSpace string, leaseId string)
}
