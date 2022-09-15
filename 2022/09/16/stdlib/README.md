# Go Vulnerability in Standard library

Using docker you can generate two different versions of the binary to trigger errors.

## Go 1.19.0

Includes an error

```
docker run --rm \
  -v "$PWD":/govulncheck -w /govulncheck \
  -e GOOS=darwin -e GOARCH=amd64 \
  golang:1.19.0-buster go build -o go119 .
```

Run `govulncheck go1191`, it will return a vulnerability.

## Go 1.19.1

Fixes that said error

```
docker run --rm \
  -v "$PWD":/govulncheck -w /govulncheck \
  -e GOOS=darwin -e GOARCH=amd64 \
  golang:1.19.1-buster go build -o go1191 .
```

Run `govulncheck go1191`, it won't return vulnerabilities.
