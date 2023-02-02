package main

import (
	"archive/tar"
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

func main() {
	f, err := os.Open(path.Join("..", "example.tar"))

	if err != nil {
		log.Fatalf("open failed %v\n", err)
	}
	defer f.Close()

	reader := tar.NewReader(f)

	for {
		hdr, err := reader.Next()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("reader.Next %v\n", err)
		}

		fmt.Printf("Contents of %s:\n", hdr.Name)
		if _, err := io.Copy(os.Stdout, reader); err != nil {
			log.Fatalf("io.Copy %v\n", err)
		}

		fmt.Println()
	}
}
