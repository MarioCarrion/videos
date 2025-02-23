# Directory-limited filesystem access

Excerpt from [official release notes](https://go.dev/doc/go1.24#directory-limited-filesystem-access):

> The new `os.Root` type provides the ability to perform filesystem operations
> within a specific directory.
>
>
> The `os.OpenRoot` function opens a directory and returns an `os.Root`.
> Methods on `os.Root` operate within the directory and do not permit paths
> that refer to locations outside the directory, including ones that follow
> symbolic links out of the directory. The methods on `os.Root` mirror most of
> the file system operations available in the os package, including for example
> `os.Root.Open`, `os.Root.Create`, `os.Root.Mkdir`, and `os.Root.Stat`,
