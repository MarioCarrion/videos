package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

func main() {
	root, err := os.OpenRoot(".")
	if err != nil {
		log.Fatalln("Couldn't open root:", err)
	}

	dirname := "test"

	if err := root.Mkdir(dirname, 0750); err != nil { // similar to `os.Mkdir`
		log.Fatalln("Couldn't make dir:", err)
	}

	//- Writing

	filename := path.Join(dirname, "example.txt")

	fwrite, err := root.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644) // similar to `os.OpenFile`
	if err != nil {
		log.Fatalln("Couldn't open file for writing:", err)
	}

	if _, err := fwrite.Write([]byte("hello world!\n")); err != nil {
		log.Fatalln("Couldn't write to file:", err)
	}

	if err := fwrite.Close(); err != nil {
		log.Fatalln("Couldn't close writing file:", err)
	}

	//- Reading

	fread, err := root.Open(filename) // similar to `os.Open`
	if err != nil {
		log.Fatalln("Couldn't open file for reading:", err)
	}

	content, err := io.ReadAll(fread)
	if err != nil {
		log.Fatalln("Couldn't reading file:", err)
	}

	fmt.Printf("%s", string(content))

	if err := fread.Close(); err != nil {
		log.Fatalln("Couldn't close reading file:", err)
	}

	//- Removing

	if err := root.Remove(filename); err != nil { // similar to `os.Remove`
		log.Fatalln("Couldn't remove file:", err)
	}

	if err := root.Remove(dirname); err != nil { // similar to `os.Remove`
		log.Fatalln("Couldn't remove directory:", err)
	}
}
