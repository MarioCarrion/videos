# `go vet`

Excerpt from [official release notes](https://go.dev/doc/go1.25#vet):

> The go vet command includes new analyzers:
>
> * `waitgroup`, which reports misplaced calls to `sync.WaitGroup.Add`; and
> * `hostport`, which reports uses of `fmt.Sprintf("%s:%d", host, port)` to construct
> addresses for `net.Dial`, as these will not work with IPv6; instead it suggests
> using `net.JoinHostPort`.

See the two examples in this folder and run `go vet ./...`:

* [waitgroup](waitgroup/)
* [hostport](hostport/)
