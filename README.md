# weave-github-search

The service is built in golang using provided gRPC spec and GitHub search API.

It uses [google/go-github](https://github.com/google/go-github) to interact with GitHub API.

It has a client to interact with the gRPC server.

## Development

| Command               | Description                                                                  |
|-----------------------|------------------------------------------------------------------------------|
| `make setup-tools`    | Install required development tools                                           |
| `make gen-grpc`       | Generate gRPC code from proto files                                          |
| `make test`           | Run tests                                                                    |
| `make lint`           | Run linter. golangci-lint must be installed (see `make setup-tools`)         |
| `make compile-server` | Compile gRPC server                                                          |
| `make compile-client` | Compile gRPC client                                                          |
| `make run-server`     | Run gRPC server on localhost. Requires `GITHUB_TOKEN` env variable to be set |
| `make build-docker`   | Build Docker image                                                           |
| `make run-docker`     | Run gRPC server using Docker. Requires `GITHUB_TOKEN` env variable to be set |
| `make run-client`     | Run interactive gRPC client                                                  |

Both of `make run-server` and `make run-docker` options will start the gRPC server on port `50051`.
The port can be changed by exporting `GRPC_PORT` environment variable before running the server.

You can then connect to the server using the interactive client mentioned below or any other gRPC client.

## File structure

```
.
├── api
│   └── gh-search
│       └── v1      # gRPC service definitions and proto files
├── bin             # compiled binaries
├── cmd
|   └── client      # interactive client for testing
├── internal
│   ├── github      # GitHub API client implementation
│   └── grpc        # gRPC service implementation
│── Dockerfile      # Dockerfile for containerizing the application
│── go.mod          # Go module file
│── go.sum          # Go module checksum file
│── main.go         # main application entrypoint
│── main_tset.go    # main application tests
│── makefile        # Makefile with build and run commands
└── README.md       # this file
```
