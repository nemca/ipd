PROJECTNAME := ipd
VERSION := $(shell cat VERSION)
BUILD := $(shell git rev-parse --short HEAD)
SHELL := /bin/bash

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

# Use linker flags to provide version/build settings
LDFLAGS=-ldflags "-s -w -X=main.version=$(VERSION) -X=main.build=$(BUILD)"

.PHONY: build run test tidy clean help

all: build

## build: Compile the binary.
build: clean tidy test
	@mkdir -p bin
	@go build $(LDFLAGS) -o $(PROJECTNAME) cmd/main.go

## run: Run the go run command.
run: test
	@go run cmd/main.go

## test: Run the go test command.
test:
	@go test -v ./...

## tidy: Run the go mod tidy command.
tidy:
	@go mod tidy

## clean: Cleanup binary.
clean:
	-@rm -f $(PROJECTNAME)

## help: Show this message.
help: Makefile
	@echo "Available targets:"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
