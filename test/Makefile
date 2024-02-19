include ../Makefile.defs
include Makefile.defs

.PHONY: kind-init
kind-init: check_env prepare
	@echo -e "\033[35mInit kind for the cluster \033[0m"
	make setup_three_kind
	@if [ "$(INSTALL_CILIUM)" == "true" ] ; then \
  		echo -e "\033[35m install cilium for the clusters \033[0m" ;\
		make install_cilium_on_all_cluster ; \
	fi ;
	make install_proscope
	@echo "success init kind clusters"
	@echo ""
	@echo "-----------------------------------------------------------------------------------------------------"
	@echo "       ip family: $(E2E_IP_FAMILY)"
	@for CLUSTER_ID in $(KIND_CLUSTER_IDS); do \
		E2E_KIND_CLUSTER_NAME=$$(echo $(KIND_CLUSTER_NAMES) | cut -d ' ' -f $$CLUSTER_ID); \
		KIND_CLUSTER_KUBECONFIG=$$(echo $(KIND_CLUSTER_KUBECONFIGS) | cut -d ' ' -f $$CLUSTER_ID); \
		echo -e "       cluster name: " $$E2E_KIND_CLUSTER_NAME; \
		echo -e "       cluster kubeconfig:" $$KIND_CLUSTER_KUBECONFIG ; \
		echo "" ;\
	done
	@echo "-----------------------------------------------------------------------------------------------------"
	@echo ""


.PHONY: setup_one_kind
setup_one_kind: E2E_KIND_CLUSTER_NAME ?=
setup_one_kind: E2E_KIND_SERVICE_CIDR ?=
setup_one_kind: E2E_KIND_POD_CIDR ?=
setup_one_kind: KIND_CLUSTER_KUBECONFIG ?=
setup_one_kind:
	export E2E_IP_FAMILY=$(E2E_IP_FAMILY) ; \
	echo E2E_IP_FAMILY=$(E2E_IP_FAMILY) ; \
	export CLUSTER_DIR=$(CLUSTER_DIR) ; \
	echo CLUSTER_DIR=$(CLUSTER_DIR) ; \
	export E2E_KUBE_PROXY_MODE=$(E2E_KUBE_PROXY_MODE) ; \
	echo E2E_KUBE_PROXY_MODE=$(E2E_KUBE_PROXY_MODE) ; \
	export E2E_KIND_NODE_IMAGE=$(E2E_KIND_NODE_IMAGE) ; \
	echo E2E_KIND_NODE_IMAGE=$(E2E_KIND_NODE_IMAGE) ; \
	export GLOBAL_KIND_CONFIG_PATH=$(GLOBAL_KIND_CONFIG_PATH) ; \
	echo GLOBAL_KIND_CONFIG_PATH=$(GLOBAL_KIND_CONFIG_PATH) ; \
	export E2E_DISABLE_DEFAULT_CNI=$(E2E_DISABLE_DEFAULT_CNI) ; \
	echo E2E_DISABLE_DEFAULT_CNI=$(E2E_DISABLE_DEFAULT_CNI) ; \
	export E2E_KIND_CLUSTER_NAME=$$E2E_KIND_CLUSTER_NAME ; \
	echo E2E_KIND_CLUSTER_NAME=$$E2E_KIND_CLUSTER_NAME ; \
	export E2E_KIND_SERVICE_CIDR=$$E2E_KIND_SERVICE_CIDR ; \
	echo E2E_KIND_SERVICE_CIDR=$$E2E_KIND_SERVICE_CIDR ; \
	export E2E_KIND_POD_CIDR=$$E2E_KIND_POD_CIDR ; \
	echo E2E_KIND_POD_CIDR=$$E2E_KIND_POD_CIDR ; \
	export KIND_CLUSTER_KUBECONFIG=$$KIND_CLUSTER_KUBECONFIG ; \
	echo KIND_CLUSTER_KUBECONFIG=$$KIND_CLUSTER_KUBECONFIG ; \
	echo setup $$E2E_KIND_CLUSTER_NAME kind cluster; \
	bash $(ROOT_DIR)/test/scripts/init-kind.sh ;


