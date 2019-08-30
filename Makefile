GO ?= go
SHELL := /bin/sh
GOBIN_DIR=${GOBIN}
PROJECT_DIR=$(shell pwd)
PROJECT_NAME=$(shell basename $(PROJECT_DIR))

build:
	GO111MODULE=on $(GO) build -o build/$(PROJECT_NAME)

clean:
	$(GO) clean ./...
	rm -rf build

# Run all tests for the project.
test:
	go test -v ./...

# Install shu to $GOBIN.
install:
	GO111MODULE=on $(GO) install

# Remove shu from $GOBIN.
uninstall:
	@rm -f $(GOBIN_DIR)/$(PROJECT_NAME)

.PHONY: build clean test install uninstall
