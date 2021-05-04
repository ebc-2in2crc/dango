.DEFAULT_GOAL := help

GOCMD := env GO111MODULE=on go
GOMOD := $(GOCMD) mod
GOBUILD := $(GOCMD) build
GOINSTALL := $(GOCMD) install
GOCLEAN := $(GOCMD) clean
GOTEST := $(GOCMD) test
GOGET := $(GOCMD) get
GOFMT := $(GOCMD) fmt
NAME := dango
CURRENT := $(shell pwd)
BUILDDIR := ./build
BINDIR := $(BUILDDIR)/bin

VERSION := $(shell git describe --tags --abbrev=0)
LDFLAGS := -X 'main.version=$(VERSION)'

export GO111MODULE=on

.PHONY: deps
## Install dependencies
deps:
	$(GOMOD) download

.PHONY: devel-deps
## Install dependencies for develop
devel-deps: deps
	sh -c '\
	tmpdir=$$(mktemp -d); \
	cd $$tmpdir; \
	$(GOGET) \
		github.com/golangci/golangci-lint/cmd/golangci-lint \
		github.com/goreleaser/goreleaser \
		github.com/Songmu/make2help/cmd/make2help \
		github.com/tcnksm/ghr; \
	rm -rf $$tmpdir'

.PHONY: build
## Build binaries
build: deps
	$(GOBUILD) -ldflags "$(LDFLAGS)" -o $(BINDIR)/$(NAME) .

.PHONY: cross-build
## Cross build binaries
cross-build:
	goreleaser build --rm-dist

.PHONY: package
## Make package
package:
	goreleaser release --snapshot --skip-publish --rm-dist

.PHONY: release
## Release package to Github
release:
	goreleaser release --rm-dist

.PHONY: test
## Run tests
test: deps
	$(GOTEST) -v ./...

.PHONY: lint
## Lint
lint: devel-deps
	golangci-lint run ./...

.PHONY: fmt
## Format source codes
fmt: deps
	$(GOFMT) ./...

.PHONY: clean
clean:
	$(GOCLEAN)
	rm -rf $(BUILDDIR)

.PHONY: dockerfile
## Update Dockerfile
dockerfile:
	sed -e "s/<VERSION>/$(VERSION)/g" Dockerfile.base > Dockerfile

.PHONY: help
## Show help
help:
	@make2help $(MAKEFILE_LIST)
