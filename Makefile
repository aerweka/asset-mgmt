# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get
GORUN = 

# Build target
BINARY_NAME = ./cmd

all: test build

build:
	$(GOBUILD) -o $(BINARY_NAME) main.go -v

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run:
	# build
	./$(BINARY_NAME)

run-dev:
	reflex -s -r '\.go$$' go run cmd/main.go

.PHONY: all build test clean run