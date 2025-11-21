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
