APPNAME=serverapp

.PHONY: build
build:
	go build -mod=vendor -o serverapp main.go

.PHONY: test
test:
	go test ./...

.PHONY: all
all: build

.PHONY: oc-build
oc-build: oc cluster-build-config cluster-build

.PHONY: cluster-build-config
cluster-build-config: oc
	$(OC) apply -f config/openshift_build.yaml

.PHONY: cluster-build
cluster-build: oc
	$(OC) start-build $(APPNAME)

.PHONY: deploy
deploy: oc apply-deployment expose-svc

.PHONY: apply-deployment
apply-deployment: oc
	$(OC) apply -f config/openshift_deployment.yaml

.PHONY: expose-svc
expose-svc: oc
	$(OC) expose svc/serverapp

.PHONY: oc
OC = ./bin/oc
oc: ## Download oc locally if necessary
## You may need to change this for your OS and arch
# Linux x86: https://mirror.openshift.com/pub/openshift-v4/x86_64/clients/ocp/stable/openshift-client-linux.tar.gz
# Mac x86: https://mirror.openshift.com/pub/openshift-v4/x86_64/clients/ocp/stable/openshift-client-mac.tar.gz
# Mac ARM: (m-series): https://mirror.openshift.com/pub/openshift-v4/aarch64/clients/ocp/stable/openshift-client-mac-arm64.tar.gz
# Windows: https://mirror.openshift.com/pub/openshift-v4/x86_64/clients/ocp/stable/openshift-client-windows.zip
ifeq (,$(wildcard $(OC)))
ifeq (,$(shell which oc 2>/dev/null))
	@{ \
	set -e ;\
	mkdir -p $(dir $(OC)) ;\
	curl -sSLo openshift-client-linux.tar.gz https://mirror.openshift.com/pub/openshift-v4/x86_64/clients/ocp/stable/openshift-client-linux.tar.gz ;\
	tar -xvf openshift-client-linux.tar.gz -C bin && rm -rf openshift-client-linux.tar.gz ;\
	chmod +x bin/oc bin/kubectl ;\
	}
else
OC = $(shell which oc)
endif
endif

