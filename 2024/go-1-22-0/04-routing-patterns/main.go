package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/" {
			http.NotFound(w, req)
			return
		}

		fmt.Fprintf(w, "Welcome to the home page!\n")
	})

	mux.HandleFunc("GET /{id}", func(w http.ResponseWriter, req *http.Request) {
		id := req.PathValue("id")
		fmt.Fprintf(w, "Your id is: %v\n", id)
	})

	mux.HandleFunc("GET /{id}/value", func(w http.ResponseWriter, req *http.Request) {
		id := req.PathValue("id")
		fmt.Fprintf(w, "Your id is: %v\n", id)
	})

	// XXX: Causes a panic
	// mux.HandleFunc("GET /val1/{val2}", func(w http.ResponseWriter, req *http.Request) {
	// 	val2 := req.PathValue("val2")
	// 	fmt.Fprintf(w, "Your child id is: %v\n", val2)
	// })

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		log.Fatal(err)
	}
}
