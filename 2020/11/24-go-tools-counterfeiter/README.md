# Go Tool: counterfeiter

[Published Video](https://youtu.be/ENqwq64TsDk)

1. Install [`direnv`](https://github.com/direnv/direnv)
1. Then:

```
go mod download
go install github.com/maxbrunsfeld/counterfeiter/v6
go test ./...
```

## Problem

Implement a service for persisting Books using ISBN, details related to that book will be accessible via a different third party service.