.PHONY: setup_three_kind
setup_three_kind:
	@for CLUSTER_ID in $(KIND_CLUSTER_IDS); do \
		E2E_KIND_CLUSTER_NAME=$$(echo $(KIND_CLUSTER_NAMES) | cut -d ' ' -f $$CLUSTER_ID); \
		if [ "$(E2E_IP_FAMILY)" == "ipv4" ] ; then \
			E2E_KIND_SERVICE_CIDR=$$(echo $(KIND_CLUSTER_SERVICE_IPV4_CIDRS) | cut -d ' ' -f $$CLUSTER_ID); \
			E2E_KIND_POD_CIDR=$$(echo $(KIND_CLUSTER_POD_IPV4_CIDRS) | cut -d ' ' -f $$CLUSTER_ID); \
		elif [ "$(E2E_IP_FAMILY)" == "ipv6" ] ; then \
			E2E_KIND_SERVICE_CIDR=$$(echo $(KIND_CLUSTER_SERVICE_IPV6_CIDRS) | cut -d ' ' -f $$CLUSTER_ID); \
			E2E_KIND_POD_CIDR=$$(echo $(KIND_CLUSTER_POD_IPV6_CIDRS) | cut -d ' ' -f $$CLUSTER_ID); \
		else \
			E2E_KIND_SERVICE_CIDR=$$(echo $(KIND_CLUSTER_SERVICE_DUAL_CIDRS) | cut -d ' ' -f $$CLUSTER_ID); \
			E2E_KIND_POD_CIDR=$$(echo $(KIND_CLUSTER_POD_DUAL_CIDRS) | cut -d ' ' -f $$CLUSTER_ID); \
		fi ; \
		KIND_CLUSTER_KUBECONFIG=$$(echo $(KIND_CLUSTER_KUBECONFIGS) | cut -d ' ' -f $$CLUSTER_ID); \
		make setup_one_kind -e E2E_KIND_CLUSTER_NAME=$$E2E_KIND_CLUSTER_NAME \
				-e E2E_KIND_SERVICE_CIDR=$$E2E_KIND_SERVICE_CIDR \
				-e E2E_KIND_POD_CIDR=$$E2E_KIND_POD_CIDR \
				-e KIND_CLUSTER_KUBECONFIG=$$KIND_CLUSTER_KUBECONFIG; \
	done


.PHONY: install_cilium
install_cilium: E2E_KIND_CLUSTER_NAME ?=
install_cilium: CILIUM_CLUSTER_ID ?=
install_cilium: E2E_IP_FAMILY ?= $E2E_IP_FAMILY
install_cilium: KIND_CLUSTER_KUBECONFIG ?=
install_cilium: HTTP_PROXY ?=
install_cilium: CILIUM_CLUSTER_IPV4_CIDR ?=
install_cilium: CILIUM_CLUSTER_IPV6_CIDR ?=
install_cilium: E2E_CILIUM_IMAGE_REPO ?= $(E2E_CILIUM_IMAGE_REPO)
install_cilium:
	echo "install cilium on" $(E2E_KIND_CLUSTER_NAME) ; \
	export E2E_KIND_CLUSTER_NAME=$(E2E_KIND_CLUSTER_NAME); \
	echo E2E_KIND_CLUSTER_NAME=$(E2E_KIND_CLUSTER_NAME); \
	export CILIUM_CLUSTER_ID=$(CILIUM_CLUSTER_ID); \
	echo CILIUM_CLUSTER_ID=$(CILIUM_CLUSTER_ID); \
	export E2E_IP_FAMILY=$(E2E_IP_FAMILY); \
	echo E2E_IP_FAMILY=$(E2E_IP_FAMILY); \
	export KIND_CLUSTER_KUBECONFIG=$(KIND_CLUSTER_KUBECONFIG); \
	echo KIND_CLUSTER_KUBECONFIG=$(KIND_CLUSTER_KUBECONFIG); \
	export HTTP_PROXY=$(HTTP_PROXY); \
	echo HTTP_PROXY=$(HTTP_PROXY); \
	export E2E_CILIUM_IMAGE_REPO=$(E2E_CILIUM_IMAGE_REPO); \
	echo E2E_CILIUM_IMAGE_REPO=$(E2E_CILIUM_IMAGE_REPO); \
	export CILIUM_CLUSTER_IPV4_CIDR=$(CILIUM_CLUSTER_IPV4_CIDR); \
	echo CILIUM_CLUSTER_IPV4_CIDR=$(CILIUM_CLUSTER_IPV4_CIDR); \
	export CILIUM_CLUSTER_IPV6_CIDR=$(CILIUM_CLUSTER_IPV6_CIDR); \
	echo CILIUM_CLUSTER_IPV6_CIDR=$(CILIUM_CLUSTER_IPV6_CIDR); \
	bash $(ROOT_DIR)/test/scripts/install-cilium.sh ;


