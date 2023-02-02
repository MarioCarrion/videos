# Code Coverage

Excerpt from official release notes:

> The go build, go install, and other build-related commands now support a -cover flag that builds the specified target with code coverage instrumentation. This is described in more detail in the Cover section below.

## Example

```
go build -cover -o main
mkdir -o cover
GOCOVERDIR=cover ./main
go tool covdata textfmt -i=cover -o profile.txt
go tool cover -html=profile.txt
```
