# New experimental `testing/synctest` package

Excerpt from [official release notes](https://go.dev/doc/go1.24#testing-synctest):

> The new experimental `testing/synctest` package provides support for testing concurrent code.
>
> The `synctest.Run` function starts a group of goroutines in an isolated “bubble”. Within the bubble, time package functions operate on a fake clock.
> The `synctest.Wait` function waits for all goroutines in the current bubble to block.

## Running

Use `GOEXPERIMENT=synctest go test -v ./...`