.PHONY: install_cilium_on_all_cluster
install_cilium_on_all_cluster:
	@for CLUSTER_ID in $(KIND_CLUSTER_IDS); do \
		E2E_KIND_CLUSTER_NAME=$$(echo $(KIND_CLUSTER_NAMES) | cut -d ' ' -f $$CLUSTER_ID); \
  		CILIUM_CLUSTER_ID=$$CLUSTER_ID ; \
		KIND_CLUSTER_KUBECONFIG=$$(echo $(KIND_CLUSTER_KUBECONFIGS) | cut -d ' ' -f $$CLUSTER_ID); \
		CILIUM_CLUSTER_IPV4_CIDR=$$(echo $(CILIUM_CLUSTER_IPV4_CIDRS) | cut -d ' ' -f $$CLUSTER_ID); \
		CILIUM_CLUSTER_IPV6_CIDR=$$(echo $(CILIUM_CLUSTER_IPV6_CIDRS) | cut -d ' ' -f $$CLUSTER_ID); \
		make install_cilium -e E2E_KIND_CLUSTER_NAME=$$E2E_KIND_CLUSTER_NAME -e CILIUM_CLUSTER_ID=$$CILIUM_CLUSTER_ID \
			-e E2E_IP_FAMILY=$(E2E_IP_FAMILY) -e KIND_CLUSTER_KUBECONFIG=$$KIND_CLUSTER_KUBECONFIG \
			-e HTTP_PROXY=$(HTTP_PROXY) -e E2E_CILIUM_IMAGE_REPO=$(E2E_CILIUM_IMAGE_REPO) \
			-e CILIUM_CLUSTER_IPV4_CIDR=$$CILIUM_CLUSTER_IPV4_CIDR -e CILIUM_CLUSTER_IPV6_CIDR=$$CILIUM_CLUSTER_IPV6_CIDR; \
	done


.PHONY: prepare
prepare:
	echo -e "\033[35m Check The Tools For Ready: \033[0m"
	$(QUIET) JUST_CLI_CHECK=true scripts/install-tools.sh
	$(QUIET) mkdir -p $(DOWNLOAD_DIR)
	$(QUIET) IMAGE_LIST="" ; \
		 [ -n "$(PYROSCOPE_LOCAL_PORT)" ] && IMAGE_LIST+=" $(IMAGE_PYROSCOPE_NAME) " ; \
		 IMAGE_LIST+=" $(TEST_IMAGE_NAME) " ; \
		 ARCH=$(ARCH) IMAGE_LIST="$${IMAGE_LIST}" \
		 	CNI_PACKAGE_VERSION=$(CNI_PACKAGE_VERSION)  scripts/prepare.sh $(DOWNLOAD_DIR)
	- sudo sysctl -w fs.inotify.max_user_watches=524288 || true
	- sudo sysctl -w fs.inotify.max_user_instances=8192 || true
	- sudo sysctl -w net.ipv6.conf.all.disable_ipv6=0


