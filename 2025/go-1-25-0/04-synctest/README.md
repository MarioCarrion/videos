# New `testing/synctest` package

Excerpt from [official release notes](https://go.dev/doc/go1.25#new-testingsynctest-package):

> The new `testing/synctest` package provides support for testing concurrent code.
> 
> The `Test` function runs a test function in an isolated "bubble". Within the
> bubble, time is virtualized: time package functions operate on a fake clock
> and the clock moves forward instantaneously if all goroutines in the bubble
> are blocked.
> 
> The `Wait` function waits for all goroutines in the current bubble to block.
> 
> This package was first available in Go 1.24 under `GOEXPERIMENT=synctest`,
> with a slightly different API. The experiment has now graduated to general
> availability. The old API is still present if `GOEXPERIMENT=synctest` is set,
> but will be removed in Go 1.26.

Blog post by the Go team: [Testing Time (and other asynchronicities)](https://go.dev/blog/testing-time)

## Running

Use `go test -race -v ./...`
