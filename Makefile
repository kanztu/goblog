GOCMD=go
BINARY_NAME=goblog
BINARY_NAME_WASM=main.wasm
BUILD_FOLDER=build
BUILD_WASM_FOLDER=wasm
GOROOT := $(shell go env GOROOT)


SWAG_PATH=./docs/
SWAG_BIN=swag

all: build

build: gen_docs
	mkdir -p $(BUILD_FOLDER) $(BUILD_WASM_FOLDER)
	$(GOCMD) build -o $(BUILD_FOLDER)/$(BINARY_NAME) cmd/webserver/webserver.go
	GOOS=js GOARCH=wasm $(GOCMD) build -o $(BUILD_WASM_FOLDER)/$(BINARY_NAME_WASM) cmd/wasm/blog_common/*

gen_docs:
	$(SWAG_BIN) init -g cmd/webserver/webserver.go

clean:
	rm -rf $(BUILD_FOLDER) $(BUILD_WASM_FOLDER)/$(BINARY_NAME_WASM) $(SWAG_PATH) ./static/wasm_exec.js
