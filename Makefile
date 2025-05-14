APP_NAME    := oidc
BIN_DIR     := bin
VERSION     := v0.1.0
MODULE_NAME := github.com/wallanaq/oidc-cli
BUILD_DIR   := $(BIN_DIR)/$(APP_NAME)
GO_FILES    := $(shell find . -name '*.go' -type f)

LDFLAGS :=
LDFLAGS += -X $(MODULE_NAME)/pkg/cmd/root.version=$(VERSION)

.PHONY: all build clean run test

build: $(BUILD_DIR)

$(BUILD_DIR): $(GO_FILES)
	@echo ">> Building $(APP_NAME)..."
	@mkdir -p $(BIN_DIR)
	go build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR) ./cmd/oidccli

all: build

run: build
	@$(BUILD_DIR)

test:
	go test ./...

clean:
	@echo ">> Cleaning up..."
	@rm -rf $(BIN_DIR)
