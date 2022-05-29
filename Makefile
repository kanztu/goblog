GOCMD=go
BINARY_NAME=goblog
BINARY_NAME_WASM=blog_common.wasm
BUILD_FOLDER=build
BUILD_WASM_FOLDER=src/wasm

SWAG_PATH=./docs/
SWAG_BIN=swag

all: build gen_docs

build:
	mkdir -p $(BUILD_FOLDER) $(BUILD_WASM_FOLDER)
	$(GOCMD) build -o $(BUILD_FOLDER)/$(BINARY_NAME) cmd/webserver/webserver.go
	# GOOS=js GOARCH=wasm $(GOCMD) build -o $(BUILD_WASM_FOLDER)/$(BINARY_NAME_WASM) cmd/wasm/blog_common/*

gen_docs:
	$(SWAG_BIN) init -g cmd/webserver/webserver.go

clean:
	rm -r $(BUILD_FOLDER) $(BUILD_WASM_FOLDER) $(SWAG_PATH)
