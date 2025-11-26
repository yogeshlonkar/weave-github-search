# weave-github-search

The service is built in golang using provided gRPC spec and GitHub search API.

It uses [google/go-github](https://github.com/google/go-github) to interact with GitHub API.

It has a client to interact with the gRPC server.

## Development

Required tools:

- [protoc](https://grpc.io/docs/protoc-installation/)
- protoc-gen-go - `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
- protoc-gen-go-grpc - `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`

### Generating gRPC code

```shell
protoc --proto_path=api --go_out=api --go_opt=paths=source_relative --go-grpc_out=api --go-grpc_opt=paths=source_relative api/gh-search/v1/*.proto
```

## Running tests

```shell
go test ./...
```

## Running the server

Run the following command to start the gRPC server:

```shell
export GITHUB_TOKEN=#YOUR_GITHUB_TOKEN_HERE
go run cmd/server/main.go
```

## Running the client

Run the following command to start the interactive gRPC client:

```shell
go run cmd/client/main.go
```

## File structure

```
.
├── api
│   └── gh-search
│       └── v1      # gRPC service definitions and proto files
├── cmd
|   └── client      # interactive client for testing
├── internal
│   ├── github      # GitHub API client implementation
│   └── grpc        # gRPC service implementation
│── main.go         # main application entrypoint
```
