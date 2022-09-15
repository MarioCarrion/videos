# Go Vulnerability in Third Party library

* [Go Package Discovery](https://pkg.go.dev/gopkg.in/yaml.v3?tab=versions)

Compile the binary:

```
go build .
```

Run `govulncheck thirdpartylib`, it will return a vulnerability.

This report is similar to other tools/services such as [Snyk](https://snyk.io/):

![Snyk Report](snyk_report.png)
