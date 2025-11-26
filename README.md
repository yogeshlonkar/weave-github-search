# weave-github-search

The service is built in golang using provided gRPC spec and GitHub search API.

It uses [google/go-github](https://github.com/google/go-github) to interact with GitHub API.

It has a client to interact with the gRPC server.

## Development

```shell
make setup-tools
```

### Generating gRPC code

```shell
make gen-grpc
```

## Running tests

```shell
make test
```

## Running the server

### On localhost

Run the following command to start the gRPC server:

```shell
export GITHUB_TOKEN=#YOUR_GITHUB_TOKEN_HERE
make run-server
```

### Using Docker

Build the Docker image:

```shell
make build-docker
export GITHUB_TOKEN=#YOUR_GITHUB_TOKEN_HERE
make run-docker
```

Both of above options will start the gRPC server on port `50051`.
The port can be changed by exporting `GRPC_PORT` environment variable before running the server.

You can then connect to the server using the interactive client mentioned below or any other gRPC client.

## Running the client

Run the following command to start the interactive gRPC client:

```shell
make run-client
```

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
