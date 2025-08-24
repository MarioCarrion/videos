# New experimental `encoding/json/v2` package

Excerpt from [official release notes](https://go.dev/doc/go1.25#json_v2):

> Go 1.25 includes a new, experimental JSON implementation, which can be
> enabled by setting the environment variable `GOEXPERIMENT=jsonv2` at build
> time.
>
> When enabled, two new packages are available:
>
> * The `encoding/json/v2` package is a major revision of the encoding/json
> package.
> * The `encoding/json/jsontext` package provides lower-level processing of JSON
> syntax.
>
> In addition, when the "jsonv2" `GOEXPERIMENT` is enabled:
>
> The `encoding/json` package uses the new JSON implementation. Marshaling and
> unmarshaling behavior is unaffected, but the text of errors returned by
> package function may change.
> The `encoding/json` package contains a number of new options which may be
> used to configure the marshaler and unmarshaler.

Make sure to review the [new v2 behavior](https://pkg.go.dev/encoding/json@master#hdr-Migrating_to_v2),
because by default it's no backwards compatible without applying a set of
options beforehand when using the `v2` package; using `v1` will continue
behaving the same way.

## Running

Use `GOEXPERIMENT=jsonv2 go build -o main . && ./main`
