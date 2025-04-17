# Makefile for Go project

APP_NAME := go_imagefile_upload
SRC := $(shell find . -type f -name '*.go')
BUILD_DIR := ./bin
BUILD_PATH := $(BUILD_DIR)/$(APP_NAME)

.PHONY: all build run clean test

default: run

all: clean build test

build: $(SRC)
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_PATH) ./cmd

run: build
	@echo "Running $(APP_NAME)..."
	@$(BUILD_PATH)

test:
	@echo "Running tests..."
	@go test ./...

clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)