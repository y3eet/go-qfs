# Simple Makefile for a Go project

# Build the application
all: build test

build:
	@echo "Building..."
	@go build -o main cmd/api/main.go

build-all:
	@echo "Building all..."
	@cd frontend && bun install && bun run build
	@go build -o main cmd/api/main.go

# Run the application
run:
	@go run main.go

# Test the application
test:
	@echo "Testing..."
	@go test ./... -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

prepare:
	@go install github.com/wailsapp/wails/v2/cmd/wails@latest
	@cd frontend && bun install && bun run build
	@go mod tidy && cp .env.example .env

# Live Reload
dev:
	@cd frontend && bun dev & \
	if command -v air > /dev/null; then \
		air; \
		echo "Watching..."; \
	else \
		read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
		if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
			go install github.com/air-verse/air@latest; \
			air; \
			echo "Watching..."; \
		else \
			echo "You chose not to install air. Exiting..."; \
			exit 1; \
		fi; \
	fi

.PHONY: all build run test clean watch
