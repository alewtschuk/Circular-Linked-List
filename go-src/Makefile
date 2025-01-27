# Makefile for Go Testing Project

# Variables
GO := go

## Default target: Run tests
all: check

## Run tests
check:
	@echo "Running tests..."
	$(GO) test ./... -v

## Clean the project
clean:
	@echo "Cleaning up..."
	$(GO) clean -testcache

## Help command to show available targets
help:
	@echo "Makefile for Go testing project"
	@echo
	@echo "Available targets:"
	@echo "  make           Run tests (default)"
	@echo "  make check     Run tests"
	@echo "  make clean     Clean test cache"
	@echo "  make help      Show this help message"