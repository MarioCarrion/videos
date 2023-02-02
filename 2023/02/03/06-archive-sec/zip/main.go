package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

func main() {
	f, err := os.Open(path.Join("..", "example.zip"))

	if err != nil {
		log.Fatalf("open failed %v\n", err)
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		log.Fatalf("f.Stat failed %v\n", err)
	}

	r, err := zip.NewReader(f, fi.Size())
	if err != nil {
		log.Fatalf("NewReader failed %v\n", err)
	}

	for _, f := range r.File {
		fmt.Printf("Contents of %s:\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}

		if _, err := io.Copy(os.Stdout, rc); err != nil {
			log.Fatalf("io.Copy %v\n", err)
		}

		rc.Close()
		fmt.Println()
	}
}
