# Backward behavior using GODEBUG

**Note** This example is using the code describing the new [tarinsecurepath security improvement added in Go 1.20](../../../02/03/06-archive-sec).

Excerpt from official release notes:

> Go 1.20 introduced support for rejecting insecure paths in tar and zip archives, controlled by the tarinsecurepath setting and the zipinsecurepath setting. These default to tarinsecurepath=1 and zipinsecurepath=1, preserving the behavior of earlier versions of Go. A future version of Go may change the defaults to tarinsecurepath=0 and zipinsecurepath=0.

## Demo

There are two parts to this demo, first you need to create the unsecure tar file and then run the binary in two different modes:
1. Using the `go:debug` directive and
1. Not using the `go:debug` directive

### Steps

1. Create unsecure tar file: `tar -cvf example.tar example ../../../../LICENSE`
1. Compile binary: `go build -o godebug .`
1. Run binary with and without **GODEBUG=tarinsecurepath=0**
    1. Running it preserving the default Go: `./godebug`
    1. Running it enabling the new behavior: `GODEBUG=tarinsecurepath=0 ./godebug`
1. If the `main` package is modified to include `//go:debug tarinsecurepath=0` then it will override to fail by default
