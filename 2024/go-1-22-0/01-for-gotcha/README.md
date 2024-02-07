# Fix `for` gotcha

Excerpt from official release notes:

> The long-standing "for" loop gotcha with accidental sharing of loop variables between iterations is now resolved.

This means there's no need to create a new variable inside the `for`.

## Running example

## With 1.22 (new behavior)

* Run `go build main.go`, it will create a binary `main`
* Verify it was built using Go 1.22, `go version -m main`
* Run it, you should see the new behavior

## With 1.21 (old behavior)

Use a full prefix of your Go compiler, for example `/usr/local/bin/go`:

* Run `/usr/local/bin/go build main.go`, it will replace/create a binary `main`
* Verify it was built using Go 1.21, `go version -m main`
* Run it, you should see the new behavior
