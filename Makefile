GO = go
BIN_NAME = wordle
BIN_DIR = bin
TARGET = $(BIN_DIR)/$(BIN_NAME)

.PHONY: run

all: test build

build:
	$(GO) build -o $(TARGET) -v

test:
	$(GO) test -v ./...

run: test build
	./$(TARGET)

clean:
	$(GO) clean
	rm -fr $(BIN_DIR)
