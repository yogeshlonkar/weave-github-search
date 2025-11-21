# weave-github-search

## Development

Required tools:

- [protoc](https://grpc.io/docs/protoc-installation/)
- protoc-gen-go
- protoc-gen-go-grpc

### Generating gRPC code

```shell
protoc --proto_path=api --go_out=api --go_opt=paths=source_relative --go-grpc_out=api --go-grpc_opt=paths=source_relative api/gh_search/v1/*.proto
```

## Running tests

```shell
go test ./...
```

## Running the server

Run the following command to start the gRPC server:

```shell
go run cmd/gh-search-server/main.go
```

## File structure

```
.
├── api
│   └── gh-search
│       └── v1              # gRPC service definitions and proto files
├── cmd
│   └── gh-search-server    # main application entrypoint
├── internal
│   ├── gh-client           # GitHub API client implementation
│   └── gh-search-service   # gRPC service implementation
```
