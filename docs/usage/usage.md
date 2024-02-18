# Usage

```shell

cat <<EOF > kcluster1.yaml
apiVersion: koffloader.koffloader.io/v1
kind: Kcluster
metadata:
  name: test1
spec:
  ipVersion: 4
  subnet: "1.0.0.0/8"
EOF

kubectl apply -f kcluster1.yaml


cat <<EOF > kcluster2.yaml
apiVersion: koffloader.koffloader.io/v1
kind: Kcluster
metadata:
  name: test2
spec:
  ipVersion: 4
  subnet: "2.0.0.0/8"
EOF

kubectl apply -f kcluster2.yaml


```


