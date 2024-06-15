# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get
GORUN = $(GOCMD) run

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

db-down:
	sudo docker compose -f ./docker-compose.dev.yml down

db-up:
	sudo docker compose -f ./docker-compose.dev.yml up -d

http:
	$(GORUN) cmd/main.go

reset-db: db-down db-up

.PHONY: all build test clean run