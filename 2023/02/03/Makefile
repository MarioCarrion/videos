install:
	go install golang.org/dl/go1.20@latest
	go1.20 download
	mkdir -p bin
	ln -s `go env GOPATH`/bin/go1.20 bin/go
