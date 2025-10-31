.PHONY: build clean install help test deps run

# Document targets using a comment starting with "##" to display them in the help output.

# Binary name
BINARY_NAME=pkm

# Build directory
BIN_DIR=bin

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod

build: ## Build the binary to bin/pkm
	@mkdir -p $(BIN_DIR)
	$(GOBUILD) -o $(BIN_DIR)/$(BINARY_NAME) -v

clean: ## Remove build artifacts
	$(GOCLEAN)
	rm -rf $(BIN_DIR)

install: build ## Install the binary to GOPATH/bin
	$(GOCMD) install

test: ## Run tests
	$(GOTEST) -v ./...

deps: ## Download and tidy dependencies
	$(GOMOD) download
	$(GOMOD) tidy

run: build ## Build and run the application
	./$(BIN_DIR)/$(BINARY_NAME)

help: ## Display this help message
	@echo "Available targets:"
	@awk 'BEGIN {FS = ":.*##"; printf "\n"} /^[a-zA-Z_-]+:.*##/ { printf "  %-10s %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

# Default target
.DEFAULT_GOAL := build
