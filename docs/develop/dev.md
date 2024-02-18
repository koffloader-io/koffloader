# develop

## local develop

1. `make build_local_image`

2. `make e2e_init`

3. `make e2e_run`

4. check proscope, browser visits http://NodeIP:4040

5. check metric, 

## chart develop

helm repo add koffloader https://koffloader-io.github.io/koffloader

## test 


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
