// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"context"
	"github.com/spidernet-io/rocktemplate/api/v1/grpcService"
	"github.com/spidernet-io/rocktemplate/pkg/debug"
	"github.com/spidernet-io/rocktemplate/pkg/grpcManager"
	"github.com/spidernet-io/rocktemplate/pkg/types"
	"github.com/spidernet-io/rocktemplate/pkg/utils"
	"time"
)

func SetupUtility() {
	// run gops
	d := debug.New(rootLogger)
	if types.AgentConfig.GopsPort != 0 {
		d.RunGops(int(types.AgentConfig.GopsPort))
	}

	if types.AgentConfig.PyroscopeServerAddress != "" {
		d.RunPyroscope(types.AgentConfig.PyroscopeServerAddress, types.AgentConfig.PodName)
	}
}

func testGrpc() {
	// ---- grpc server
	rootLogger.Info("start grpc server")
	alternateDNS := []string{}
	// add pod name
	alternateDNS = append(alternateDNS, types.AgentConfig.PodName)
	tlsCertPath := "/tmp/cert.crt"
	tlsKeyPath := "/tmp/key.crt"
	tlsCaPath := "/tmp/ca.crt"
	// generate self-signed certificates
	if e := utils.NewServerCertKeyForLocalNode(alternateDNS, tlsCertPath, tlsKeyPath, tlsCaPath); e != nil {
		rootLogger.Sugar().Fatalf("failed to generate certiface, error=%v", e)
	}
	t := grpcManager.NewGrpcServer(rootLogger, tlsCertPath, tlsKeyPath)
	// listen on ipv4 and ipv6
	t.Run(":3000")
	time.Sleep(time.Second * 10)

	// self test grpc
	rootLogger.Info("test grpc server")
	client := grpcManager.NewGrpcClient(rootLogger, true)
	p := &grpcService.ExecRequestMsg{
		Command:       "ls /",
		Timeoutsecond: 10,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	if reslist, err := client.SendRequestForExecRequest(ctx, []string{"127.0.0.1:3000"}, []*grpcService.ExecRequestMsg{p}); err != nil {
		rootLogger.Sugar().Errorf("grpc failed to send request, error=%v", err)
	} else {
		rootLogger.Sugar().Infof("grpc response: %+v", reslist[0])
	}
	cancel()

}

func DaemonMain() {
	rootLogger.Sugar().Infof("config: %+v", types.AgentConfig)

	SetupUtility()

	SetupHttpServer()

	RunMetricsServer(types.AgentConfig.PodName)

	testGrpc()

	// -----------
	rootLogger.Info("hello world")
	time.Sleep(time.Hour)
}
