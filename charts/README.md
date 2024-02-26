# koffloader

## Introduction

## Features

## Parameters

### Global parameters

| Name                           | Description                                | Value           |
| ------------------------------ | ------------------------------------------ | --------------- |
| `global.imageRegistryOverride` | Global Docker image registry               | `""`            |
| `global.imageTagOverride`      | Global Docker image tag                    | `""`            |
| `global.name`                  | instance name                              | `koffloader`    |
| `global.clusterDnsDomain`      | cluster dns domain                         | `cluster.local` |
| `global.commonAnnotations`     | Annotations to add to all deployed objects | `{}`            |
| `global.commonLabels`          | Labels to add to all deployed objects      | `{}`            |
| `global.configName`            | the configmap name                         | `koffloader`    |

### feature parameters

| Name                 | Description | Value   |
| -------------------- | ----------- | ------- |
| `feature.enableIPv4` | enable ipv4 | `true`  |
| `feature.enableIPv6` | enable ipv6 | `false` |

### koffloaderAgent parameters

| Name                                                         | Description                                                                                      | Value                            |
| ------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | -------------------------------- |
| `koffloaderAgent.name`                                       | the koffloaderAgent name                                                                         | `koffloader-agent`               |
| `koffloaderAgent.cmdBinName`                                 | the binary name of koffloaderAgent                                                               | `/usr/bin/agent`                 |
| `koffloaderAgent.hostnetwork`                                | enable hostnetwork mode of koffloaderAgent pod                                                   | `true`                           |
| `koffloaderAgent.image.registry`                             | the image registry of koffloaderAgent                                                            | `ghcr.io`                        |
| `koffloaderAgent.image.repository`                           | the image repository of koffloaderAgent                                                          | `koffloader-io/koffloader-agent` |
| `koffloaderAgent.image.pullPolicy`                           | the image pullPolicy of koffloaderAgent                                                          | `IfNotPresent`                   |
| `koffloaderAgent.image.digest`                               | the image digest of koffloaderAgent, which takes preference over tag                             | `""`                             |
| `koffloaderAgent.image.tag`                                  | the image tag of koffloaderAgent, overrides the image tag whose default is the chart appVersion. | `""`                             |
| `koffloaderAgent.image.imagePullSecrets`                     | the image imagePullSecrets of koffloaderAgent                                                    | `[]`                             |
| `koffloaderAgent.serviceAccount.create`                      | create the service account for the koffloaderAgent                                               | `true`                           |
| `koffloaderAgent.serviceAccount.annotations`                 | the annotations of koffloaderAgent service account                                               | `{}`                             |
| `koffloaderAgent.service.annotations`                        | the annotations for koffloaderAgent service                                                      | `{}`                             |
| `koffloaderAgent.service.type`                               | the type for koffloaderAgent service                                                             | `ClusterIP`                      |
| `koffloaderAgent.priorityClassName`                          | the priority Class Name for koffloaderAgent                                                      | `system-node-critical`           |
| `koffloaderAgent.affinity`                                   | the affinity of koffloaderAgent                                                                  | `{}`                             |
| `koffloaderAgent.extraArgs`                                  | the additional arguments of koffloaderAgent container                                            | `[]`                             |
| `koffloaderAgent.extraEnv`                                   | the additional environment variables of koffloaderAgent container                                | `[]`                             |
| `koffloaderAgent.extraVolumes`                               | the additional volumes of koffloaderAgent container                                              | `[]`                             |
| `koffloaderAgent.extraVolumeMounts`                          | the additional hostPath mounts of koffloaderAgent container                                      | `[]`                             |
| `koffloaderAgent.podAnnotations`                             | the additional annotations of koffloaderAgent pod                                                | `{}`                             |
| `koffloaderAgent.podLabels`                                  | the additional label of koffloaderAgent pod                                                      | `{}`                             |
| `koffloaderAgent.resources.limits.cpu`                       | the cpu limit of koffloaderAgent pod                                                             | `1000m`                          |
| `koffloaderAgent.resources.limits.memory`                    | the memory limit of koffloaderAgent pod                                                          | `1024Mi`                         |
| `koffloaderAgent.resources.requests.cpu`                     | the cpu requests of koffloaderAgent pod                                                          | `100m`                           |
| `koffloaderAgent.resources.requests.memory`                  | the memory requests of koffloaderAgent pod                                                       | `128Mi`                          |
| `koffloaderAgent.securityContext`                            | the security Context of koffloaderAgent pod                                                      | `{}`                             |
| `koffloaderAgent.httpServer.port`                            | the http Port for koffloaderAgent, for health checking                                           | `5710`                           |
| `koffloaderAgent.httpServer.startupProbe.failureThreshold`   | the failure threshold of startup probe for koffloaderAgent health checking                       | `60`                             |
| `koffloaderAgent.httpServer.startupProbe.periodSeconds`      | the period seconds of startup probe for koffloaderAgent health checking                          | `2`                              |
| `koffloaderAgent.httpServer.livenessProbe.failureThreshold`  | the failure threshold of startup probe for koffloaderAgent health checking                       | `6`                              |
| `koffloaderAgent.httpServer.livenessProbe.periodSeconds`     | the period seconds of startup probe for koffloaderAgent health checking                          | `10`                             |
| `koffloaderAgent.httpServer.readinessProbe.failureThreshold` | the failure threshold of startup probe for koffloaderAgent health checking                       | `3`                              |
| `koffloaderAgent.httpServer.readinessProbe.periodSeconds`    | the period seconds of startup probe for koffloaderAgent health checking                          | `10`                             |
| `koffloaderAgent.prometheus.enabled`                         | enable template agent to collect metrics                                                         | `false`                          |
| `koffloaderAgent.prometheus.port`                            | the metrics port of template agent                                                               | `5711`                           |
| `koffloaderAgent.prometheus.serviceMonitor.install`          | install serviceMonitor for template agent. This requires the prometheus CRDs to be available     | `false`                          |
| `koffloaderAgent.prometheus.serviceMonitor.namespace`        | the serviceMonitor namespace. Default to the namespace of helm instance                          | `""`                             |
| `koffloaderAgent.prometheus.serviceMonitor.annotations`      | the additional annotations of koffloaderAgent serviceMonitor                                     | `{}`                             |
| `koffloaderAgent.prometheus.serviceMonitor.labels`           | the additional label of koffloaderAgent serviceMonitor                                           | `{}`                             |
| `koffloaderAgent.prometheus.prometheusRule.install`          | install prometheusRule for template agent. This requires the prometheus CRDs to be available     | `false`                          |
| `koffloaderAgent.prometheus.prometheusRule.namespace`        | the prometheusRule namespace. Default to the namespace of helm instance                          | `""`                             |
| `koffloaderAgent.prometheus.prometheusRule.annotations`      | the additional annotations of koffloaderAgent prometheusRule                                     | `{}`                             |
| `koffloaderAgent.prometheus.prometheusRule.labels`           | the additional label of koffloaderAgent prometheusRule                                           | `{}`                             |
| `koffloaderAgent.prometheus.grafanaDashboard.install`        | install grafanaDashboard for template agent. This requires the prometheus CRDs to be available   | `false`                          |
| `koffloaderAgent.prometheus.grafanaDashboard.namespace`      | the grafanaDashboard namespace. Default to the namespace of helm instance                        | `""`                             |
| `koffloaderAgent.prometheus.grafanaDashboard.annotations`    | the additional annotations of koffloaderAgent grafanaDashboard                                   | `{}`                             |
| `koffloaderAgent.prometheus.grafanaDashboard.labels`         | the additional label of koffloaderAgent grafanaDashboard                                         | `{}`                             |
| `koffloaderAgent.debug.logLevel`                             | the log level of template agent [debug, info, warn, error, fatal, panic]                         | `info`                           |
| `koffloaderAgent.debug.gopsPort`                             | the gops port of template agent                                                                  | `5712`                           |

