install: gen-grpc
	@mkdir -p bin
	go build -o bin/server main.go
	go build -o bin/client cmd/client/main.go

build-docker:
	docker build -t gh-search-app .

clean:
	@rm -rf bin/
	@go clean -testcache

gen-grpc: setup-tools
	protoc --proto_path=api --go_out=api --go_opt=paths=source_relative --go-grpc_out=api --go-grpc_opt=paths=source_relative api/gh-search/v1/*.proto

setup-tools:
	@if ! command -v protoc &> /dev/null; then \
		echo "protoc could not be found, installing..."; \
		brew install protobuf; \
	fi
	@if [ ! -f "$$(go env GOPATH)/bin/protoc-gen-go" ]; then \
		echo "protoc-gen-go not found, installing..."; \
		go install google.golang.org/protobuf/cmd/protoc-gen-go@latest; \
	fi
	@if [ ! -f "$$(go env GOPATH)/bin/protoc-gen-go-grpc" ]; then \
		echo "protoc-gen-go-grpc not found, installing..."; \
		go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest; \
	fi

test: gen-grpc
	go test ./...

run-server: install
	@if [ -z "$GITHUB_TOKEN" ]; then \
		echo "GITHUB_TOKEN is not set"; \
		exit 1; \
	fi
	./bin/server

run-client: install
	./bin/client

run-docker:
	@if [ -z "$(GITHUB_TOKEN)" ]; then \
		echo "GITHUB_TOKEN is not set"; \
		exit 1; \
	fi
	@echo "Running Docker container..."
	@docker run -e GITHUB_TOKEN=$(GITHUB_TOKEN) -p 50051:50051 --name gh-search-container --rm gh-search-app
