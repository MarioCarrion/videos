# Go 1.18: What is new?

This directory includes examples to (some) features I consider worth mentioning included in Go 1.18.0.

* [@Golang Twitter announcement](https://twitter.com/golang/status/1503787326060875782)
* [Download](https://go.dev/dl/#go1.18)
* [Release notes](https://go.dev/doc/go1.18)
* Docker Image for Alpine
  * [3.14](https://hub.docker.com/layers/golang/library/golang/1.18-alpine3.14/images/sha256-b8eeab34b03b7b97988e42b75223c764d26ed4b5b5231c6c14f2695c8cd86847?context=explore): `docker pull golang:1.18-alpine3.14`
  * [3.15](https://hub.docker.com/layers/golang/library/golang/1.18-alpine3.15/images/sha256-2d37a990a83e584b1def08811a104146037e5813a3a24f30c26980562503e71c?context=explore): `docker pull golang:1.18-alpine3.15`
* [Homebrew Formula](https://github.com/Homebrew/homebrew-core/pull/91369)

`NOTE` For `goenv` users: make sure you reinstall the binaries you use for your editor, like `gopls` and `gopls` otherwise you will run into issues, using [`direnv`](https://mariocarrion.com/2020/11/20/golang-go-tool-direnv.html) the tools will be _sandboxed_ to this folder.
