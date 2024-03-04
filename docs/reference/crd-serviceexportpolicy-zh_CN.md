# ServiceExportPolicy

ServiceExportPolicy crd 表示 koffloader 的 service 发现策略。

## Sample YAML

```yaml
apiVersion: koffloader.io/koffloader/v1beta1
kind: ServiceExportPolicy
metadata:
  name: service-policy
spec:
  serviceSelector:
    matchLabels:
      app: test
  ciliumServiceFeature:
    share: ture
    affinity: remote
    global: true
  kclusterSelector:
    matchLabels:
      app: test  
status:
  matchService:
    - name: test
      namespace: default
  matchKCluster:
    - cluster1
    - cluster2
```

## ServiceExportPolicy definition

### Metadata

| Field | Description                   | Schema  | Validation |
|-------|-------------------------------|---------|------------|
| name  | ServiceExportPolicy 资源名称   | string  | required   |

### Spec

| Field                | Description                        | Schema                                                                                                                                   | Validation | Values    | Default |
|----------------------|------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------|------------|-----------|---------|
| serviceSelector      | 指定哪些 service 需要发现                  | [labelSelector](https://github.com/kubernetes/kubernetes/blob/v1.29.0/staging/src/k8s.io/apimachinery/pkg/apis/meta/v1/types.go#L1213)   | required   |           |         |
| kclusterSelector     | 指定哪些 kcluster 生效该策略                | [labelSelector](https://github.com/kubernetes/kubernetes/blob/v1.29.0/staging/src/k8s.io/apimachinery/pkg/apis/meta/v1/types.go#L1213)   | required   |           |         |
| ciliumServiceFeature | 使用 cilium 进行多集群连通时，service 的发现策略配置 | [ciliumServiceFeature](./crd-kcluster-zh_CN.md#koffloaderClusterConfig)                                                                  | optional   |           |         |



#### ciliumServiceFeature

| Field    | Description               | Schema | Validation | Values            |
|----------|---------------------------|--------|------------|-------------------|
| share    | service 是否可以被其他集群发现       | bool   | optional   |                   |
| affinity | service 负载的优先级            | string | optional   | local,remote,none |
| global   | service 是否作为多集群全局 service | bool   | optional   |                   |


### Status (subresource)

| Field               | Description        | Schema                                                    |
|---------------------|--------------------|-----------------------------------------------------------|
| matchService        | 策略生效在哪些 service    | [matchService](./crd-kcluster-zh_CN.md#matchService) 数组   |
| matchKCluster       | 策略生效在哪些 kcluster   | string 数组                                                 |

#### matchService

| Field     | Description       | Schema    | 
|-----------|-------------------|-----------|
| name      | service 名称        | string    |        
| namespace | service 命名空间      | string    |