CONTROLLER_GEN_VERSION := v0.6.1
OPERATOR_SDK_VERSION := v1.21.0
KUSTOMIZE_VERSION := v4.1.2
OPM_VERSION := v1.21.0
SRC_CATALOG := registry.redhat.io/redhat/redhat-operator-index:v4.11

PACKAGE := fuse-apicurito

DEFAULT_VERSION := 7.13.0
DEFAULT_PREVIOUS_VERSION := 7.12.0
DEFAULT_TAG := 1.13
DEFAULT_IMAGE := registry-proxy.engineering.redhat.com/rh-osbs/fuse7-fuse-apicurito-rhel8-operator
DEFAULT_NAMESPACE := $(shell $(CONFIG_DIR)/script/namespace.sh)
DEFAULT_CHANNEL := $(PACKAGE)-$(shell v=$(DEFAULT_VERSION) && echo "$${v%\.[0-9]}.x")
DEFAULT_CSV := manifests/bases/apicurito.clusterserviceversion.yaml
BUNDLE_INFIX := manifest

CSV_SUPPORT := Fuse Apicurito
CSV_DISPLAY_NAME := Red Hat Integration - API Designer

#
# Vars that can be overridden by external env vars
#
VERSION ?= $(DEFAULT_VERSION)
PREVIOUS_VERSION ?= $(DEFAULT_PREVIOUS_VERSION)
DEBUG ?= false
KUBE_USER ?= developer
IMAGE ?= $(DEFAULT_IMAGE)
TAG ?= $(DEFAULT_TAG)
LEGACY ?= false
NAMESPACE ?= $(DEFAULT_NAMESPACE)

CHANNELS ?= $(DEFAULT_CHANNEL)
# Default bundle image tag
BUNDLE_IMG ?= $(IMAGE)-$(BUNDLE_INFIX):$(TAG)
# Options for 'bundle-build'
ifneq ($(origin CHANNELS), undefined)
BUNDLE_CHANNELS := --channels=$(CHANNELS)
endif
ifneq ($(origin DEFAULT_CHANNEL), undefined)
BUNDLE_DEFAULT_CHANNEL := --default-channel=$(DEFAULT_CHANNEL)
endif
ifneq ($(origin PACKAGE), undefined)
BUNDLE_PACKAGE := --package=$(PACKAGE)
endif
BUNDLE_METADATA_OPTS ?= $(BUNDLE_CHANNELS) $(BUNDLE_DEFAULT_CHANNEL) $(BUNDLE_PACKAGE)

SHORT_VERSION := $(shell v=$(VERSION) && echo "$${v%.*}")
BUNDLE_DIR := bundle
PKG := ../pkg
CRD := crd/bases
MANAGER := manager
MANIFESTS := manifests
RBAC := rbac
TIMESTAMP=$(shell date -u '+%F %T %Z')
DOMAIN := apicur.io
CATALOG := catalog


CSV_NAME := $(PACKAGE).v$(VERSION)
CSV_REPLACES := $(PACKAGE).v$(PREVIOUS_VERSION)
CSV_FILENAME := $(PACKAGE).clusterserviceversion.yaml
CSV_PATH := $(MANIFESTS)/bases/$(CSV_FILENAME)

RELATED_IMAGES_PATCH := $(MANAGER)/patches/related-images-patch.yaml

#
# Timestamp for the building of the operator
#
BUILD_TIME := $(shell date +%Y-%m-%dT%H:%M:%S%z)

TMPL := tmpl
YAML := yaml
GEN := gen
NAMESPACE_VAR := {NAMESPACE}
TAG_VAR := {TAG}
IMAGE_VAR := {IMAGE}
VERSION_VAR := {VERSION}
SHORT_VERSION_VAR := {SHORT_VERSION}
