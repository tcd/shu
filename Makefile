GO ?= go
SHELL := /bin/sh
GOBIN_DIR=${GOBIN}
PROJECT_DIR=$(shell pwd)
PROJECT_NAME=$(shell basename $(PROJECT_DIR))

# Install shu to $GOBIN.
install:
	GO111MODULE=on $(GO) install

# Remove shu from $GOBIN.
uninstall:
	@rm -f $(GOBIN_DIR)/$(PROJECT_NAME)

# Run all tests for the project.
test:
	go test -v ./...

build:
	GO111MODULE=on $(GO) build -o build/$(PROJECT_NAME)

clean:
	$(GO) clean ./...
	rm -rf build

mod:
	GO111MODULE=on go mod tidy

.PHONY: build clean test install uninstall
