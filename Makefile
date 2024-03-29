include ./Makefile.Common

RUN_CONFIG?=local/config.yaml
CMD?=

VERSION=$(shell git describe --always --match "v[0-9]*" HEAD)
COREVERSION="0.97.0"

COMP_REL_PATH=cmd/otelcontribcol/components.go
MOD_NAME=github.com/open-telemetry/opentelemetry-collector-contrib

GROUP ?= all
FOR_GROUP_TARGET=for-$(GROUP)-target

FIND_MOD_ARGS=-type f -name "go.mod"
TO_MOD_DIR=dirname {} \; | sort | grep -E '^./'
EX_COMPONENTS=-not -path "./receiver/*" -not -path "./processor/*" -not -path "./exporter/*" -not -path "./extension/*" -not -path "./connector/*"
EX_INTERNAL=-not -path "./internal/*"
EX_PKG=-not -path "./pkg/*"
EX_CMD=-not -path "./cmd/*"

# NONROOT_MODS includes ./* dirs (excludes . dir)
NONROOT_MODS := $(shell find . $(FIND_MOD_ARGS) -exec $(TO_MOD_DIR) )

RECEIVER_MODS_0 := $(shell find ./receiver/[a-f]* $(FIND_MOD_ARGS) -exec $(TO_MOD_DIR) )
RECEIVER_MODS_1 := $(shell find ./receiver/[g-o]* $(FIND_MOD_ARGS) -exec $(TO_MOD_DIR) )
RECEIVER_MODS_2 := $(shell find ./receiver/[p]* $(FIND_MOD_ARGS) -exec $(TO_MOD_DIR) ) # Prometheus is special and gets its own section.
RECEIVER_MODS_3 := $(shell find ./receiver/[q-z]* $(FIND_MOD_ARGS) -exec $(TO_MOD_DIR) )
RECEIVER_MODS := $(RECEIVER_MODS_0) $(RECEIVER_MODS_1) $(RECEIVER_MODS_2) $(RECEIVER_MODS_3)
PROCESSOR_MODS := $(shell find ./processor/* $(FIND_MOD_ARGS) -exec $(TO_MOD_DIR) )
EXPORTER_MODS_0 := $(shell find ./exporter/[a-m]* $(FIND_MOD_ARGS) -exec $(TO_MOD_DIR) )
EXPORTER_MODS_1 := $(shell find ./exporter/[n-z]* $(FIND_MOD_ARGS) -exec $(TO_MOD_DIR) )
EXPORTER_MODS := $(EXPORTER_MODS_0) $(EXPORTER_MODS_1)
EXTENSION_MODS := $(shell find ./extension/* $(FIND_MOD_ARGS) -exec $(TO_MOD_DIR) )
CONNECTOR_MODS := $(shell find ./connector/* $(FIND_MOD_ARGS) -exec $(TO_MOD_DIR) )
INTERNAL_MODS := $(shell find ./internal/* $(FIND_MOD_ARGS) -exec $(TO_MOD_DIR) )
PKG_MODS := $(shell find ./pkg/* $(FIND_MOD_ARGS) -exec $(TO_MOD_DIR) )
CMD_MODS_0 := $(shell find ./cmd/[a-m]* $(FIND_MOD_ARGS) -exec $(TO_MOD_DIR) )
CMD_MODS_1 := $(shell find ./cmd/[n-z]* $(FIND_MOD_ARGS) -not -path "./cmd/otelcontribcol/*" -exec $(TO_MOD_DIR) )
CMD_MODS := $(CMD_MODS_0) $(CMD_MODS_1)
OTHER_MODS := $(shell find . $(EX_COMPONENTS) $(EX_INTERNAL) $(EX_PKG) $(EX_CMD) $(FIND_MOD_ARGS) -exec $(TO_MOD_DIR) ) $(PWD)
ALL_MODS := $(RECEIVER_MODS) $(PROCESSOR_MODS) $(EXPORTER_MODS) $(EXTENSION_MODS) $(CONNECTOR_MODS) $(INTERNAL_MODS) $(PKG_MODS) $(CMD_MODS) $(OTHER_MODS)

FIND_INTEGRATION_TEST_MODS={ find . -type f -name "*integration_test.go" & find . -type f -name "*e2e_test.go" -not -path "./testbed/*"; }
INTEGRATION_MODS := $(shell $(FIND_INTEGRATION_TEST_MODS) | xargs $(TO_MOD_DIR) | uniq)

ifeq ($(GOOS),windows)
	EXTENSION := .exe
endif

.DEFAULT_GOAL := all

all-modules:
	@echo $(NONROOT_MODS) | tr ' ' '\n' | sort

all-groups:
	@echo "receiver-0: $(RECEIVER_MODS_0)"
	@echo "\nreceiver-1: $(RECEIVER_MODS_1)"
	@echo "\nreceiver-2: $(RECEIVER_MODS_2)"
	@echo "\nreceiver-3: $(RECEIVER_MODS_3)"
	@echo "\nreceiver: $(RECEIVER_MODS)"
	@echo "\nprocessor: $(PROCESSOR_MODS)"
	@echo "\nexporter-0: $(EXPORTER_MODS_0)"
	@echo "\nexporter-1: $(EXPORTER_MODS_1)"
	@echo "\nextension: $(EXTENSION_MODS)"
	@echo "\nconnector: $(CONNECTOR_MODS)"
	@echo "\ninternal: $(INTERNAL_MODS)"
	@echo "\npkg: $(PKG_MODS)"
	@echo "\ncmd-0: $(CMD_MODS_0)"
	@echo "\ncmd-1: $(CMD_MODS_1)"
	@echo "\nother: $(OTHER_MODS)"

.PHONY: all
all: install-tools all-common goporto multimod-verify gotest otelcontribcol

.PHONY: all-common
all-common:
	@$(MAKE) $(FOR_GROUP_TARGET) TARGET="common"

.PHONY: e2e-test
e2e-test: otelcontribcol oteltestbedcol
	$(MAKE) --no-print-directory -C testbed run-tests

.PHONY: integration-test
integration-test:
	@$(MAKE) for-integration-target TARGET="mod-integration-test"

.PHONY: integration-tests-with-cover
integration-tests-with-cover:
	@$(MAKE) for-integration-target TARGET="do-integration-tests-with-cover"

# Long-running e2e tests
.PHONY: stability-tests
stability-tests: otelcontribcol
	@echo Stability tests are disabled until we have a stable performance environment.
	@echo To enable the tests replace this echo by $(MAKE) -C testbed run-stability-tests

.PHONY: gogci
gogci:
	$(MAKE) $(FOR_GROUP_TARGET) TARGET="gci"

.PHONY: gotidy
gotidy:
	$(MAKE) $(FOR_GROUP_TARGET) TARGET="tidy"

.PHONY: gomoddownload
gomoddownload:
	$(MAKE) $(FOR_GROUP_TARGET) TARGET="moddownload"

.PHONY: gotest
gotest:
	$(MAKE) $(FOR_GROUP_TARGET) TARGET="test"

.PHONY: gotest-with-cover
gotest-with-cover:
	@$(MAKE) $(FOR_GROUP_TARGET) TARGET="test-with-cover"
	$(GOCMD) tool covdata textfmt -i=./coverage/unit -o ./$(GROUP)-coverage.txt

.PHONY: gointegration-test
gointegration-test:
	$(MAKE) $(FOR_GROUP_TARGET) TARGET="mod-integration-test"

.PHONY: gofmt
gofmt:
	$(MAKE) $(FOR_GROUP_TARGET) TARGET="fmt"

.PHONY: golint
golint:
	$(MAKE) $(FOR_GROUP_TARGET) TARGET="lint"

.PHONY: gogovulncheck
gogovulncheck:
	$(MAKE) $(FOR_GROUP_TARGET) TARGET="govulncheck"

.PHONY: goporto
goporto: $(PORTO)
	$(PORTO) -w --include-internal --skip-dirs "^cmd$$" ./

.PHONY: for-all
for-all:
	@echo "running $${CMD} in root"
	@$${CMD}
	@set -e; for dir in $(NONROOT_MODS); do \
	  (cd "$${dir}" && \
	  	echo "running $${CMD} in $${dir}" && \
	 	$${CMD} ); \
	done

COMMIT?=HEAD
MODSET?=contrib-core
REMOTE?=git@github.com:open-telemetry/opentelemetry-collector-contrib.git
.PHONY: push-tags
push-tags: $(MULTIMOD)
	$(MULTIMOD) verify
	set -e; for tag in `$(MULTIMOD) tag -m ${MODSET} -c ${COMMIT} --print-tags | grep -v "Using" `; do \
		echo "pushing tag $${tag}"; \
		git push ${REMOTE} $${tag}; \
	done;

# Define a delegation target for each module
.PHONY: $(ALL_MODS)
$(ALL_MODS):
	@echo "Running target '$(TARGET)' in module '$@' as part of group '$(GROUP)'"
	$(MAKE) --no-print-directory -C $@ $(TARGET)

# Trigger each module's delegation target
.PHONY: for-all-target
for-all-target: $(ALL_MODS)

.PHONY: for-receiver-target
for-receiver-target: $(RECEIVER_MODS)

.PHONY: for-receiver-0-target
for-receiver-0-target: $(RECEIVER_MODS_0)

.PHONY: for-receiver-1-target
for-receiver-1-target: $(RECEIVER_MODS_1)

.PHONY: for-receiver-2-target
for-receiver-2-target: $(RECEIVER_MODS_2)

.PHONY: for-receiver-3-target
for-receiver-3-target: $(RECEIVER_MODS_3)

.PHONY: for-processor-target
for-processor-target: $(PROCESSOR_MODS)

.PHONY: for-exporter-target
for-exporter-target: $(EXPORTER_MODS)

.PHONY: for-exporter-0-target
for-exporter-0-target: $(EXPORTER_MODS_0)

.PHONY: for-exporter-1-target
for-exporter-1-target: $(EXPORTER_MODS_1)

.PHONY: for-extension-target
for-extension-target: $(EXTENSION_MODS)

.PHONY: for-connector-target
for-connector-target: $(CONNECTOR_MODS)

.PHONY: for-internal-target
for-internal-target: $(INTERNAL_MODS)

.PHONY: for-pkg-target
for-pkg-target: $(PKG_MODS)

.PHONY: for-cmd-target
for-cmd-target: $(CMD_MODS)

.PHONY: for-cmd-0-target
for-cmd-0-target: $(CMD_MODS_0)

.PHONY: for-cmd-1-target
for-cmd-1-target: $(CMD_MODS_1)

.PHONY: for-other-target
for-other-target: $(OTHER_MODS)

.PHONY: for-integration-target
for-integration-target: $(INTEGRATION_MODS)

# Debugging target, which helps to quickly determine whether for-all-target is working or not.
.PHONY: all-pwd
all-pwd:
	$(MAKE) $(FOR_GROUP_TARGET) TARGET="pwd"

.PHONY: run
run:
	cd ./cmd/otelcontribcol && GO111MODULE=on $(GOCMD) run --race . --config ../../${RUN_CONFIG} ${RUN_ARGS}

.PHONY: docker-component # Not intended to be used directly
docker-component: check-component
	GOOS=linux GOARCH=amd64 $(MAKE) $(COMPONENT)
	cp ./bin/$(COMPONENT)_linux_amd64 ./cmd/$(COMPONENT)/$(COMPONENT)
	docker build -t $(COMPONENT) ./cmd/$(COMPONENT)/
	rm ./cmd/$(COMPONENT)/$(COMPONENT)

.PHONY: check-component
check-component:
ifndef COMPONENT
	$(error COMPONENT variable was not defined)
endif

.PHONY: docker-otelcontribcol
docker-otelcontribcol:
	COMPONENT=otelcontribcol $(MAKE) docker-component

.PHONY: docker-telemetrygen
docker-telemetrygen:
	GOOS=linux GOARCH=$(GOARCH) $(MAKE) telemetrygen
	cp bin/telemetrygen_* cmd/telemetrygen/
	cd cmd/telemetrygen && docker build --platform linux/$(GOARCH) -t telemetrygen:latest .
	rm cmd/telemetrygen/telemetrygen_*

.PHONY: generate
generate: install-tools
	cd ./internal/tools && go install go.opentelemetry.io/collector/cmd/mdatagen
	$(MAKE) for-all CMD="$(GOCMD) generate ./..."

.PHONY: githubgen-install
githubgen-install:
	cd cmd/githubgen && $(GOCMD) install .

.PHONY: gengithub
gengithub: githubgen-install
	githubgen

.PHONY: gendistributions
gendistributions: githubgen-install
	githubgen distributions

.PHONY: update-codeowners
update-codeowners: gengithub generate

FILENAME?=$(shell git branch --show-current)
.PHONY: chlog-new
chlog-new: $(CHLOGGEN)
	$(CHLOGGEN) new --config $(CHLOGGEN_CONFIG) --filename $(FILENAME)

.PHONY: chlog-validate
chlog-validate: $(CHLOGGEN)
	$(CHLOGGEN) validate --config $(CHLOGGEN_CONFIG)

.PHONY: chlog-preview
chlog-preview: $(CHLOGGEN)
	$(CHLOGGEN) update --config $(CHLOGGEN_CONFIG) --dry

.PHONY: chlog-update
chlog-update: $(CHLOGGEN)
	$(CHLOGGEN) update --config $(CHLOGGEN_CONFIG) --version $(VERSION)

.PHONY: genotelcontribcol
genotelcontribcol: $(BUILDER)
	$(BUILDER) --skip-compilation --config cmd/otelcontribcol/builder-config.yaml --output-path cmd/otelcontribcol
	$(MAKE) --no-print-directory -C cmd/otelcontribcol fmt

# Build the Collector executable.
.PHONY: otelcontribcol
otelcontribcol:
	cd ./cmd/otelcontribcol && GO111MODULE=on CGO_ENABLED=0 $(GOCMD) build -trimpath -o ../../bin/otelcontribcol_$(GOOS)_$(GOARCH)$(EXTENSION) \
		-tags $(GO_BUILD_TAGS) .

.PHONY: genoteltestbedcol
genoteltestbedcol: $(BUILDER)
	$(BUILDER) --skip-compilation --config cmd/oteltestbedcol/builder-config.yaml --output-path cmd/oteltestbedcol
	$(MAKE) --no-print-directory -C cmd/oteltestbedcol fmt

# Build the Collector executable, with only components used in testbed.
.PHONY: oteltestbedcol
oteltestbedcol:
	cd ./cmd/oteltestbedcol && GO111MODULE=on CGO_ENABLED=0 $(GOCMD) build -trimpath -o ../../bin/oteltestbedcol_$(GOOS)_$(GOARCH)$(EXTENSION) \
		-tags $(GO_BUILD_TAGS) .

# Build the telemetrygen executable.
.PHONY: telemetrygen
telemetrygen:
	cd ./cmd/telemetrygen && GO111MODULE=on CGO_ENABLED=0 $(GOCMD) build -trimpath -o ../../bin/telemetrygen_$(GOOS)_$(GOARCH)$(EXTENSION) \
		-tags $(GO_BUILD_TAGS) .

.PHONY: update-otel
update-otel:$(MULTIMOD)
	$(MULTIMOD) sync -s=true -o ../opentelemetry-collector-contrib -m stable 
	git add . && git commit -s -m "[chore] multimod update stable modules" ; \
	$(MULTIMOD) sync -s=true -o ../opentelemetry-collector-contrib -m beta 
	git add . && git commit -s -m "[chore] multimod update beta modules" ; \
	sed -i -E -e "s/v[0-9]+\.[0-9]+\.[0-9].*$\/v$(COREVERSION)/" -e "5,6s/[0-9]+\.[0-9]+\.[0-9]/$(COREVERSION)/" ./cmd/otelcontribcol/builder-config.yaml 
	sed -i -E -e "s/v[0-9]+\.[0-9]+\.[0-9].*$\/v$(COREVERSION)/" -e "5,6s/[0-9]+\.[0-9]+\.[0-9]/$(COREVERSION)/" ./cmd/oteltestbedcol/builder-config.yaml 
	$(MAKE) gotidy
	$(MAKE) genotelcontribcol
	$(MAKE) genoteltestbedcol
	$(MAKE) oteltestbedcol

.PHONY: otel-from-tree
otel-from-tree:
	# This command allows you to make changes to your local checkout of otel core and build
	# contrib against those changes without having to push to github and update a bunch of
	# references. The workflow is:
	#
	# 1. Hack on changes in core (assumed to be checked out in ../opentelemetry-collector from this directory)
	# 2. Run `make otel-from-tree` (only need to run it once to remap go modules)
	# 3. You can now build contrib and it will use your local otel core changes.
	# 4. Before committing/pushing your contrib changes, undo by running `make otel-from-lib`.
	$(MAKE) for-all CMD="$(GOCMD) mod edit -replace go.opentelemetry.io/collector=$(SRC_PARENT_DIR)/opentelemetry-collector"

.PHONY: otel-from-lib
otel-from-lib:
	# Sets opentelemetry core to be not be pulled from local source tree. (Undoes otel-from-tree.)
	$(MAKE) for-all CMD="$(GOCMD) mod edit -dropreplace go.opentelemetry.io/collector"

.PHONY: build-examples
build-examples:
	docker-compose -f examples/demo/docker-compose.yaml build
	cd examples/secure-tracing/certs && $(MAKE) clean && $(MAKE) all && docker-compose -f ../docker-compose.yaml build
	docker-compose -f exporter/splunkhecexporter/example/docker-compose.yml build

.PHONY: deb-rpm-package
%-package: ARCH ?= amd64
%-package:
	GOOS=linux GOARCH=$(ARCH) $(MAKE) otelcontribcol
	docker build -t otelcontribcol-fpm internal/buildscripts/packaging/fpm
	docker run --rm -v $(CURDIR):/repo -e PACKAGE=$* -e VERSION=$(VERSION) -e ARCH=$(ARCH) otelcontribcol-fpm

# Verify existence of READMEs for components specified as default components in the collector.
.PHONY: checkdoc
checkdoc: $(CHECKFILE)
	$(CHECKFILE) --project-path $(CURDIR) --component-rel-path $(COMP_REL_PATH) --module-name $(MOD_NAME) --file-name "README.md"

# Verify existence of metadata.yaml for components specified as default components in the collector.
.PHONY: checkmetadata
checkmetadata: $(CHECKFILE)
	$(CHECKFILE) --project-path $(CURDIR) --component-rel-path $(COMP_REL_PATH) --module-name $(MOD_NAME) --file-name "metadata.yaml"

.PHONY: checkapi
checkapi:
	$(GOCMD) run cmd/checkapi/main.go .

.PHONY: all-checklinks
all-checklinks:
	$(MAKE) $(FOR_GROUP_TARGET) TARGET="checklinks"

# Function to execute a command. Note the empty line before endef to make sure each command
# gets executed separately instead of concatenated with previous one.
# Accepts command to execute as first parameter.
define exec-command
$(1)

endef

# List of directories where certificates are stored for unit tests.
CERT_DIRS := receiver/sapmreceiver/testdata \
             receiver/signalfxreceiver/testdata \
             receiver/splunkhecreceiver/testdata \
             receiver/mongodbatlasreceiver/testdata/alerts/cert \
             receiver/mongodbreceiver/testdata/certs \
             receiver/cloudflarereceiver/testdata/cert

# Generate certificates for unit tests relying on certificates.
.PHONY: certs
certs:
	$(foreach dir, $(CERT_DIRS), $(call exec-command, @internal/buildscripts/gen-certs.sh -o $(dir)))

.PHONY: multimod-verify
multimod-verify: $(MULTIMOD)
	@echo "Validating versions.yaml"
	$(MULTIMOD) verify

.PHONY: multimod-prerelease
multimod-prerelease: $(MULTIMOD)
	$(MULTIMOD) prerelease -s=true -b=false -v ./versions.yaml -m contrib-base
	$(MAKE) gotidy

.PHONY: multimod-sync
multimod-sync: $(MULTIMOD)
	$(MULTIMOD) sync -a=true -s=true -o ../opentelemetry-collector
	$(MAKE) gotidy

.PHONY: crosslink
crosslink: $(CROSSLINK)
	@echo "Executing crosslink"
	$(CROSSLINK) --root=$(shell pwd) --prune

.PHONY: clean
clean:
	@echo "Removing coverage files"
	find . -type f -name 'coverage.txt' -delete
	find . -type f -name 'coverage.html' -delete
	find . -type f -name 'coverage.out' -delete
	find . -type f -name 'integration-coverage.txt' -delete
	find . -type f -name 'integration-coverage.html' -delete

.PHONY: genconfigdocs
genconfigdocs:
	cd cmd/configschema && $(GOCMD) run ./docsgen all

.PHONY: generate-gh-issue-templates
generate-gh-issue-templates:
	cd cmd/githubgen && $(GOCMD) install .
	githubgen issue-templates

.PHONY: checks
checks:
	$(MAKE) checkdoc
	$(MAKE) checkmetadata
	$(MAKE) checkapi
	$(MAKE) -j4 goporto
	$(MAKE) crosslink
	$(MAKE) -j4 gotidy
	$(MAKE) genotelcontribcol
	$(MAKE) genoteltestbedcol
	$(MAKE) gendistributions
	$(MAKE) -j4 generate
	$(MAKE) multimod-verify
	git diff --exit-code || (echo 'Some files need committing' &&  git status && exit 1)
