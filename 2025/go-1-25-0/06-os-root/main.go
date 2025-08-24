package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path"
	"path/filepath"
)

func main() {
	root, err := os.OpenRoot(".")
	if err != nil {
		log.Fatalln("Couldn't open root:", err)
	}

	fullPath := path.Join("parent", "data")
	if err := root.MkdirAll(fullPath, 0750); err != nil { // Similar to `os.MkdirAll`
		log.Fatalln("Couldn't mkdirall:", err)
	}

	newFilePath := path.Join(fullPath, "file.txt")
	if err := root.WriteFile(newFilePath, []byte("hello world!\n"), 0644); err != nil { // Similar to `os.WriteFile`
		log.Fatalln("Couldn't write file:", err)
	}

	newLinkedFilename := path.Join("parent", "linked.txt")
	if err := root.Symlink(newFilePath, newLinkedFilename); err != nil { // Similar to `os.Symlink`
		log.Fatalln("Couldn't link to data:", err)
	}

	fmt.Println("Directory structure (with files created):")
	walkRootDir()

	if err := root.RemoveAll("parent"); err != nil { // Similar to `os.RemoveAll`
		log.Fatalln("Couldn't removeall:", err)
	}

	fmt.Println("\nDirectory structure (with all deleted):")
	walkRootDir()
}

func walkRootDir() {
	filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			log.Fatalln("error accessing a path:", err)
		}

		fmt.Printf("%q\n", path)
		return nil
	})
}
