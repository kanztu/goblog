GOCMD=go
BINARY_NAME=goblog
BUILD_FOLDER=build

all: build

build:
	mkdir -p $(BUILD_FOLDER)
	$(GOCMD) build -o $(BUILD_FOLDER)/$(BINARY_NAME) cmd/webserver/webserver.go

clean:
	rm -r $(BUILD_FOLDER)
