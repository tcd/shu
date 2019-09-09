GO ?= go
PACKR ?= packr2
SHELL := /bin/sh
GOBIN_DIR=${GOBIN}
PROJECT_DIR=$(shell pwd)
PROJECT_NAME=$(shell basename $(PROJECT_DIR))

# Install shu to $GOBIN.
install:
	GO111MODULE=on $(PACKR) install

# Remove shu from $GOBIN.
uninstall:
	@rm -f $(GOBIN_DIR)/$(PROJECT_NAME)

# Run all tests for the project.
test:
	go test -v ./...

build:
	GO111MODULE=on $(PACKR) build -o build/$(PROJECT_NAME)

clean:
	$(GO) clean ./...
	$(PACKR) clean
	rm -rf build

mod:
	GO111MODULE=on go mod tidy

.PHONY: build clean test install uninstall
