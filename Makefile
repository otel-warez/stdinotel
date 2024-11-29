-include Makefile.Common

GO ?= go

Makefile.Common:
	@wget -q https://raw.githubusercontent.com/otel-warez/build-tools/refs/heads/main/Makefile.Common

internal/tools/empty_test.go:
	@mkdir -p internal/tools
	@wget -q https://raw.githubusercontent.com/otel-warez/build-tools/refs/heads/main/tools/empty_test.go -O internal/tools/empty_test.go

internal/tools/go.mod:
	@mkdir -p internal/tools
	@wget -q https://raw.githubusercontent.com/otel-warez/build-tools/refs/heads/main/tools/go.mod -O internal/tools/go.mod

internal/tools/go.sum:
	@mkdir -p internal/tools
	@wget -q https://raw.githubusercontent.com/otel-warez/build-tools/refs/heads/main/tools/go.sum -O internal/tools/go.sum

internal/tools/tools.go:
	@mkdir -p internal/tools
	@wget -q https://raw.githubusercontent.com/otel-warez/build-tools/refs/heads/main/tools/tools.go -O internal/tools/tools.go

.PHONY: setup
setup: internal/tools/empty_test.go internal/tools/go.mod internal/tools/go.sum internal/tools/tools.go Makefile.Common

# Build the Collector executable.
.PHONY: stdinotel
stdinotel: setup
	mkdir -p bin
	cd ./cmd/stdinotel && GOOS=linux GOARCH=arm64 GO111MODULE=on CGO_ENABLED=0 "${GO}" build -trimpath -o ../../bin/stdinotel_linux_arm64 .
	cd ./cmd/stdinotel && GOOS=linux GOARCH=amd64 GO111MODULE=on CGO_ENABLED=0 "${GO}" build -trimpath -o ../../bin/stdinotel_linux_amd64 .
	cd ./cmd/stdinotel && GOOS=darwin GOARCH=arm64 GO111MODULE=on CGO_ENABLED=0 "${GO}" build -trimpath -o ../../bin/stdinotel_darwin_arm64 .
	cd ./cmd/stdinotel && GOOS=darwin GOARCH=amd64 GO111MODULE=on CGO_ENABLED=0 "${GO}" build -trimpath -o ../../bin/stdinotel_darwin_amd64 .

