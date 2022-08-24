# Go 1.19: What is new?

Go 1.19.0 was released on August 2nd, 2022; in this directory you will find examples of some new features I consider worth mentioning added in this release.

* [@Golang Twitter announcement](https://twitter.com/golang/status/1554515292390408192)
* [Download](https://go.dev/dl/#go1.19)
* [Release notes](https://go.dev/doc/go1.18)
* [New API Changes for 1.19](https://github.com/golang/go/issues/53310)

* Other ways to use Go:
  * Docker Images
    * [All](https://hub.docker.com/_/golang?tab=tags&page=1&name=1.19)
    * [Alpine 3.16]( https://hub.docker.com/layers/golang/library/golang/1.19.0-alpine3.16/images/sha256-78767a599f5758e3fd60fd065de9dc40aec1ad3b86ded2bddc92640a465cb515?context=explore ): `docker pull golang:1.19.0-alpine3.16`
  * [Homebrew Formula PR](https://github.com/Homebrew/homebrew-core/pull/107165)

## New features

### Tools

* [Doc Comments](01-comments/)
* [vet](02-vet/)

### Core library

* [New atomic types](03-atomic-types/)
* Minor changes to the library
    * [flag.TextVar](04-flag-textvar/)

### PATH lookups

Command and LookPath no longer allow results from a PATH search to be found relative to the current directory. This removes a common source of security problems but may also break existing programs that depend on using, say, exec.Command("prog") to run a binary named prog (or, on Windows, prog.exe) in the current directory. See the os/exec package documentation for information about how best to update such programs.

On Windows, Command and LookPath now respect the NoDefaultCurrentDirectoryInExePath environment variable, making it possible to disable the default implicit search of “.” in PATH lookups on Windows systems.

