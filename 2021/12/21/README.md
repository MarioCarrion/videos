# Tips for Writing Idiomatic Go - Part 1

Install `golangci-lint`

```
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

## Use gofmt

> Gofmt's style is nobody's favourite, but gofmt is everybody's favourite. - Rob Pike

https://pkg.go.dev/cmd/gofmt

## Receiver Names

Use one or two letter abbreviation for naming a method receiver matching its
type. Avoid identifiers from another languages like `me`, `this` or `self`.

Keep that name consistent across all methods, using `revive` should help you doing that.

## Naming packages 

> Name your packages after what they provide, not what they contain. - [Dave Cheney](https://dave.cheney.net/2019/01/08/avoid-package-names-like-base-util-or-common)

Make them short, unique and representative. Avoid generic ones like `common` and `util`.

## Group imports by their _origin_

It allows quickly to determine where the packages are coming from, what packages are from the standard library, _local_ and third party.

Using `gci` should help you doing that.

## Use short names for variables with limited scope

> The basic rule: the further from its declaration that a name is used, the
> more descriptive the name must be. For a method receiver, one or two letters
> is sufficient. Common variables such as loop indices and readers can be a
> single letter (i, r). More unusual things and global variables need more
> descriptive names.

[Example](https://github.com/MarioCarrion/todo-api-microservice-example/blob/1a59011bc6c520eb45997acc0c2cfff7fe39f8dd/internal/elasticsearch/task.go#L153)

Using `varnamelen` should help you doing that.
