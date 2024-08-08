SRC_DIR := ./cmd
OUTPUT_DIR := ./bin
OUTPUT_NAME := zero-agency

GO_BIN := go

.PHONY: all build clean

all: build

build:
	@echo "Building the Go project..."
	$(GO_BIN) build -o $(OUTPUT_DIR)/$(OUTPUT_NAME) $(SRC_DIR)/*.go

clean:
	@echo "Cleaning up..."
	rm -f $(OUTPUT_DIR)/$(OUTPUT_NAME)
