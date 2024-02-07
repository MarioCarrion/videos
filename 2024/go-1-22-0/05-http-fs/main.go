package main

import (
	"embed"
	"log"
	"net/http"
)

//go:embed files
var f embed.FS

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServerFS(f))
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
