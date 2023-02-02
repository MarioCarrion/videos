# Go 1.20: What is new?

**NOTE** Install [`direnv`](https://mariocarrion.com/2020/11/20/golang-go-tool-direnv.html) and then run `make install` to install the version you get by using `go install`.

---

Go 1.20.0 was released on August 2nd, 2022; in this directory you will find examples of some new features I consider worth mentioning added in this release.

* [@Golang Twitter announcement](https://twitter.com/golang/status/1620875197569187840)
* [Download](https://go.dev/dl/#go1.20)
* [Release notes](https://go.dev/doc/go1.20)
* [New API Changes for 1.20](https://github.com/golang/go/issues/57126)

* Other ways to use Go:
  * Docker Images
    * [All](https://hub.docker.com/_/golang?tab=tags&page=1&name=1.20)
    * [Alpine 3.17](https://hub.docker.com/layers/library/golang/1.20.0-alpine3.17/images/sha256-ebceb16dc094769b6e2a393d51e0417c19084ba20eb8967fb3f7675c32b45774?context=explore): `docker pull golang:1.20.0-alpine3.17`
    * [Debian Bullseye](https://hub.docker.com/layers/library/golang/1.20.0-bullseye/images/sha256-61dafe97cc20b12faef7a744a1a3e43411c45b9908c3ccf9cadd05566e473e1d?context=explore) `docker pull golang:1.20.0-bullseye`
  * [Homebrew Formula PR](https://github.com/Homebrew/homebrew-core/pull/122082)

## New features

### Tools

* [Accept -C](01-accept-c/)
* [Skip for generate and test](02-skip/)
* [Code coverage](03-cc/)
* [vet incorrect times](04-vet/)

### Core library

* [Wrapping multiple errors](05-wrap-errors/)
* Minor changes to the library
    * [archive/tar and archive/zip security changes](06-archive-sec/)
    * [context's new WithCancelCause](07-context/)
    * [time new layout constants and compare](08-time/)
