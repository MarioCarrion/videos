# `go get -tool`

* [Official release notes](https://go.dev/doc/go1.24#go-command)

The new `tool` command supports versioning multiple modules using the same tool. However, different versions can coexist, negating the need to use tools such as `direnv,` that sandbox your environment variables, such as `PATH` or `GOBIN.`

## Instructions

For example, to get the latest version of `golangci-lint` using the new `tool` support, run:

```
go get -tool github.com/golangci/golangci-lint/cmd/golangci-lint
```

Then you can run that binary using `go tool`:

```
go tool golangci-lint run ./...
```
