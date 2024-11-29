-include Makefile.Common

GO ?= go

.PHONY: setup
setup: Makefile.Common internal/tools
	wget https://raw.githubusercontent.com/otel-warez/build-tools/refs/heads/main/Makefile.Common
	mkdir -p internal/tools
	wget https://raw.githubusercontent.com/otel-warez/build-tools/refs/heads/main/tools/Makefile internal/tools/Makefile
	wget https://raw.githubusercontent.com/otel-warez/build-tools/refs/heads/main/tools/empty_test.go internal/tools/empty_test.go
	wget https://raw.githubusercontent.com/otel-warez/build-tools/refs/heads/main/tools/go.mod internal/tools/go.mod
	wget https://raw.githubusercontent.com/otel-warez/build-tools/refs/heads/main/tools/go.sum internal/tools/go.sum
	wget https://raw.githubusercontent.com/otel-warez/build-tools/refs/heads/main/tools/tools.go internal/tools/tools.go

# Build the Collector executable.
.PHONY: stdinotel
stdinotel: setup
	mkdir -p bin
	cd ./cmd/stdinotel && GOOS=linux GOARCH=arm64 GO111MODULE=on CGO_ENABLED=0 "${GO}" build -trimpath -o ../../bin/stdinotel_linux_arm64 .
	cd ./cmd/stdinotel && GOOS=linux GOARCH=amd64 GO111MODULE=on CGO_ENABLED=0 "${GO}" build -trimpath -o ../../bin/stdinotel_linux_amd64 .
	cd ./cmd/stdinotel && GOOS=darwin GOARCH=arm64 GO111MODULE=on CGO_ENABLED=0 "${GO}" build -trimpath -o ../../bin/stdinotel_darwin_arm64 .
	cd ./cmd/stdinotel && GOOS=darwin GOARCH=amd64 GO111MODULE=on CGO_ENABLED=0 "${GO}" build -trimpath -o ../../bin/stdinotel_darwin_amd64 .

