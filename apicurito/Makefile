CONFIG_DIR := ./config

# User customisable variables
include $(CONFIG_DIR)/vars/Makefile

ifeq (, $(shell command -v git 2> /dev/null))
GIT_COMMIT=$(shell git rev-list -1 HEAD)
else
GIT_COMMIT=release-$(VERSION)
endif

# kernel-style V=1 build verbosity
ifeq ("$(origin V)", "command line")
       BUILD_VERBOSE = $(V)
endif

ifeq ($(BUILD_VERBOSE),1)
       Q =
else
       Q = @
endif

#export CGO_ENABLED:=0

.PHONY: all
all: build

.PHONY: format
format:
	./scripts/go-fmt.sh

.PHONY: go-generate
go-generate: dep
	$(Q)go generate ./...

.PHONY: sdk-generate
sdk-generate: dep
	operator-sdk generate k8s

.PHONY: vet
vet:
	./scripts/go-vet.sh

.PHONY: test
test:
	./scripts/go-test.sh

.PHONY: lint
lint:
	# Temporarily disabled
	# ./scripts/go-lint.sh
	# ./scripts/yaml-lint.sh

.PHONY: build
build:
	@(./scripts/go-build.sh \
		-b $(BUILD_TIME) \
		-g $(GIT_COMMIT) \
		-i $(IMAGE) \
		-p $(PREVIOUS_VERSION) \
		-t $(TAG) \
		-v $(VERSION))

.PHONY: clean
clean:
	rm -rf build/_output

.PHONY: csv
csv:
	./scripts/go-csv.sh
