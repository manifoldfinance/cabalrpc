BINDIR      := $(CURDIR)/bin
DIST_DIRS   := find * -type d -exec
TARGETS     := darwin/amd64 linux/amd64 linux/386 linux/arm linux/arm64 linux/ppc64le linux/s390x windows/amd64
TARGET_OBJS ?= darwin-amd64.tar.gz darwin-amd64.tar.gz.sha256 darwin-amd64.tar.gz.sha256sum linux-amd64.tar.gz linux-amd64.tar.gz.sha256 linux-amd64.tar.gz.sha256sum linux-386.tar.gz linux-386.tar.gz.sha256 linux-386.tar.gz.sha256sum linux-arm.tar.gz linux-arm.tar.gz.sha256 linux-arm.tar.gz.sha256sum linux-arm64.tar.gz linux-arm64.tar.gz.sha256 linux-arm64.tar.gz.sha256sum linux-ppc64le.tar.gz linux-ppc64le.tar.gz.sha256 linux-ppc64le.tar.gz.sha256sum linux-s390x.tar.gz linux-s390x.tar.gz.sha256 linux-s390x.tar.gz.sha256sum windows-amd64.zip windows-amd64.zip.sha256 windows-amd64.zip.sha256sum
BINNAME     ?= cabalrpc

GOPATH        = $(shell go env GOPATH)

# go option
PKG        := ./...
TAGS       :=
TESTS      := .
TESTFLAGS  :=
LDFLAGS    := -w -s
GOFLAGS    :=
SRC        := $(shell find . -type f -name '*.go' -print)

# Required for globs to work correctly
SHELL      = /usr/bin/env bash


.PHONY: all
all: build

# ------------------------------------------------------------------------------
#  build

.PHONY: build
build: $(BINDIR)/$(BINNAME)

$(BINDIR)/$(BINNAME): $(SRC)
	GO111MODULE=on go build $(GOFLAGS) -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o $(BINDIR)/$(BINNAME) ./cmd/cabalrpc

# ------------------------------------------------------------------------------
#  test

.PHONY: test
test: build
ifeq ($(ARCH),s390x)
test: TESTFLAGS += -v
else
test: TESTFLAGS += -race -v
endif
test: test-unit

.PHONY: test-unit
test-unit:
	@echo
	@echo "==> Running unit tests <=="
	GO111MODULE=on go test $(GOFLAGS) -run $(TESTS) $(PKG) $(TESTFLAGS)

.PHONY: test-coverage
test-coverage:
	@echo
	@echo "==> Running unit tests with coverage <=="
	@ ./scripts/coverage.sh

.PHONY: coverage
coverage:
	@scripts/coverage.sh

.PHONY: clean
clean:
	@rm -rf $(BINDIR) ./_dist
