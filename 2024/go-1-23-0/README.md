# Go 1.23: What is new?

* [Video](https://youtu.be/EL4hg73mT2A)
* [Blog Post](https://mariocarrion.com/2024/09/02/what-is-new-in-go-1-23.html)

Go 1.23.0 was released on August 13th, 2024. This directory includes examples of some new features added in this release.

* [@Golang X announcement](https://x.com/golang/status/1823404920583663807)
* [Download](https://go.dev/dl/#go1.23.0)
* [Release notes](https://go.dev/doc/go1.23)
* [Official Blog Post](https://go.dev/blog/go1.23)
* [New API Changes for 1.23](https://github.com/golang/go/issues/67843)

* Other ways to get Go 1.23:
  * Docker Images:
    * [All](https://hub.docker.com/_/golang/tags?page=&page_size=&ordering=&name=1.23.0)
    * [Debian 12 (Bookworm)](https://hub.docker.com/layers/library/golang/1.23.0-bookworm/images/sha256-e3e4449f4e4e7f740122cd1d6c63fe1fb8cd16d526f6c68fc90773111f5b89b1?context=explore) `docker pull golang:1.23.0-bookworm`
    * [Alpine 3.20](https://hub.docker.com/layers/library/golang/1.23.0-alpine3.20/images/sha256-ce1e987ea7759217351b74977a384cea8f44631f1c1add04d1703f13dd3ee850?context=explore): `docker pull golang:1.23.0-alpine3.20`
  * [Homebrew Formula PR](https://github.com/Homebrew/homebrew-core/pull/175310)

## New features

### Changes to the language

* [Range over func](01-range-over-func/)

### Standard Library

* [`unique` package](02-unique/)

### Iterators

* [`iter` package changes](03-iter/)
* [`slices` package changes](04-slices/)
* [`maps` package changes](05-maps/)

### Minor changes to the library

* [`net/http.Request.Pattern` and `slices.Repeat`](06-minor/)
