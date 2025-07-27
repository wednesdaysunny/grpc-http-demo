.PHONY: build run clean proto dev test

# Build the application
build:
	go build -o bin/server cmd/main.go

# Run the application
run: build
	./bin/server

# Clean build artifacts
clean:
	rm -rf bin/

# Generate protobuf files
proto:
	buf generate

# Run in development mode with auto-reload
dev:
	go run cmd/main.go

# Test the application
test:
	go test ./...

# Install dependencies
deps:
	go mod tidy
	go mod download

# Format code
fmt:
	go fmt ./...

# Lint code
lint:
	golangci-lint run

# Install required tools
install-tools:
	go install github.com/bufbuild/buf/cmd/buf@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest