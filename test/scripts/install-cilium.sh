#!/bin/bash

# SPDX-License-Identifier: Apache-2.0
# Copyright Authors of koffloader

set -o errexit -o nounset -o pipefail

CURRENT_FILENAME=$( basename $0 )
CURRENT_DIR_PATH=$(cd $(dirname $0); pwd)

[ -z "$E2E_KIND_CLUSTER_NAME" ] && echo "error, miss E2E_KIND_CLUSTER_NAME " && exit 1
echo "$CURRENT_FILENAME : E2E_KIND_CLUSTER_NAME $E2E_KIND_CLUSTER_NAME "

[ -z "$CILIUM_CLUSTER_ID" ] && echo "error, miss CILIUM_CLUSTER_ID " && exit 1
echo "$CURRENT_FILENAME : CILIUM_CLUSTER_ID $CILIUM_CLUSTER_ID "

[ -z "$E2E_IP_FAMILY" ] && echo "error, miss E2E_IP_FAMILY" && exit 1
echo "$CURRENT_FILENAME : E2E_IP_FAMILY $E2E_IP_FAMILY "

[ -z "$KIND_CLUSTER_KUBECONFIG" ] && echo "error, miss KIND_CLUSTER_KUBECONFIG " && exit 1
[ ! -f "$KIND_CLUSTER_KUBECONFIG" ] && echo "error, could not find file $KIND_CLUSTER_KUBECONFIG " && exit 1
echo "$CURRENT_FILENAME : KIND_CLUSTER_KUBECONFIG $KIND_CLUSTER_KUBECONFIG "

E2E_CILIUM_IMAGE_REPO=${E2E_CILIUM_IMAGE_REPO:-"quay.io"}
CILIUM_VERSION=${CILIUM_VERSION:-""}

CILIUM_HELM_OPTIONS=" --set cluster.id=${CILIUM_CLUSTER_ID} \
                            --set cluster.name=${E2E_KIND_CLUSTER_NAME} \
                            --set image.repository=${E2E_CILIUM_IMAGE_REPO}/cilium/cilium \
                            --set image.useDigest=false \
                            --set certgen.image.repository=${E2E_CILIUM_IMAGE_REPO}/cilium/certgen \
                            --set hubble.relay.image.repository=${E2E_CILIUM_IMAGE_REPO}/cilium/hubble-relay \
                            --set hubble.relay.image.useDigest=false \
                            --set hubble.ui.backend.image.repository=${E2E_CILIUM_IMAGE_REPO}/cilium/hubble-ui-backend \
                            --set hubble.ui.frontend.image.repository=${E2E_CILIUM_IMAGE_REPO}/cilium/hubble-ui \
                            --set etcd.image.repository=${E2E_CILIUM_IMAGE_REPO}/cilium/cilium-etcd-operator \
                            --set operator.image.repository=${E2E_CILIUM_IMAGE_REPO}/cilium/operator  \
                            --set operator.image.useDigest=false  \
                            --set preflight.image.repository=${E2E_CILIUM_IMAGE_REPO}/cilium/cilium \
                            --set preflight.image.useDigest=false \
                            --set nodeinit.image.repository=${E2E_CILIUM_IMAGE_REPO}/cilium/startup-script "


if [ -n "${CILIUM_VERSION}" ] ; then
    CILIUM_HELM_OPTIONS+=" --version ${CILIUM_VERSION} "
fi

case ${E2E_IP_FAMILY} in
  ipv4)
      CILIUM_HELM_OPTIONS+=" --set ipam.operator.clusterPoolIPv4PodCIDRList=${CILIUM_CLUSTER_IPV4_CIDR} \
                             --set ipv4.enabled=true \
                             --set ipv6.enabled=false "
    ;;
  ipv6)
      CILIUM_HELM_OPTIONS+=" --set ipam.operator.clusterPoolIPv6PodCIDRList=${CILIUM_CLUSTER_IPV6_CIDR} \
                             --set ipv4.enabled=false \
                             --set ipv6.enabled=true \
                             --set tunnel=disabled \
                             --set ipv6NativeRoutingCIDR=${CILIUM_CLUSTER_IPV6_CIDR} \
                             --set autoDirectNodeRoutes=true \
                             --set enableIPv6Masquerade=true "
    ;;
  dual)
      CILIUM_HELM_OPTIONS+=" --set ipam.operator.clusterPoolIPv4PodCIDRList=${CILIUM_CLUSTER_IPV4_CIDR} \
                             --set ipam.operator.clusterPoolIPv6PodCIDRList=${CILIUM_CLUSTER_IPV6_CIDR} \
                             --set ipv4.enabled=true \
                             --set ipv6.enabled=true "
    ;;
  *)
    echo "the value of E2E_IP_FAMILY: ipv4 or ipv6 or dual"
    exit 1
esac


echo "CILIUM_HELM_OPTIONS: ${CILIUM_HELM_OPTIONS}"

[ -z "${HTTP_PROXY}" ] || export https_proxy=${HTTP_PROXY}

helm repo remove cilium &>/dev/null || true
helm repo add cilium https://helm.cilium.io
helm repo update cilium

HELM_IMAGES_LIST=` helm template test cilium/cilium ${CILIUM_HELM_OPTIONS} | grep " image: " | tr -d '"'| awk '{print $2}' | uniq `

[ -z "${HELM_IMAGES_LIST}" ] && echo "can't found image of cilium" && exit 1
LOCAL_IMAGE_LIST=`docker images | awk '{printf("%s:%s\n",$1,$2)}'`

for IMAGE in ${HELM_IMAGES_LIST}; do
  if ! grep ${IMAGE} <<< ${LOCAL_IMAGE_LIST}; then
      echo "===> docker pull ${IMAGE}... "
      docker pull ${IMAGE}
  fi
  echo "===> load image ${IMAGE} to kind..."
  kind load docker-image ${IMAGE} --name $E2E_KIND_CLUSTER_NAME
done

# Install cilium
helm upgrade --install cilium cilium/cilium -n kube-system --debug --kubeconfig ${KIND_CLUSTER_KUBECONFIG} ${CILIUM_HELM_OPTIONS}

# no matching resources found
sleep 3

kubectl wait --for=condition=ready -l k8s-app=cilium --timeout=300s pod -n kube-system \
--kubeconfig ${KIND_CLUSTER_KUBECONFIG}

echo -e "\033[35m Succeed to install cilium  on $E2E_KIND_CLUSTER_NAME \033[0m"