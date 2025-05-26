APP_NAME    := oidc
BIN_DIR     := bin
VERSION     := v0.2.0
MODULE_NAME := github.com/wallanaq/oidc-cli
INSTALL_DIR := /usr/local/bin/
BUILD_DIR   := $(BIN_DIR)/$(APP_NAME)
GO_FILES    := $(shell find . -name '*.go' -type f)

LDFLAGS :=
LDFLAGS += -X $(MODULE_NAME)/internal/version.Current=$(VERSION)

all: build

build: clean $(BUILD_DIR)

$(BUILD_DIR): $(GO_FILES)
	@echo ">> Building $(APP_NAME)..."
	@mkdir -p $(BIN_DIR)
	go build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR) ./cmd/oidccli

install: build
	@echo ">> Installing $(APP_NAME)..."
	@sudo mv $(BUILD_DIR) $(INSTALL_DIR)

run: build
	@$(BUILD_DIR)

test:
	go test ./...

clean:
	@echo ">> Cleaning up..."
	@rm -rf $(BIN_DIR)

.PHONY: all build install clean run test
