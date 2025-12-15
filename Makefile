# Shadis Makefile

# Variables
BINARY_NAME=shadis
BUILD_DIR=bin
CMD_DIR=cmd/shadis

# Go commands
GO=go
GOBUILD=$(GO) build
GORUN=$(GO) run
GOTEST=$(GO) test
GOMOD=$(GO) mod

# Build flags
LDFLAGS=-ldflags "-s -w"

.PHONY: all build run dev clean test tidy help

## help: Show this help message
help:
	@echo "Shadis - A Redis clone in Go"
	@echo ""
	@echo "Usage:"
	@echo "  make <target>"
	@echo ""
	@echo "Targets:"
	@grep -E '^## ' Makefile | sed 's/## /  /'

## build: Build the binary
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) ./$(CMD_DIR)
	@echo "Binary built: $(BUILD_DIR)/$(BINARY_NAME)"

## run: Build and run the server
run: build
	@echo "Starting $(BINARY_NAME)..."
	./$(BUILD_DIR)/$(BINARY_NAME)

## dev: Run in development mode with hot reload (go run)
dev:
	@echo "Starting $(BINARY_NAME) in dev mode..."
	SHADIS_LOG_LEVEL=debug $(GORUN) ./$(CMD_DIR)

## clean: Remove build artifacts
clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	@echo "Done."

## test: Run tests
test:
	@echo "Running tests..."
	$(GOTEST) -v ./...

## tidy: Tidy go modules
tidy:
	@echo "Tidying modules..."
	$(GOMOD) tidy

# Default target
all: build

