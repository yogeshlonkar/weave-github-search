SHELL := /bin/bash
export GRPC_PORT ?= 50051

.PHONY: compile-server compile-client

compile-server: gen-grpc
	@mkdir -p bin
	@echo "Building server..."
	@go build -ldflags="-s -w" -o bin/server main.go

compile-client: gen-grpc
	@mkdir -p bin
	@echo "Building client..."
	@go build -ldflags="-s -w" -o bin/client cmd/client/main.go

build-docker:
	docker build -t gh-search-app .

clean:
	@rm -rf bin/
	@go clean -testcache

lint:
	golangci-lint run ./...

gen-grpc: setup-tools
	@echo "Generating gRPC code from .proto files..."
	protoc --proto_path=api --go_out=api --go_opt=paths=source_relative --go-grpc_out=api --go-grpc_opt=paths=source_relative api/gh-search/v1/*.proto

setup-tools:
	@if ! command -v protoc &> /dev/null; then \
		echo "protoc could not be found, installing..."; \
		brew install protobuf; \
	fi
	@if [ ! -f "$$(go env GOPATH)/bin/protoc-gen-go" ]; then \
		echo "protoc-gen-go not found, installing..."; \
		go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.10; \
	fi
	@if [ ! -f "$$(go env GOPATH)/bin/protoc-gen-go-grpc" ]; then \
		echo "protoc-gen-go-grpc not found, installing..."; \
		go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.6.0; \
	fi
	@if ! command -v golangci-lint &> /dev/null; then \
		echo "golangci-lint could not be found, installing..."; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
	fi
	@echo "all tools are set up."

test: gen-grpc
	go test ./...

run-server: compile
	./bin/server

run-client: compile-client
	./bin/client

run-docker: build-docker
	@echo "Running Docker container..."
	@docker run \
		--rm \
		-e GITHUB_TOKEN=$(GITHUB_TOKEN) \
		-e GRPC_PORT=$(GRPC_PORT) \
		--expose $(GRPC_PORT) \
		-p $(GRPC_PORT):$(GRPC_PORT) \
		--name gh-search-container \
		gh-search-app
