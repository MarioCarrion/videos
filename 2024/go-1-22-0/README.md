# Go 1.22: What is new?

**NOTE** Install [`direnv`](https://mariocarrion.com/2020/11/20/golang-go-tool-direnv.html) and then run `make install` to install the version you get by using `go install`.

---

Go 1.22.0 was released on February 6th, 2024. This directory includes examples of some new features worth mentioning that were added in this release.

* [@Golang Twitter announcement](https://twitter.com/golang/status/1755005673975218662)
* [Download](https://go.dev/dl/#go1.22.0)
* [Release notes](https://go.dev/doc/go1.22)
* [Official Blog Post](https://go.dev/blog/go1.22)
* [New API Changes for 1.22](https://github.com/golang/go/issues/64343)

* Other ways to use Go:
  * Docker Images
    * [All](https://hub.docker.com/_/golang/tags?page=1&name=1.22)
    * [Debian 12 (Bookworm)](https://hub.docker.com/layers/library/golang/1.22.0-bookworm/images/sha256-e1ed73ccd2f67731fa03ec551d11df697838e9e0a5054163e7753d34add15e99?context=explore) `docker pull golang:1.22.0-bookworm`
    * [Alpine 3.19](https://hub.docker.com/layers/library/golang/1.22.0-alpine3.19/images/sha256-3325c5593767d8f1fd580e224707ca5e847a1679470a027aaa3c71711ce16613?context=explore): `docker pull golang:1.22.0-alpine3.19`
  * [Homebrew Formula PR](https://github.com/Homebrew/homebrew-core/pull/157782)

## New features

### Changes to the language

* [`for` gotcha](01-for-gotcha/)
* [`for` range over integers](02-for-over-integers/)

### Core library

* [New math/rand/v2 package](03-rand-v2/)
* [Enhanced routing patterns](04-routing-patterns/)

#### Minor changes to the library

* [net/http: ServeFileFS, FileServerFS, and NewFileTransportFS](05-http-fs/)
* [slices: Concat](06-slices-concat/)
