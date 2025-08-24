# `go doc`

Excerpt from [official release notes](https://go.dev/doc/go1.25#go-command):

> The new `go doc -http` option will start a documentation server showing
> documentation for the requested object, and open the documentation in a
> browser window.

For example if you do: `go doc -http` in your local Go module, it will default
to it, however you can also use a package name for example:

```
go doc -http errors
```

That will automatically open your default browser pointing to a local
pkg.go.dev-like page displaying the `errors` documentation.
