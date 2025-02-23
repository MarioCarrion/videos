# Go 1.24.0: What's new?

* [Video](https://youtu.be/qjiIA0220ms)
* [Blog Post](https://mariocarrion.com/2025/02/23/what-is-new-in-go-1-24.html)

Go **1.24.0** was released on _February 11th, 2025_. This directory includes examples of some new features added in this release.

* [@Golang X announcement](https://x.com/golang/status/1889383557182689695)
* [Download](https://go.dev/dl/#go1.24.0)
* [Release notes](https://go.dev/doc/go1.24)
* [Official Blog Post](https://go.dev/blog/go1.24)
* [1.24 Milestone on Github](https://github.com/golang/go/milestone/322?closed=1)

* Other ways to get Go 1.24.0:
  * Docker Images:
    * [All](https://hub.docker.com/_/golang/tags?name=1.24)
    * [Debian Bullseye](https://hub.docker.com/layers/library/golang/1.24-bullseye/images/sha256-641d2e1480c57b97df7c8fc8585822264aa37c8d0b25f2905046624a72d77807) `docker pull golang:1.24-bullseye`
    * [Alpine 3.21](https://hub.docker.com/layers/library/golang/1.24.0-alpine3.21/images/sha256-a88aef3ac8a2c2fd547f65eba6e8220aace8eabba5d93c00c9f76bdfba578c68): `docker pull golang:1.24.0-alpine3.21`
  * [Homebrew Formula Pull Request](https://github.com/Homebrew/homebrew-core/pull/201070)

# New features in Go 1.24.0

* [Updating](00-updating/)

## Tools

* [Go command](01-go-tool/)

## Standard library

* [Directory-limited filesystem access](02-os-root/)
* [`encoding/json` omitzero](03-json-omitzero/)
* [`strings` package updates](04-strings/)
* [`testing` package updates](05-testing/)
* [New experimental `testing/synctest` package](06-synctest/)
