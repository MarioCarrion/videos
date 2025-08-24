# Updating

Starting with Go **1.21**, there's no need to install a version explicitly; instead, editing the `go.mod` to indicate the version to use should suffice.

For example, assuming you have a `go.mod` where **1.24.6** is being used, the content of this file in its most basic representation would be:

```
module github.com/MarioCarrion/videos/2025/go-1-25-0/00-updating

go 1.24.6
```

If you wanted this `go.mod` to start using **1.25.0**, you will do:

```
go mod edit -go=1.25.0
```

That's it, you should see the new updated `go.mod`:

```
module github.com/MarioCarrion/videos/2025/go-1-25-0/00-updating

go 1.25.0
```
