# Container-aware `GOMAXPROCS`

Excerpt from [official release notes](https://go.dev/doc/go1.25#container-aware-gomaxprocs):

> The default behavior of the `GOMAXPROCS` has changed. In prior versions of Go,
> `GOMAXPROCS` defaults to the number of logical CPUs available at startup
> (`runtime.NumCPU`). Go 1.25 introduces two changes:
>
> On Linux, the runtime considers the CPU bandwidth limit of the cgroup
> containing the process, if any. If the CPU bandwidth limit is lower than the
> number of logical CPUs available, `GOMAXPROCS` will default to the lower limit.
> In container runtime systems like Kubernetes, cgroup CPU bandwidth limits
> generally correspond to the "CPU limit" option. The Go runtime does not
> consider the “CPU requests” option.
>
> On all OSes, the runtime periodically updates `GOMAXPROCS` if the number of
> logical CPUs available or the cgroup CPU bandwidth limit change.
>
> Both of these behaviors are automatically disabled if `GOMAXPROCS` is set
> manually via the `GOMAXPROCS` environment variable or a call to
> `runtime.GOMAXPROCS`. They can also be disabled explicitly with the `GODEBUG`
> settings `containermaxprocs=0` and `updatemaxprocs=0`, respectively.
>
> In order to support reading updated cgroup limits, the runtime will keep
> cached file descriptors for the cgroup files for the duration of the process
> lifetime.

Blog post by the Go team: [Container-aware GOMAXPROCS](https://go.dev/blog/container-aware-gomaxprocs)

## Running

Try using the example using Docker with default values:

```
docker run --rm -v $(pwd):/pwd golang:1.25-bookworm go run /pwd/main.go
```

Then set the CPUs explicitly:

```
docker run --rm --cpus=4 -v $(pwd):/pwd golang:1.25-bookworm go run /pwd/main.go
```

The printed values will be different because of the `--cpus` argument.
