# vet incorrect times

Excerpt from official release notes:

> The vet tool now reports use of the time format 2006-02-01 (yyyy-dd-mm) with Time.Format and time.Parse. This format does not appear in common date standards, but is frequently used by mistake when attempting to use the ISO 8601 date format (yyyy-mm-dd).

## Demo

If you're using `gopls` together with your editor it should report the warning, you can explicitly call it via:


```
go vet main.go
```

It should output something like this:


```
# command-line-arguments
./main.go:10:26: 2006-02-01 should be 2006-01-02
```
