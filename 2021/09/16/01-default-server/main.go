package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.POST("/hello", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("io.ReadAll", err)

			return
		}

		fmt.Println("written")

		fmt.Fprint(w, "Hello ", string(body))
	})

	s := &http.Server{
		Addr:        ":8080",
		Handler:     router,
		ReadTimeout: 500 * time.Millisecond,
	}

	log.Fatal(s.ListenAndServe())
}
