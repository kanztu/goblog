GOCMD=go
BINARY_NAME=goblog
BINARY_NAME_WASM=blog_common.wasm
BUILD_FOLDER=build
BUILD_WASM_FOLDER=src/wasm

all: build

build:
	mkdir -p $(BUILD_FOLDER) $(BUILD_WASM_FOLDER)
	$(GOCMD) build -o $(BUILD_FOLDER)/$(BINARY_NAME) cmd/webserver/webserver.go
	GOOS=js GOARCH=wasm $(GOCMD) build -o $(BUILD_WASM_FOLDER)/$(BINARY_NAME_WASM) cmd/wasm/blog_common/*

gen_docs:
	
clean:
	rm -r $(BUILD_FOLDER) $(BUILD_WASM_FOLDER)
