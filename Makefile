include Makefile.Common

OTELCOL_BUILDER_VERSION ?= 0.96.0
OTELCOL_BUILDER_DIR ?= ${HOME}/bin
OTELCOL_BUILDER ?= ${OTELCOL_BUILDER_DIR}/ocb
GO ?= go
BUILDER ?= ocb

.PHONY: install-tools
install-tools: $(TOOLS_BIN_NAMES)

$(TOOLS_BIN_DIR):
	mkdir -p $@

$(TOOLS_BIN_NAMES): $(TOOLS_BIN_DIR) $(TOOLS_MOD_DIR)/go.mod
	cd $(TOOLS_MOD_DIR) && $(GOCMD) build -o $@ -trimpath $(filter %/$(notdir $@),$(TOOLS_PKG_NAMES))

.PHONY: ocb
ocb:
ifeq (, $(shell command -v ocb 2>/dev/null))
	@{ \
	[ ! -x '$(OTELCOL_BUILDER)' ] || exit 0; \
	set -e ;\
	os=$$(uname | tr A-Z a-z) ;\
	machine=$$(uname -m) ;\
	[ "$${machine}" != x86 ] || machine=386 ;\
	[ "$${machine}" != x86_64 ] || machine=amd64 ;\
	echo "Installing ocb ($${os}/$${machine}) at $(OTELCOL_BUILDER_DIR)";\
	mkdir -p $(OTELCOL_BUILDER_DIR) ;\
	curl -sfLo $(OTELCOL_BUILDER) "https://github.com/open-telemetry/opentelemetry-collector/releases/download/cmd%2Fbuilder%2Fv$(OTELCOL_BUILDER_VERSION)/ocb_$(OTELCOL_BUILDER_VERSION)_$${os}_$${machine}" ;\
	chmod +x $(OTELCOL_BUILDER) ;\
	}
else
OTELCOL_BUILDER=$(shell command -v ocb)
endif

# Build the Collector executable.
.PHONY: stdinotel
stdinotel: ocb
	mkdir -p bin
	cd ./cmd/stdinotel && mkdir -p _build
	cd ./cmd/stdinotel && "${OTELCOL_BUILDER}" --skip-compilation=true --config manifest.yaml
	cd ./cmd/stdinotel/_build && GOOS=linux GOARCH=arm64 GO111MODULE=on CGO_ENABLED=0 "${GO}" build -trimpath -o ../../../bin/stdinotel_linux_arm64 .
	cd ./cmd/stdinotel/_build && GOOS=linux GOARCH=amd64 GO111MODULE=on CGO_ENABLED=0 "${GO}" build -trimpath -o ../../../bin/stdinotel_linux_amd64 .

.PHONY: lint
lint: $(LINT) checklicense misspell
	$(LINT) run --allow-parallel-runners --build-tags integration

.PHONY: tidy
tidy:
	rm -fr go.sum
	$(GOCMD) mod tidy -compat=1.19