.PHONY: install_proscope
install_proscope:
	if [ -n "$(PYROSCOPE_LOCAL_PORT)" ] ; then \
  		echo "install proscope " ; \
		docker stop $(PYROSCOPE_CONTAINER_NAME) &>/dev/null || true ; \
		docker rm $(PYROSCOPE_CONTAINER_NAME) &>/dev/null || true ; \
		ServerAddress=$$(docker network inspect kind -f {{\(index\ $$.IPAM.Config\ 0\).Gateway}}) ; \
		echo "setup pyroscope on $${ServerAddress}:$(PYROSCOPE_LOCAL_PORT)" ; \
		docker run -d --name $(PYROSCOPE_CONTAINER_NAME) -p $(PYROSCOPE_LOCAL_PORT):4040 $(PYROSCOPE_IMAGE_NAME) server ; \
		echo "finish setuping pyroscope " ; \
      fi

.PHONY: check_env
check_env:
	$(QUIET) [ -n "$(E2E_CLUSTER_NAME)" ] || { echo "error, miss E2E_CLUSTER_NAME " ; false ; }
	$(QUIET) ( [ "$(E2E_IP_FAMILY)" == "ipv4" ] || [ "$(E2E_IP_FAMILY)" == "ipv6" ]  || [ "$(E2E_IP_FAMILY)" == "dual" ] ) \
			|| { echo "error, E2E_IP_FAMILY=$(E2E_IP_FAMILY) must be ipv4/ipv6/dual" ;  exit 1 ; }

#==================

