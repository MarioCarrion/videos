# Accept -C

Excerpt from official release notes:

> The go subcommands now accept -C <dir> to change directory to <dir> before performing the command, which may be useful for scripts that need to execute commands in multiple different modules

## Example

To build [example1](example1/):

```
go build -C example1 github.com/MarioCarrion/videos/2023/02/03/01-accept-c/example1
```

To build [example2](example2/):

```
go build -C example2 github.com/MarioCarrion/videos/2023/02/03/01-accept-c/example2
```
