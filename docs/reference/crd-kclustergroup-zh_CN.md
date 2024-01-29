# KClusterGroup

KClusterGroup crd 表示多集群连通的配置。

## Sample YAML

```yaml
apiVersion: koffloader.io/koffloader/v1beta1
kind: KClusterGroup
metadata:
  name: koffloader-cluster1
spec:
  clusterConnectorType: cilium/submariner
  ciliumClusterMeshServiceType: nodeport
  kclusterSelector:
    matchLabels:
      app: test
status:
  matchKCluster:
    - cluster1
    - cluster2
  clusterConnector: cilium
```

## KClusterGroup definition

### Metadata

| Field | Description         | Schema   | Validation   |
|-------|---------------------|----------|--------------|
| name  | KMultiCluster 资源名称  | string   | required     |

### Spec

| Field                        | Description                                    | Schema                                                                                                                                 | Validation | Values                          | Default |
|------------------------------|------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------|------------|---------------------------------|---------|
| kclusterSelector             | 指定哪些 kcluster 组成 KMultiCluster 多集群连通           | [labelSelector](https://github.com/kubernetes/kubernetes/blob/v1.29.0/staging/src/k8s.io/apimachinery/pkg/apis/meta/v1/types.go#L1213) | required   |                                 |         |
| clusterConnectorType         | 指定打通多集群连通性的方式                                  | string                                                                                                                                 | required   |                                 |         |
| ciliumClusterMeshServiceType | 指定 cilium clustermesh apiserver 的 service 类型   | string                                                                                                                                 | required   | NodePort、ClusterIP、LoadBalancer |         |


### Status (subresource)

| Field                    | Description                     | Schema      |
|--------------------------|---------------------------------|-------------|
| clusterConnector         | 多集群连通性类型                        | string      |
| matchKCluster            | 哪些 kcluster 资源组成 KMultiCluster  | string 数组   |
