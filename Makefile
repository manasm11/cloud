# Makefile for cloud CLI
# Cross-platform compatible (Windows, macOS, Linux)

.PHONY: help build run test test-coverage lint fmt vet clean deps install-tools ci pre-commit

# Default target
.DEFAULT_GOAL := help

# Detect OS
ifeq ($(OS),Windows_NT)
    DETECTED_OS := Windows
    EXE_EXT := .exe
    RM_CMD := powershell -Command "Remove-Item -Recurse -Force -ErrorAction SilentlyContinue"
    MKDIR_CMD := powershell -Command "New-Item -ItemType Directory -Force -Path"
    NULL_DEVICE := NUL
    # Race detector requires CGO, which may not be available on Windows
    RACE_FLAG :=
else
    DETECTED_OS := $(shell uname -s)
    EXE_EXT :=
    RM_CMD := rm -rf
    MKDIR_CMD := mkdir -p
    NULL_DEVICE := /dev/null
    RACE_FLAG := -race
endif

# Variables
APP_NAME := cloud
VERSION := $(shell git describe --tags --always --dirty 2>$(NULL_DEVICE) || echo dev)
LDFLAGS := -ldflags "-X main.version=$(VERSION) -X main.buildTime=unknown -X main.gitCommit=$(VERSION)"

## help: Show this help message
help:
	@echo Usage: make [target]
	@echo ""
	@echo Targets:
	@echo   build          - Build the application
	@echo   run            - Run the application
	@echo   test           - Run all tests
	@echo   test-coverage  - Run tests with coverage report
	@echo   lint           - Run linter
	@echo   fmt            - Format code
	@echo   vet            - Run go vet
	@echo   clean          - Clean build artifacts
	@echo   deps           - Download dependencies
	@echo   install-tools  - Install development tools
	@echo   ci             - Run all CI checks
	@echo   pre-commit     - Run pre-commit checks

## build: Build the application
build:
	go build $(LDFLAGS) -o bin/$(APP_NAME)$(EXE_EXT) ./cmd/$(APP_NAME)

## run: Run the application
run:
	go run ./cmd/$(APP_NAME)

## test: Run all tests
test:
	go test -v $(RACE_FLAG) -cover ./...

## test-short: Run tests without verbose output
test-short:
	go test $(RACE_FLAG) -cover ./...

## test-coverage: Run tests with coverage report
test-coverage:
	go test -v $(RACE_FLAG) -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo Coverage report generated: coverage.html

## lint: Run linter
lint:
	golangci-lint run ./...

## fmt: Format code
fmt:
	go fmt ./...
	-goimports -w .

## vet: Run go vet
vet:
	go vet ./...

## clean: Clean build artifacts
clean:
ifeq ($(OS),Windows_NT)
	@powershell -Command "if (Test-Path bin) { Remove-Item -Recurse -Force bin }"
	@powershell -Command "if (Test-Path coverage.out) { Remove-Item -Force coverage.out }"
	@powershell -Command "if (Test-Path coverage.html) { Remove-Item -Force coverage.html }"
else
	rm -rf bin/
	rm -f coverage.out coverage.html
endif

## deps: Download dependencies
deps:
	go mod download
	go mod tidy

## install-tools: Install development tools
install-tools:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install golang.org/x/tools/cmd/goimports@latest

## ci: Run all CI checks
ci: lint vet test-short

## pre-commit: Run pre-commit checks
pre-commit: fmt lint vet test-short
