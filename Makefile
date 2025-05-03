APP_NAME := oidc
BIN_DIR := bin
BUILD_DIR := $(BIN_DIR)/$(APP_NAME)
GO_FILES := $(shell find . -name '*.go' -type f)

VERSION := v0.1.0
LDFLAGS := -X main.version=$(VERSION)

.PHONY: all build clean run test

build: $(BUILD_DIR)

$(BUILD_DIR): $(GO_FILES)
	@echo ">> Building $(APP_NAME)..."
	@mkdir -p $(BIN_DIR)
	go build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR) ./cmd/oidc

run: build
	@$(BUILD_DIR)

test:
	go test ./...

clean:
	@echo ">> Cleaning up..."
	@rm -rf $(BIN_DIR)
