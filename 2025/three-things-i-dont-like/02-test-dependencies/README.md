# Test Dependencies

[proposal: cmd/go: mark test dependencies in go.mod](https://github.com/golang/go/issues/44743)

Test and tool dependencies are tracked in the same `go.mod`; thus, downloading them is unnecessary when building a production artifact.

## Example

> Using old `tools.go` paradigm because `go get -tool` does not support build tags yet (2025-04).

If you already have [`golang-migrate/migrate`](github.com/golang-migrate/migrate) installed in your `PATH`, make sure to use something like [`direnv`](https://mariocarrion.com/2021/10/15/learning-golang-versioning-tools-as-dependencies.html) to _"sandbox"_ this example.

Make sure you install the `migrate` tool using:

```
go install -tags 'pgx5' github.com/golang-migrate/migrate/v4/cmd/migrate
```

Another tool used in this example is `sqlc`, you don't need to install it but if you want you can run:

```
go install github.com/sqlc-dev/sqlc/cmd/sqlc
```

## Docker errors?

To run the tests you need **Docker**. Current code worked using `Docker Desktopr` on MacOS with the following versions:
* Version: `4.38.0 (181591)`, and
* Engine: `27.5.1`.

If you're getting errors when running `go test ./...` such as:

> Couldn't start resource: dial unix /var/run/docker.sock: connect: no such file or directory

Go to `"Settings"` -> `"Advanced"` -> `"Allow the default Docker socked to be used (requires password)"` and enabled it; if it was enabled already, disable it, `"Apply & Restart"`, enable it again and `"Apply & Restart"`.
