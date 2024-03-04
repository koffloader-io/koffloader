# KCluster

kcluster crd 表示 koffloader 的 cluster 集群资源，存储集群的 kubeconfig 连接信息以及打通集群连接的配置信息。

## Sample YAML

```yaml
apiVersion: koffloader.io/koffloader/v1beta1
kind: KCluster
metadata:
  name: cluster1
spec:
  clusterConnector:
    type: cilium/submariner
    ciliumNamespace: kube-system
  kubeconfig:
    secretName: cluster1
    secretNamespace: koffloader-system
status:
  kmulticluster: koffloader-clusters
  clusterConnector: cilium
```

## KCluster definition

### Metadata

| Field | Description     | Schema  | Validation |
|-------|-----------------|---------|------------|
| name  | kcluster 资源名称   | string  | required   |

### Spec

| Field            | Description                            | Schema                                                                            | Validation | Values    | Default |
|------------------|----------------------------------------|-----------------------------------------------------------------------------------|------------|-----------|---------|
| clusterConnector | 指定使用哪种方式打通集群的连通性（cilium 或 submariner）  | [koffloaderClusterConnector](./crd-kcluster-zh_CN.md#koffloaderClusterConnector)  | required   |           |         |
| kubeconfig       | 集群 kubeconfig 连接信息存储                   | [koffloaderClusterConfig](./crd-kcluster-zh_CN.md#koffloaderClusterConfig)        | required   |           |         |



#### koffloaderClusterConnector

| Field           | Description    | Schema | Validation | Values            |
|-----------------|----------------|--------|------------|-------------------|
| type            | 指定打通多集群连通性的方式  | string | required   | cilium、submariner |
| ciliumNamespace | cilium 所在的命名空间 | string | optional   |                   |


#### koffloaderClusterConfig

| Field           | Description                   | Schema | Validation | Values   |
|-----------------|-------------------------------|--------|------------|----------|
| secretName      | 集群 kubeconfig 存储的 secret 名称   | string | required   |          |
| secretNamespace | 集群 kubeconfig 存储的 secret 命名空间 | string | required   |          |

### Status (subresource)

| Field                | Description             | Schema   |
|----------------------|-------------------------|----------|
| kmulticluster        | 所属 kmulticluster 资源名称   | string   |
| clusterConnector     | 集群连通器类型                 | string   |