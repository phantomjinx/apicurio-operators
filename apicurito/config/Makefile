#
# Make Options
#
MK_OPTIONS := -s

#
# Allows for resources to be loaded from outside the root location of
# the kustomize config file. Ensures that resource don't need to be
# copied around the file system.
#
# See https://kubectl.docs.kubernetes.io/faq/kustomize
#
KOPTIONS := --load-restrictor LoadRestrictionsNone

CONFIG_DIR := .

# User customisable variables
include $(CONFIG_DIR)/vars/Makefile

.PHONY: kustomize setup operator app

kubectl:
ifeq (, $(shell which kubectl))
$(error "No kubectl found in PATH. Please install and re-run")
endif

#
# Setup the installation by installing crds, roles and granting
# privileges for the installing user.
#
setup: kubectl
	$(MAKE) $(MK_OPTIONS) -C $(RBAC) init
	#@ Must be invoked by a user with cluster-admin privileges
ifeq ($(LEGACY), true)
	@cd setup && \
		$(KUSTOMIZE) edit remove resource ../crd/bases && \
		$(KUSTOMIZE) edit add resource ../crd/deprecated
else
	@cd setup && \
		$(KUSTOMIZE) edit remove resource ../crd/deprecated && \
		$(KUSTOMIZE) edit add resource ../crd/bases
endif
ifeq ($(DEBUG), false)
	@$(KUSTOMIZE) build $(KOPTIONS) setup | kubectl apply -f -
else
	@$(KUSTOMIZE) build $(KOPTIONS) setup
endif

#
# Install the operator deployment and related resources
#
operator: kubectl
	$(MAKE) $(MK_OPTIONS) -C $(MANAGER) init
	#@ Can be invoked by a user with namespace privileges (rather than a cluster-admin)
ifeq ($(DEBUG), false)
	@$(KUSTOMIZE) build $(KOPTIONS) operator | kubectl apply -f -
else
	@$(KUSTOMIZE) build $(KOPTIONS) operator
endif

#
# Installs the operator deployment and in addition installs a default CR
#
app: kubectl
	$(MAKE) $(MK_OPTIONS) -C $(MANAGER) init
	#@ Can be invoked by a user with namespace privileges (rather than a cluster-admin)
ifeq ($(DEBUG), false)
	@$(KUSTOMIZE) build $(KOPTIONS) app | kubectl apply -f -
else
	@$(KUSTOMIZE) build $(KOPTIONS) app
endif


#
##################################
# B U N D L E  G E N E R A T I O N
##################################
#

# Produce CRDs that work back to Kubernetes 1.11 (no version conversion)
CRD_OPTIONS ?= "crd:trivialVersions=true"

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

# find or download controller-gen
# download controller-gen if necessary
controller-gen:
ifeq (, $(shell command -v controller-gen 2> /dev/null))
	@{ \
	set -e ;\
	CONTROLLER_GEN_TMP_DIR=$$(mktemp -d) ;\
	cd $$CONTROLLER_GEN_TMP_DIR ;\
	go mod init tmp ;\
	go get sigs.k8s.io/controller-tools/cmd/controller-gen@$(CONTROLLER_GEN_VERSION) ;\
	rm -rf $$CONTROLLER_GEN_TMP_DIR ;\
	}
CONTROLLER_GEN=$(GOBIN)/controller-gen
else
CONTROLLER_GEN=$(shell command -v controller-gen 2> /dev/null)
endif

kustomize:
ifeq (, $(shell command -v kustomize 2> /dev/null))
	@{ \
	set -e ;\
	KUSTOMIZE_GEN_TMP_DIR=$$(mktemp -d) ;\
	cd $$KUSTOMIZE_GEN_TMP_DIR ;\
	go mod init tmp ;\
	go get sigs.k8s.io/kustomize/kustomize/v4@$(KUSTOMIZE_VERSION) ;\
	rm -rf $$KUSTOMIZE_GEN_TMP_DIR ;\
	}
KUSTOMIZE=$(GOBIN)/kustomize
else
KUSTOMIZE=$(shell command -v kustomize 2> /dev/null)
endif

opm:
ifeq (, $(shell command -v opm 2> /dev/null))
	@{ \
	set -e ;\
	OPM_GEN_TMP_DIR=$$(mktemp -d) ;\
	cd $$OPM_GEN_TMP_DIR ;\
	go mod init tmp ;\
	go get \
		-ldflags '-w -extldflags "-static"' -tags "json1" \
		-ldflags "-X '$(OPM_PKG)/cmd/opm/version.opmVersion=$(OPM_VERSION)' -X '$(OPM_PKG)/cmd/opm/version.buildDate=$(BUILD_TIME)'" \
		$(OPM_PKG)/cmd/opm@$(OPM_VERSION) ;\
	rm -rf $$OPM_GEN_TMP_DIR ;\
	}
OPM=$(GOBIN)/opm
else
OPM=$(shell command -v opm 2> /dev/null)
endif

operator-sdk:
ifeq (, $(shell command -v operator-sdk 2> /dev/null))
	@{ \
	set -e ;\
	curl \
		-L https://github.com/operator-framework/operator-sdk/releases/download/$(OPERATOR_SDK_VERSION)/operator-sdk_linux_amd64 \
		-o operator-sdk ;\
	chmod +x operator-sdk ;\
	mv operator-sdk $(GOBIN)/ ;\
	}
