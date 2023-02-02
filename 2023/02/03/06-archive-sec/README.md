# archive/tar and archive/zip security changes

Excerpt from official release notes:

> archive/tar

> When the GODEBUG=tarinsecurepath=0 environment variable is set, Reader.Next method will now return the error ErrInsecurePath for an entry with a file name that is an absolute path, refers to a location outside the current directory, contains invalid characters, or (on Windows) is a reserved name such as NUL. A future version of Go may disable insecure paths by default.

> archive/zip

> When the GODEBUG=zipinsecurepath=0 environment variable is set, NewReader will now return the error ErrInsecurePath when opening an archive which contains any file name that is an absolute path, refers to a location outside the current directory, contains invalid characters, or (on Windows) is a reserved names such as NUL. A future version of Go may disable insecure paths by default.

## Demo

## TAR

1. Create unsecure tar file: `tar -cvf example.tar example ../../../../LICENSE`
1. Compile binary: `cd tar && go build .`
1. Run binary with and without **GODEBUG=tarinsecurepath=0**
    1. `./tar`
    1. `GODEBUG=tarinsecurepath=0 ./tar`

## ZIP

1. Create unsecure tar file: `zip -v -r example.zip example/ ../../../../LICENSE`
1. Compile binary: `cd zip && go build .`
1. Run binary with and without **GODEBUG=zipinsecurepath=0**
    1. `./zip`
    1. `GODEBUG=zipinsecurepath=0 ./zip`
