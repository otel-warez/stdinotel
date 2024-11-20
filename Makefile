include Makefile.Common

GO ?= go

.PHONY: install-tools
install-tools: $(TOOLS_BIN_NAMES)

$(TOOLS_BIN_DIR):
	mkdir -p $@

$(TOOLS_BIN_NAMES): $(TOOLS_BIN_DIR) $(TOOLS_MOD_DIR)/go.mod
	cd $(TOOLS_MOD_DIR) && $(GOCMD) build -o $@ -trimpath $(filter %/$(notdir $@),$(TOOLS_PKG_NAMES))

# Build the Collector executable.
.PHONY: stdinotel
stdinotel:
	mkdir -p bin
	cd ./cmd/stdinotel && GOOS=linux GOARCH=arm64 GO111MODULE=on CGO_ENABLED=0 "${GO}" build -trimpath -o ../../bin/stdinotel_linux_arm64 .
	cd ./cmd/stdinotel && GOOS=linux GOARCH=amd64 GO111MODULE=on CGO_ENABLED=0 "${GO}" build -trimpath -o ../../bin/stdinotel_linux_amd64 .

.PHONY: lint
lint: $(LINT) checklicense misspell
	$(LINT) run --allow-parallel-runners --build-tags integration

.PHONY: tidy
tidy:
	rm -fr go.sum
	$(GOCMD) mod tidy -compat=1.23