OPERATOR_SDK=$(GOBIN)/operator-sdk
else
OPERATOR_SDK=$(shell command -v operator-sdk 2> /dev/null)
endif

.PHONY: kustomize operator-sdk bundle manifests generate generate-deepcopy generate-crds

generate-deepcopy: controller-gen
	@$(CONTROLLER_GEN) paths="$(PKG)/..." object

generate-crds: controller-gen
	@$(CONTROLLER_GEN) crd crd:crdVersions=v1 \
		paths=$(PKG)/apis/... \
		output:artifacts:config=$(CRD) \
		output:crd:dir=$(CRD)
	@sed -i '/creationTimestamp: null/d' $(CRD)/$(DOMAIN)_*.yaml

# Generate code
generate: generate-deepcopy generate-crds

# Tailor the manifest according to default values for this project
# Note. to make the bundle this name must match that specified in PROJECT
#
pre-bundle:
# bundle name must match that which appears in PROJECT file
	@sed -i 's/projectName: .*/projectName: $(PACKAGE)/' PROJECT
# finds the single CSV file and renames it
	@find $(MANIFESTS)/bases -type f -name "*.clusterserviceversion.yaml" -execdir mv '{}' $(CSV_FILENAME) ';'
	@sed -i 's~^    containerImage: .*~    containerImage: $(IMAGE):$(TAG)~' $(CSV_PATH)
	@sed -i 's/^    support: .*/    support: $(CSV_SUPPORT)/' $(CSV_PATH)
	@sed -i "s/^    olm.skipRange: .*/    olm.skipRange: '>=7.7.0 <$(VERSION)'/" $(CSV_PATH)
	@sed -i 's/^  name: .*.\(v.*\)/  name: $(CSV_NAME)/' $(CSV_PATH)
	@sed -i 's/^  displayName: .*/  displayName: $(CSV_DISPLAY_NAME)/' $(CSV_PATH)
	@sed -i 's/^  replaces: .*/  replaces: $(CSV_REPLACES)/' $(CSV_PATH)
	@sed -i 's/^  version: .*/  version: $(VERSION)/' $(CSV_PATH)

# Generate bundle manifests and metadata, then validate generated files.
bundle: generate-crds kustomize pre-bundle
	@$(MAKE) $(MK_OPTIONS) -C $(RBAC) init
	@$(MAKE) $(MK_OPTIONS) -C $(MANAGER) init
# (Re)-generate the manifests directory
	@$(OPERATOR_SDK) generate kustomize manifests \
		--apis-dir $(PKG)/apis \
		--input-dir $(MANIFESTS) \
		--output-dir $(MANIFESTS)
# Sets the operator image to the preferred image:tag
	@cd $(MANIFESTS) && $(KUSTOMIZE) edit set image $(DEFAULT_IMAGE)=$(IMAGE):$(TAG)
# Generates the bundle complete with manifests
	@$(KUSTOMIZE) build $(KOPTIONS) $(MANIFESTS) | \
		$(OPERATOR_SDK) generate bundle \
		-q --overwrite --version $(VERSION) \
		--kustomize-dir $(MANIFESTS) $(BUNDLE_METADATA_OPTS)
# Add a timestamp to the new bundle manifest
	@sed -i 's/    createdAt: .*/    createdAt: $(TIMESTAMP)/' $(BUNDLE_DIR)/$(MANIFESTS)/$(CSV_FILENAME)
# Moves the docker file into the bundle directory
	@mv bundle.Dockerfile $(BUNDLE_DIR)/Dockerfile && \
		sed -i 's/bundle\///g' $(BUNDLE_DIR)/Dockerfile
# Add Product LABELS to Dockerfile
	@sed -i 's/version=\".*\"/version=\"$(TAG)\"/' Dockerfile.labels
	@cat Dockerfile.labels >> $(BUNDLE_DIR)/Dockerfile
	@$(OPERATOR_SDK) bundle validate $(BUNDLE_DIR)

# Build the bundle image.
.PHONY: bundle-build
bundle-build: bundle
	cd bundle && docker build -f Dockerfile -t $(BUNDLE_IMG) .

.PHONY: bundle-push
bundle-push: bundle-build
	docker push $(BUNDLE_IMG)

.PHONY: catalog
catalog: opm bundle-build
	SRC_CATALOG=$(SRC_CATALOG) CATALOG_DIR=$(CATALOG) \
	OPM=$(OPM) BUNDLE_IMAGE=$(BUNDLE_IMG) CSV_NAME=$(CSV_NAME) \
	CSV_REPLACES=$(CSV_REPLACES) CHANNEL=$(DEFAULT_CHANNEL) \
	PACKAGE=$(PACKAGE) ./script/build-catalog.sh

clean:
	@find . -name "*.gen.*" -delete
ifneq ($(CSV_PATH), $(DEFAULT_CSV))
ifneq ("$(wildcard $(CSV_PATH))","")
	@mv $(CSV_PATH) $(DEFAULT_CSV)
endif
endif
	@rm -rf bundle
