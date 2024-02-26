// Copyright 2024 Authors of koffloader-io
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"context"
	"github.com/go-logr/logr"
	"github.com/koffloader-io/koffloader/pkg/ciliumManager"
	"github.com/koffloader-io/koffloader/pkg/multiClusterManager"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	"os"
	"os/signal"
	"path/filepath"
	"sigs.k8s.io/controller-runtime/pkg/metrics/server"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"syscall"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.uber.org/zap"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	controllerruntimelog "sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/koffloader-io/koffloader/pkg/debug"
	koffloaderv1beta1 "github.com/koffloader-io/koffloader/pkg/k8s/apis/koffloader.koffloader.io/v1beta1"
	"github.com/koffloader-io/koffloader/pkg/kclusterGroupManager"
	"github.com/koffloader-io/koffloader/pkg/kclusterManager"
	"github.com/koffloader-io/koffloader/pkg/serviceExportPolicyManager"
	"github.com/koffloader-io/koffloader/pkg/types"
)

var scheme = runtime.NewScheme()

func init() {
	utilruntime.Must(koffloaderv1beta1.AddToScheme(scheme))
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
	// init ~/.kube/config
	cfg, _ := clientcmd.Write(*clientcmdapi.NewConfig())
	err := os.Mkdir(types.KubeConfigLocalDir, 0750)
	if err != nil {
		rootLogger.Sugar().Errorf("falied create kubeconfig dir,reason=%v", err)
	}
	f, err := os.OpenFile(types.KubeConfigLocalPath, os.O_CREATE|os.O_WRONLY, 0660)
	if err != nil {
		rootLogger.Sugar().Errorf("falied create kubeconfig,reason=%v", err)
	}
	defer f.Close()
	_, _ = f.Write(cfg)
}

func SetupUtility() {

	// run gops
	d := debug.New(rootLogger)
	if types.ControllerConfig.GopsPort != 0 {
		d.RunGops(int(types.ControllerConfig.GopsPort))
	}

	if types.ControllerConfig.PyroscopeServerAddress != "" {
		d.RunPyroscope(types.ControllerConfig.PyroscopeServerAddress, types.ControllerConfig.PodName)
	}
}

func DaemonMain() {
	logger := rootLogger.Named("koffloader-controller")
	logger.Sugar().Infof("config: %+v", types.ControllerConfig)
	SetupUtility()

	SetupHttpServer()

	// ------

	RunMetricsServer(types.ControllerConfig.PodName)
	MetricGaugeEndpoint.Add(context.Background(), 100)
	MetricGaugeEndpoint.Add(context.Background(), -10)
	MetricGaugeEndpoint.Add(context.Background(), 5)

	attrs := attribute.NewSet(attribute.String("pod1", "value1"), attribute.Int("version", 1))
	MetricCounterRequest.Add(context.Background(), 10, metric.WithAttributeSet(attrs))

	attrs = attribute.NewSet(attribute.String("pod2", "value1"), attribute.Int("version", 1))
	MetricCounterRequest.Add(context.Background(), 5, metric.WithAttributeSet(attrs))

	MetricHistogramDuration.Record(context.Background(), 10)
	MetricHistogramDuration.Record(context.Background(), 20)

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:                 scheme,
		HealthProbeBindAddress: "0",
		Metrics: server.Options{
			CertDir:     filepath.Dir(types.ControllerConfig.TlsServerCertPath),
			BindAddress: "0",
		},

		// lease
		LeaderElection:          true,
		LeaderElectionNamespace: types.ControllerConfig.PodNamespace,
		LeaderElectionID:        types.ControllerElectorLockName,

		// webhook port
		WebhookServer: webhook.NewServer(webhook.Options{
			Port:    int(types.ControllerConfig.WebhookPort),
			CertDir: filepath.Dir(types.ControllerConfig.TlsServerCertPath),
		}),
		Client: client.Options{Cache: &client.CacheOptions{
			DisableFor: []client.Object{
				&corev1.Node{},
				&corev1.Namespace{},
				&corev1.Pod{},
				&corev1.Service{},
				&appsv1.Deployment{},
				&appsv1.StatefulSet{},
				&appsv1.ReplicaSet{},
				&appsv1.DaemonSet{},
			},
		}},
	})
	if err != nil {
		logger.Sugar().Fatalf("failed to NewManager, reason=%v", err)
	}

	// set logger for controller-runtime framework
	// The controller-runtime would print debug stack if we do not init the log previously: https://github.com/kubernetes-sigs/controller-runtime/pull/2357
	ctrl.SetLogger(logr.New(controllerruntimelog.NullLogSink{}))
	ciliumManager.InitCiliumManager(logger.Named("CiliumManager"))
	multiClusterManager.InitMultiClusterManager(mgr.GetClient(), logger.Named("MultilClusterManager"))

	// run koffloader controller
	kclusterManager.RunKClusterController(logger, mgr)
	serviceExportPolicyManager.RunServiceExportPolicyController(logger, mgr)
	kclusterGroupManager.RunKClusterGroupController(logger, mgr)

	go func() {
		logger.Info("Starting koffloader-controller runtime manager")
		if err := mgr.Start(context.Background()); err != nil {
			logger.Fatal(err.Error())
		}
	}()

	// ------------
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)
	WatchSignal(logger, sigCh)
}

func WatchSignal(logger *zap.Logger, sigCh chan os.Signal) {
	for sig := range sigCh {
		logger.Sugar().Warnw("received shutdown", "signal", sig)
		// others...

	}
}
