# Go 1.25.0: What's new?

* [Video](https://youtu.be/iK4evFHCzWw)

Go **1.25.0** was released on _August 12th, 2025_. This directory includes examples of some new features added in this release.

* [@Golang X announcement](https://x.com/golang/status/1955388473185030597)
* [Download](https://go.dev/dl/#go1.25.0)
* [Release notes](https://go.dev/doc/go1.25)
* [Official Blog Post](https://go.dev/blog/go1.25)
* [1.25 Milestone on Github](https://github.com/golang/go/milestone/364?closed=1)

* Other ways to get Go 1.25.0:
  * Docker Images:
    * [All](https://hub.docker.com/_/golang/tags?name=1.25)
    * [Debian Bookworm](https://hub.docker.com/layers/library/golang/1.25-bookworm/images/sha256-dfaaac68d880dafabeac289427956be7f71f21431e23446fe0faad79333ba59d) `docker pull golang:1.25-bookworm`
    * [Debian Trixie](https://hub.docker.com/layers/library/golang/1.25-trixie/images/sha256-69e51132953b45d446a436c43f7bbf93072ce568268e5b401be3823118118558) `docker pull golang:1.25-trixie`
    * [Alpine 3.22](https://hub.docker.com/layers/library/golang/1.25-alpine3.22/images/sha256-5dffbd69b37b8e9fc35d94817e94bc7bf4fd047147e3c5353d357bd3492992bf): `docker pull golang:1.25-alpine3.22`
  * [Homebrew Formula Pull Request](https://github.com/Homebrew/homebrew-core/pull/226636)

# New features in Go 1.25.0

* [Updating](00-updating/)

## Tools

* [`go doc`](01-go-doc/)
* [`go vet`](02-go-vet/)

## Runtime

* [Container-aware GOMAXPROCS](03-gomaxprocs/)

## Standard library 

* [New `testing/synctest` package](04-synctest/)
* [New experimental `encoding/json/v2` package](05-json-v2/)
* [New `os.Root` functions](06-os-root/)
