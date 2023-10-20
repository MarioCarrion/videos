# Go 1.21: What is new?

[<img src="https://github.com/MarioCarrion/MarioCarrion/blob/main/youtube.svg" width="20" height="20" alt="YouTube video"> Watch the video!](https://youtu.be/fO2a-5Wyh0I)

**NOTE** Install [`direnv`](https://mariocarrion.com/2020/11/20/golang-go-tool-direnv.html) and then run `make install` to install the version you get by using `go install`.

---

Go 1.21.0 was released on August 8th, 2023; this directory includes examples of some new features I consider worth mentioning added in this release.

* [@Golang Twitter announcement](https://twitter.com/golang/status/1688944844490539008)
* [Download](https://go.dev/dl/#go1.21)
* [Release notes](https://go.dev/doc/go1.21)
* [Official Blog Post](https://go.dev/blog/go1.21)
* [New API Changes for 1.21](https://github.com/golang/go/issues/60560)

* Other ways to use Go:
  * Docker Images
    * [All](https://hub.docker.com/_/golang/tags?page=1&name=1.21)
    * [Debian Bookworm](https://hub.docker.com/layers/library/golang/1.21.0-bookworm/images/sha256-291849f9186a989b8f5bce22fd10c361d9e446d56ecf335c5f7e8572e98e45a9?context=explore) `docker pull golang:1.21.0-bookworm`
    * [Alpine 3.18](https://hub.docker.com/layers/library/golang/1.21.0-alpine3.18/images/sha256-4c16f271b38a57f79514b432c329decc71af3868d914191ba767256e9c504e65?context=explore): `docker pull golang:1.21.0-alpine3.18`
  * [Homebrew Formula PR](https://github.com/Homebrew/homebrew-core/pull/134468)

## New features

### Tool improvements

* [Backward behavior using GODEBUG](01-godebug/)

### Language changes

* [New built-in functions: min, max and clear](02-new-built-in-funcs/)

### Core library

* New `log/slog` package, this is _basically_ the same package as the one I covered [in a previous episode](../04/07/). You can watch that [video here](https://youtu.be/htbGdhW3JdQ).
* [New `slices` package](03-slices/)
* [New `maps` package](04-maps/)

### Minor changes to the library

* [New `context` functions](05-context)
* [New `sync` functions](06-sync-once)