### koffloaderController parameters

| Name                                                              | Description                                                                                                                     | Value                                 |
| ----------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------- |
| `koffloaderController.name`                                       | the koffloaderController name                                                                                                   | `koffloader-controller`               |
| `koffloaderController.replicas`                                   | the replicas number of koffloaderController pod                                                                                 | `1`                                   |
| `koffloaderController.cmdBinName`                                 | the binName name of koffloaderController                                                                                        | `/usr/bin/controller`                 |
| `koffloaderController.hostnetwork`                                | enable hostnetwork mode of koffloaderController pod. Notice, if no CNI available before template installation, must enable this | `false`                               |
| `koffloaderController.image.registry`                             | the image registry of koffloaderController                                                                                      | `ghcr.io`                             |
| `koffloaderController.image.repository`                           | the image repository of koffloaderController                                                                                    | `koffloader-io/koffloader-controller` |
| `koffloaderController.image.pullPolicy`                           | the image pullPolicy of koffloaderController                                                                                    | `IfNotPresent`                        |
| `koffloaderController.image.digest`                               | the image digest of koffloaderController, which takes preference over tag                                                       | `""`                                  |
| `koffloaderController.image.tag`                                  | the image tag of koffloaderController, overrides the image tag whose default is the chart appVersion.                           | `""`                                  |
| `koffloaderController.image.imagePullSecrets`                     | the image imagePullSecrets of koffloaderController                                                                              | `[]`                                  |
| `koffloaderController.serviceAccount.create`                      | create the service account for the koffloaderController                                                                         | `true`                                |
| `koffloaderController.serviceAccount.annotations`                 | the annotations of koffloaderController service account                                                                         | `{}`                                  |
| `koffloaderController.service.annotations`                        | the annotations for koffloaderController service                                                                                | `{}`                                  |
| `koffloaderController.service.type`                               | the type for koffloaderController service                                                                                       | `ClusterIP`                           |
| `koffloaderController.priorityClassName`                          | the priority Class Name for koffloaderController                                                                                | `system-node-critical`                |
| `koffloaderController.affinity`                                   | the affinity of koffloaderController                                                                                            | `{}`                                  |
| `koffloaderController.extraArgs`                                  | the additional arguments of koffloaderController container                                                                      | `[]`                                  |
| `koffloaderController.extraEnv`                                   | the additional environment variables of koffloaderController container                                                          | `[]`                                  |
| `koffloaderController.extraVolumes`                               | the additional volumes of koffloaderController container                                                                        | `[]`                                  |
| `koffloaderController.extraVolumeMounts`                          | the additional hostPath mounts of koffloaderController container                                                                | `[]`                                  |
| `koffloaderController.podAnnotations`                             | the additional annotations of koffloaderController pod                                                                          | `{}`                                  |
| `koffloaderController.podLabels`                                  | the additional label of koffloaderController pod                                                                                | `{}`                                  |
| `koffloaderController.securityContext`                            | the security Context of koffloaderController pod                                                                                | `{}`                                  |
| `koffloaderController.resources.limits.cpu`                       | the cpu limit of koffloaderController pod                                                                                       | `500m`                                |
| `koffloaderController.resources.limits.memory`                    | the memory limit of koffloaderController pod                                                                                    | `1024Mi`                              |
| `koffloaderController.resources.requests.cpu`                     | the cpu requests of koffloaderController pod                                                                                    | `100m`                                |
| `koffloaderController.resources.requests.memory`                  | the memory requests of koffloaderController pod                                                                                 | `128Mi`                               |
| `koffloaderController.podDisruptionBudget.enabled`                | enable podDisruptionBudget for koffloaderController pod                                                                         | `false`                               |
| `koffloaderController.podDisruptionBudget.minAvailable`           | minimum number/percentage of pods that should remain scheduled.                                                                 | `1`                                   |
| `koffloaderController.httpServer.port`                            | the http Port for koffloaderController, for health checking and http service                                                    | `5720`                                |
| `koffloaderController.httpServer.startupProbe.failureThreshold`   | the failure threshold of startup probe for koffloaderController health checking                                                 | `30`                                  |
| `koffloaderController.httpServer.startupProbe.periodSeconds`      | the period seconds of startup probe for koffloaderController health checking                                                    | `2`                                   |
| `koffloaderController.httpServer.livenessProbe.failureThreshold`  | the failure threshold of startup probe for koffloaderController health checking                                                 | `6`                                   |
| `koffloaderController.httpServer.livenessProbe.periodSeconds`     | the period seconds of startup probe for koffloaderController health checking                                                    | `10`                                  |
| `koffloaderController.httpServer.readinessProbe.failureThreshold` | the failure threshold of startup probe for koffloaderController health checking                                                 | `3`                                   |
| `koffloaderController.httpServer.readinessProbe.periodSeconds`    | the period seconds of startup probe for koffloaderController health checking                                                    | `10`                                  |
| `koffloaderController.webhookPort`                                | the http port for koffloaderController webhook                                                                                  | `5722`                                |
| `koffloaderController.prometheus.enabled`                         | enable template Controller to collect metrics                                                                                   | `false`                               |
| `koffloaderController.prometheus.port`                            | the metrics port of template Controller                                                                                         | `5721`                                |
| `koffloaderController.prometheus.serviceMonitor.install`          | install serviceMonitor for template agent. This requires the prometheus CRDs to be available                                    | `false`                               |
| `koffloaderController.prometheus.serviceMonitor.namespace`        | the serviceMonitor namespace. Default to the namespace of helm instance                                                         | `""`                                  |
| `koffloaderController.prometheus.serviceMonitor.annotations`      | the additional annotations of koffloaderController serviceMonitor                                                               | `{}`                                  |
| `koffloaderController.prometheus.serviceMonitor.labels`           | the additional label of koffloaderController serviceMonitor                                                                     | `{}`                                  |
| `koffloaderController.prometheus.prometheusRule.install`          | install prometheusRule for template agent. This requires the prometheus CRDs to be available                                    | `false`                               |
| `koffloaderController.prometheus.prometheusRule.namespace`        | the prometheusRule namespace. Default to the namespace of helm instance                                                         | `""`                                  |
| `koffloaderController.prometheus.prometheusRule.annotations`      | the additional annotations of koffloaderController prometheusRule                                                               | `{}`                                  |
| `koffloaderController.prometheus.prometheusRule.labels`           | the additional label of koffloaderController prometheusRule                                                                     | `{}`                                  |
| `koffloaderController.prometheus.grafanaDashboard.install`        | install grafanaDashboard for template agent. This requires the prometheus CRDs to be available                                  | `false`                               |
| `koffloaderController.prometheus.grafanaDashboard.namespace`      | the grafanaDashboard namespace. Default to the namespace of helm instance                                                       | `""`                                  |
| `koffloaderController.prometheus.grafanaDashboard.annotations`    | the additional annotations of koffloaderController grafanaDashboard                                                             | `{}`                                  |
| `koffloaderController.prometheus.grafanaDashboard.labels`         | the additional label of koffloaderController grafanaDashboard                                                                   | `{}`                                  |
| `koffloaderController.debug.logLevel`                             | the log level of template Controller [debug, info, warn, error, fatal, panic]                                                   | `info`                                |
| `koffloaderController.debug.gopsPort`                             | the gops port of template Controller                                                                                            | `5724`                                |
| `koffloaderController.tls.method`                                 | the method for generating TLS certificates. [ provided , certmanager , auto]                                                    | `auto`                                |
| `koffloaderController.tls.secretName`                             | the secret name for storing TLS certificates                                                                                    | `template-controller-server-certs`    |
| `koffloaderController.tls.certmanager.certValidityDuration`       | generated certificates validity duration in days for 'certmanager' method                                                       | `365`                                 |
| `koffloaderController.tls.certmanager.issuerName`                 | issuer name of cert manager 'certmanager'. If not specified, a CA issuer will be created.                                       | `""`                                  |
| `koffloaderController.tls.certmanager.extraDnsNames`              | extra DNS names added to certificate when it's auto generated                                                                   | `[]`                                  |
| `koffloaderController.tls.certmanager.extraIPAddresses`           | extra IP addresses added to certificate when it's auto generated                                                                | `[]`                                  |
| `koffloaderController.tls.provided.tlsCert`                       | encoded tls certificate for provided method                                                                                     | `""`                                  |
| `koffloaderController.tls.provided.tlsKey`                        | encoded tls key for provided method                                                                                             | `""`                                  |
| `koffloaderController.tls.provided.tlsCa`                         | encoded tls CA for provided method                                                                                              | `""`                                  |
| `koffloaderController.tls.auto.caExpiration`                      | ca expiration for auto method                                                                                                   | `73000`                               |
| `koffloaderController.tls.auto.certExpiration`                    | server cert expiration for auto method                                                                                          | `73000`                               |
| `koffloaderController.tls.auto.extraIpAddresses`                  | extra IP addresses of server certificate for auto method                                                                        | `[]`                                  |
| `koffloaderController.tls.auto.extraDnsNames`                     | extra DNS names of server cert for auto method                                                                                  | `[]`                                  |
