#!/bin/bash

# SPDX-License-Identifier: Apache-2.0
# Copyright Authors of koffloader

set -o errexit -o nounset -o pipefail

CURRENT_FILENAME=$( basename $0 )
CURRENT_DIR_PATH=$(cd $(dirname $0); pwd)

[ -z "$E2E_KIND_CLUSTER_NAME" ] && echo "error, miss $E2E_KIND_CLUSTER_NAME " && exit 1
echo "$CURRENT_FILENAME : E2E_KIND_CLUSTER_NAME $E2E_KIND_CLUSTER_NAME "

[ -z "$E2E_KIND_SERVICE_CIDR" ] && echo "error, miss E2E_KIND_SERVICE_CIDR" && exit 1
echo "$CURRENT_FILENAME : E2E_KIND_SERVICE_CIDR $E2E_KIND_SERVICE_CIDR "

[ -z "$E2E_KIND_POD_CIDR" ] && echo "error, miss E2E_KIND_POD_CIDR" && exit 1
echo "$CURRENT_FILENAME : E2E_KIND_POD_CIDR $E2E_KIND_POD_CIDR "

[ -z "$E2E_IP_FAMILY" ] && echo "error, miss E2E_IP_FAMILY" && exit 1
echo "$CURRENT_FILENAME : E2E_IP_FAMILY $E2E_IP_FAMILY "

[ -z "$CLUSTER_DIR" ] && echo "error, miss CLUSTER_DIR" && exit 1
echo "$CURRENT_FILENAME : CLUSTER_DIR $CLUSTER_DIR "

[ -z "$E2E_KUBE_PROXY_MODE" ] && echo "error, miss E2E_KUBE_PROXY_MODE" && exit 1
echo "$CURRENT_FILENAME : E2E_KUBE_PROXY_MODE $E2E_KUBE_PROXY_MODE "

[ -z "$KIND_CLUSTER_KUBECONFIG" ] && echo "error, miss KIND_CLUSTER_KUBECONFIG" && exit 1
echo "$CURRENT_FILENAME : KIND_CLUSTER_KUBECONFIG $KIND_CLUSTER_KUBECONFIG "

[ -z "$GLOBAL_KIND_CONFIG_PATH" ] && echo "error, miss GLOBAL_KIND_CONFIG_PATH" && exit 1
echo "$CURRENT_FILENAME : GLOBAL_KIND_CONFIG_PATH $GLOBAL_KIND_CONFIG_PATH "

[ -z "$E2E_DISABLE_DEFAULT_CNI" ] && echo "error, miss E2E_DISABLE_DEFAULT_CNI" && exit 1
echo "$CURRENT_FILENAME : E2E_DISABLE_DEFAULT_CNI $E2E_DISABLE_DEFAULT_CNI "


# delete exist kind cluster
kind delete cluster --name ${E2E_KIND_CLUSTER_NAME} &>/dev/null
rm -rf ${CLUSTER_DIR}/${E2E_KIND_CLUSTER_NAME}
mkdir -p -v ${CLUSTER_DIR}/${E2E_KIND_CLUSTER_NAME}

NEW_KIND_YAML=${CLUSTER_DIR}/${E2E_KIND_CLUSTER_NAME}/kind-config.yaml

INSERT_LINE=` grep "insert inform" ${GLOBAL_KIND_CONFIG_PATH} -n | awk -F':' '{print $1}' `;
echo "insert after line ${INSERT_LINE}" ;
sed  ${INSERT_LINE}" a \  kubeProxyMode: ${E2E_KUBE_PROXY_MODE}" ${GLOBAL_KIND_CONFIG_PATH} > ${NEW_KIND_YAML}
sed -i ${INSERT_LINE}" a \  disableDefaultCNI: ${E2E_DISABLE_DEFAULT_CNI}" ${NEW_KIND_YAML}
sed -i ${INSERT_LINE}" a \  ipFamily: ${E2E_IP_FAMILY}" ${NEW_KIND_YAML}
sed -i ${INSERT_LINE}" a \  podSubnet: ${E2E_KIND_POD_CIDR}" ${NEW_KIND_YAML}
sed -i ${INSERT_LINE}" a \  serviceSubnet: ${E2E_KIND_SERVICE_CIDR}" ${NEW_KIND_YAML}
cat ${CLUSTER_DIR}/${E2E_KIND_CLUSTER_NAME}/kind-config.yaml

KIND_OPTION="" ; \
       	[ -n "${E2E_KIND_NODE_IMAGE}" ] && KIND_OPTION=" --image ${E2E_KIND_NODE_IMAGE} " && echo "setup kind with E2E_KIND_NODE_IMAGE=${E2E_KIND_NODE_IMAGE}"; \
            kind create cluster --config ${CLUSTER_DIR}/${E2E_KIND_CLUSTER_NAME}/kind-config.yaml \
			--name ${E2E_KIND_CLUSTER_NAME} --kubeconfig ${KIND_CLUSTER_KUBECONFIG} ${KIND_OPTION}

kubectl --kubeconfig ${KIND_CLUSTER_KUBECONFIG} taint nodes --all node-role.kubernetes.io/master- || true
kubectl --kubeconfig ${KIND_CLUSTER_KUBECONFIG} taint nodes --all node-role.kubernetes.io/control-plane- || true
for ((N=0;N<=30;N++)); do \
  sleep 1 ; \
  kubectl get node --kubeconfig ${KIND_CLUSTER_KUBECONFIG} &>/dev/null && break ; \
  echo "wait for node ready" ; \
done ; \
  kubectl get node --kubeconfig ${KIND_CLUSTER_KUBECONFIG} &>/dev/null || { echo "error, cluster is not ready" ; exit 1 ; }
echo "show kubernetes node image " && docker ps
echo "===================== deploy prometheus CRD ========== "
kubectl apply --kubeconfig ${KIND_CLUSTER_KUBECONFIG}  -f ${CURRENT_DIR_PATH}/../yamls/monitoring.coreos.com_servicemonitors.yaml
kubectl apply --kubeconfig ${KIND_CLUSTER_KUBECONFIG} -f ${CURRENT_DIR_PATH}/../yamls/monitoring.coreos.com_podmonitors.yaml
kubectl apply --kubeconfig ${KIND_CLUSTER_KUBECONFIG} -f ${CURRENT_DIR_PATH}/../yamls/monitoring.coreos.com_prometheusrules.yaml
kubectl apply --kubeconfig ${KIND_CLUSTER_KUBECONFIG} -f ${CURRENT_DIR_PATH}/../yamls/monitoring.coreos.com_probes.yaml
kubectl apply --kubeconfig ${KIND_CLUSTER_KUBECONFIG} -f ${CURRENT_DIR_PATH}/../yamls/grafanadashboards.yaml
echo "success setup kind cluster"