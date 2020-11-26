# Go Tool: counterfeiter

[Published Video](https://youtu.be/t2r2u2vd0ig)

1. Install [`direnv`](https://github.com/direnv/direnv)
1. Then:

```
go mod download
go install github.com/maxbrunsfeld/counterfeiter/v6
go test ./...
```

## Problem

Implement a service for persisting Books using ISBN, details related to that book will be accessible via a different third party service.
