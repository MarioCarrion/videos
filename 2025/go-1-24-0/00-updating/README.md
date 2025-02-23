# Updating

Thanks to Go 1.21 (or greater). There's no need to install a version explicitly; instead, editing the go.mod to indicate the version to use should suffice.

For example assuming Go **1.23.6** is installed, running `go version` prints out the following:

```
go version go1.23.6 darwin/amd64
```

Next, editing `go.mod` to update the Go version:

```
go mod edit -go=1.24.0
```

Finally, running `go version` will print out:

```
go version go1.24.0 darwin/amd64
```