# this will auto tag github ci image : agent:xxx -> github.com/koffloader-io/koffloader/agent:xxx
.PHONY: check_images_ready
check_images_ready:
	echo "check image  " ; \
	IMAGE_LIST=` helm template test $(ROOT_DIR)/charts --set global.imageTagOverride=$(PROJECT_IMAGE_VERSION)  | grep " image: " | tr -d '"'| awk '{print $$2}' ` ; \
	if [ -z "$${IMAGE_LIST}" ] ; then \
		echo "warning, failed to find image from chart " ; \
		exit 1 ;\
	else \
		echo "find image from chart : $${IMAGE_LIST} " ; \
		for IMAGE in $${IMAGE_LIST} ; do \
		  	echo "try to find image $${IMAGE} " ; \
			EXIST=` docker images | awk '{printf("%s:%s\n",$$1,$$2)}' | grep "$${IMAGE}" ` || true ; \
			if [ -z "$${EXIST}" ] ; then \
					CI_IMAGE=$${IMAGE##*/} ; \
			  		echo "try to find github CI image $${CI_IMAGE} " ; \
			  		EXIST=` docker images | awk '{printf("%s:%s\n",$$1,$$2)}' | grep "$${CI_IMAGE}" ` || true ; \
			  		if [ -z "$${EXIST}" ] ; then \
			  			echo "error, failed to find image $${IMAGE}" ; \
			  			echo "error, failed to find image $${CI_IMAGE}" ; \
			  			exit 1 ; \
			  		fi ; \
			  		docker tag $${CI_IMAGE} $${IMAGE} ; \
			fi ;\
			echo "image exists: $${IMAGE}" ; \
		done ; \
		docker images ; \
	fi


# install spidernet on global cluster
.PHONY: deploy_project
deploy_project: KIND_KUBECONFIG ?= $(E2E_KIND_KUBECONFIG_PATH)
deploy_project: KIND_CLUSTER_NAME ?= $(E2E_KIND_CLUSTER_NAME)
deploy_project:
	echo "try to load local image tag $(PROJECT_IMAGE_VERSION) " ; \
	IMAGE_LIST=` helm template test $(ROOT_DIR)/charts --set global.imageTagOverride=$(PROJECT_IMAGE_VERSION)  | grep " image: " | tr -d '"'| awk '{print $$2}' ` ; \
	if [ -z "$${IMAGE_LIST}" ] ; then \
		echo "warning, failed to find image from chart " ; \
	else \
		echo "found image from chart : $${IMAGE_LIST} " ; \
		for IMAGE in $${IMAGE_LIST} ; do \
			EXIST=` docker images | awk '{printf("%s:%s\n",$$1,$$2)}' | grep "$${IMAGE}" ` ; \
			if [ -z "$${EXIST}" ] ; then \
			  echo "docker pull $${IMAGE} to local" ; \
			  docker pull $${IMAGE} ; \
			fi ;\
			echo "load local image $${IMAGE} " ; \
			kind load docker-image $${IMAGE}  --name $(KIND_CLUSTER_NAME)  ; \
		done ; \
	fi
	- helm --kubeconfig=$(KIND_KUBECONFIG) uninstall -n $(E2E_INSTALL_NAMESPACE) project || true
	HELM_OPTION="" ; \
    	if [ -n "$(PYROSCOPE_LOCAL_PORT)" ] ; then \
			echo "add env" ; \
			ServerAddress=$$(docker network inspect kind -f {{\(index\ $$.IPAM.Config\ 0\).Gateway}}) ; \
			HELM_OPTION+=" --set koffloaderAgent.extraEnv[0].name=ENV_PYROSCOPE_PUSH_SERVER_ADDRESS  --set koffloaderAgent.extraEnv[0].value=http://$${ServerAddress}:$(PYROSCOPE_LOCAL_PORT) " ; \
			HELM_OPTION+=" --set koffloaderController.extraEnv[0].name=ENV_PYROSCOPE_PUSH_SERVER_ADDRESS  --set koffloaderController.extraEnv[0].value=http://$${ServerAddress}:$(PYROSCOPE_LOCAL_PORT) " ; \
		fi ; \
		HELM_OPTION+=" --set koffloaderAgent.debug.logLevel=debug --set koffloaderController.debug.logLevel=debug " ; \
		HELM_OPTION+=" --set koffloaderAgent.prometheus.enabled=true --set koffloaderController.prometheus.enabled=true  " ; \
		helm --kubeconfig=$(KIND_KUBECONFIG) install project $(ROOT_DIR)/charts \
				-n $(E2E_INSTALL_NAMESPACE) --create-namespace --wait --debug \
				--set global.imageTagOverride=$(PROJECT_IMAGE_VERSION) \
				$${HELM_OPTION} \
				|| { KIND_CLUSTER_NAME=$(KIND_CLUSTER_NAME) ./scripts/debugCluster.sh $(KIND_KUBECONFIG) "detail"  $(E2E_INSTALL_NAMESPACE) ; exit 1 ; } ; \
		exit 0


# test kind is ok
.PHONY: install_example_app
install_example_app: KIND_KUBECONFIG ?= $(E2E_KIND_KUBECONFIG_PATH)
install_example_app: KIND_CLUSTER_NAME ?= $(E2E_KIND_CLUSTER_NAME)
install_example_app:
	@echo "---------- install example app"
	kubectl --kubeconfig=$(KIND_KUBECONFIG) apply -f yaml/testpod.yaml
	@ if ! kubectl rollout status  deployment/test --kubeconfig $(KIND_KUBECONFIG) -w --timeout=120s ; then \
			echo "error, failed to create a test pod" ; \
			exit 1 ; \
		fi ; \
		echo "succeeded to deploy test deployment "
	@echo "========================================================"
	@echo "   deploy kind cluster $(KIND_CLUSTER_NAME)             "
	@echo "   export KUBECONFIG=$(KIND_KUBECONFIG)                 "
	@echo "   kubectl get pod -o wide -A                           "
	@echo "========================================================"
	@ KUBECONFIG=$(KIND_KUBECONFIG)  kubectl get pod -o wide -A


.PHONY: clean_e2e
clean_e2e:
	@rm -rf $(CLUSTER_DIR)
	-@  kind get clusters | xargs -n1  kind delete cluster --name
	-@ docker stop $(PYROSCOPE_CONTAINER_NAME) &>/dev/null
	-@ docker rm $(PYROSCOPE_CONTAINER_NAME) &>/dev/null



#============ e2e ====================
.PHONY: e2e_test
e2e_test: KIND_CLUSTER_NAME ?= $(E2E_KIND_CLUSTER_NAME)
e2e_test: KIND_KUBECONFIG ?= $(E2E_KIND_KUBECONFIG_PATH)
e2e_test:
	@echo -e "\033[35m Run e2e test on the cluster $(KIND_CLUSTER_NAME) \033[0m "
	@ echo -e "\033[35m [E2E] Run E2E with ginkgo label=$(E2E_GINKGO_LABELS) , timeout=$(E2E_TIMEOUT) GINKGO_OPTION=$(E2E_GINKGO_OPTION) \033[0m"
	@  NODE_LIST=` docker ps | egrep " kindest/node.* $(KIND_CLUSTER_NAME)-(control|worker)" | awk '{print $$NF }' ` ; \
		[ -n "$$NODE_LIST" ] || { echo "error, failed to find any kind nodes, please setup kind cluster $(KIND_CLUSTER_NAME) first" ; exit 1 ; } ; \
		NODE_LIST=` echo "$${NODE_LIST}" | tr -d ' ' | tr '\n' ',' ` ; \
		NODE_LIST=$${NODE_LIST%%,} ; \
		echo "find cluster node: $${NODE_LIST}" ; \
		export E2E_KIND_CLUSTER_NODE_LIST="$${NODE_LIST}" ; \
		export E2E_CLUSTER_NAME=$(KIND_CLUSTER_NAME) ; \
		if [ "$(E2E_IP_FAMILY)" == "ipv4" ] ; then \
			export E2E_IPV4_ENABLED=true ; export E2E_IPV6_ENABLED=false ; \
		elif [ "$(E2E_IP_FAMILY)" == "ipv6" ] ; then \
			export E2E_IPV4_ENABLED=false ; export E2E_IPV6_ENABLED=true ; \
		else \
			export E2E_IPV4_ENABLED=true ; export E2E_IPV6_ENABLED=true ; \
		fi ; \
		export E2E_KUBECONFIG_PATH=$(KIND_KUBECONFIG) ; [ -f "$(KIND_KUBECONFIG)" ] || { echo "error, does not exist KUBECONFIG $(E2E_KUBECONFIG)" ; exit 1 ; } ; \
		rm -f $(E2E_LOG_FILE) || true ; \
		echo "=========== before test `date` ===========" >> $(E2E_LOG_FILE) ; \
		./scripts/debugCluster.sh $(KIND_KUBECONFIG) "system" "$(E2E_LOG_FILE)" $(E2E_INSTALL_NAMESPACE) ; \
		RESULT=0 ; \
		$(ROOT_DIR)/tools/golang/ginkgo.sh \
			--race --timeout=$(E2E_TIMEOUT) --output-interceptor-mode=none  --slow-spec-threshold=15s \
			--json-report e2ereport.json --output-dir $(ROOT_DIR) --procs $(E2E_GINKGO_PROCS) \
			--label-filter="$(E2E_GINKGO_LABELS)" -randomize-suites -randomize-all  -vv --fail-fast  $(E2E_GINKGO_OPTION) \
			-r e2e/*  || RESULT=1  ; \
		echo "=========== after test `date` ===========" >> $(E2E_LOG_FILE) ; \
		./scripts/debugCluster.sh $(KIND_KUBECONFIG) "system" "$(E2E_LOG_FILE)" $(E2E_INSTALL_NAMESPACE) ; \
		KIND_CLUSTER_NAME=$(KIND_CLUSTER_NAME) ./scripts/debugCluster.sh $(KIND_KUBECONFIG) "detail" "$(E2E_LOG_FILE)" $(E2E_INSTALL_NAMESPACE) ; \
		./scripts/debugCluster.sh $(KIND_KUBECONFIG) "error" "$(E2E_LOG_FILE)" $(E2E_INSTALL_NAMESPACE) || { echo "error, found error log, datarace/pacni/longlock !!!" ; RESULT=1 ; } ; \
		if (($${RESULT} != 0)) ; then \
		   echo "failed to run e2e test"  ; \
		   exit 1 ; \
		fi ; \
		echo "" ; \
		echo "============================================" ; \
		echo "succeeded to run all test" ; \
		echo "output report to e2ereport.json" ; \
		echo "output env log to $(E2E_LOG_FILE) "
