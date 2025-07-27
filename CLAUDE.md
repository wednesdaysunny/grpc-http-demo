# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Development Commands

### Build and Run
- `make build` - Compile the application to `bin/server`
- `make run` - Build and run the server
- `make dev` - Run in development mode with `go run cmd/main.go`
- `make clean` - Remove build artifacts

### Code Quality
- `make test` - Run all tests with `go test ./...`
- `make fmt` - Format code with `go fmt ./...`
- `make lint` - Run golangci-lint for code analysis

### Dependencies and Tools
- `make deps` - Install/update Go dependencies
- `make install-tools` - Install buf and golangci-lint
- `make proto` - Generate Protocol Buffer files using buf

## Architecture Overview

This is a gRPC-Gateway demo service that provides both gRPC and HTTP APIs for user management. The service demonstrates the dual-protocol pattern where a single service implementation serves both gRPC clients and HTTP/REST clients.

### Key Components

**Service Layer (`server/`)**:
- `user_service.go` - Core gRPC service implementation with in-memory user storage
- `gateway.go` - HTTP gateway configuration that proxies REST calls to gRPC

**Entry Point (`cmd/main.go`)**:
- Starts both gRPC server (port 9091) and HTTP gateway (port 9090) concurrently
- Implements graceful shutdown handling
- Includes health check endpoint at `/health`

**Protocol Definitions (`proto/user/`)**:
- `user.proto` - Defines UserService with CRUD operations and HTTP mappings
- Generated files include Go structs, gRPC server/client code, and gateway handlers

### Service Ports
- gRPC server: 9091
- HTTP gateway: 9090
- Health check: `GET /health`
- API endpoints: `/api/v1/users/*`

### Code Generation Workflow

The project uses buf for Protocol Buffer management:
1. Edit `proto/user/user.proto` for API changes
2. Run `make proto` to regenerate Go code
3. Implement service methods in `server/user_service.go`
4. HTTP endpoints are automatically created via gRPC-Gateway annotations

### Data Storage
Currently uses in-memory storage (`map[string]*pb.User`) in UserServer. Users are identified by auto-generated IDs with timestamp-based uniqueness.