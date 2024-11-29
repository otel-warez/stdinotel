-include Makefile.Common

GO ?= go

.PHONY: setup
setup:
	wget


# Build the Collector executable.
.PHONY: stdinotel
stdinotel: setup
	mkdir -p bin
	cd ./cmd/stdinotel && GOOS=linux GOARCH=arm64 GO111MODULE=on CGO_ENABLED=0 "${GO}" build -trimpath -o ../../bin/stdinotel_linux_arm64 .
	cd ./cmd/stdinotel && GOOS=linux GOARCH=amd64 GO111MODULE=on CGO_ENABLED=0 "${GO}" build -trimpath -o ../../bin/stdinotel_linux_amd64 .
	cd ./cmd/stdinotel && GOOS=darwin GOARCH=arm64 GO111MODULE=on CGO_ENABLED=0 "${GO}" build -trimpath -o ../../bin/stdinotel_darwin_arm64 .
	cd ./cmd/stdinotel && GOOS=darwin GOARCH=amd64 GO111MODULE=on CGO_ENABLED=0 "${GO}" build -trimpath -o ../../bin/stdinotel_darwin_amd64 .

