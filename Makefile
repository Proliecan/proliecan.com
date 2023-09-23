# go compiler
GO=go

# src dir
SRC_DIR=src

# build dir
BUILD_DIR=out
# binary name
BINARY_NAME=server

# gofiles
GOFILES=$(shell find $(SRC_DIR) -name "*.go")
MAIN_FILE=$(SRC_DIR)/server.go

# run
run: build
	$(BUILD_DIR)/$(BINARY_NAME)

# build
build: $(GOFILES)
	$(GO) build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_FILE)

# clean
clean:
	rm -rf $(BUILD_DIR)/*