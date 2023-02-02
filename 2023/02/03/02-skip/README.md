# Skip

Excerpt from official release notes:

> The go generate command now accepts -skip "pattern" to skip //go:generate directives matching "pattern".

> The go test command now accepts -skip "pattern" to skip tests, subtests, or examples matching "pattern".

## Example

### Using `go generate`

Make sure you install `stringer` via: `go install golang.org/x/tools/cmd/stringer@v0.5.0`, modify one of the types, for example add a new country to the `Country` type and then run the following:

```
go generate -skip="Country" github.com/MarioCarrion/videos/2023/02/03/02-skip
```

That **won't** generate the expected code.


### Using `go test`

For example to skip `Integer` -related tests:

```
go test -skip="Integer" -v ./...
```
