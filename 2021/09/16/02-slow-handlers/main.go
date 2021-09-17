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
		defer r.Body.Close()

		fmt.Println("written")

		fmt.Fprint(w, "Hello ", string(body))
	})

	slowHandler := func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("io.ReadAll", err)

			return
		}
		defer r.Body.Close()

		fmt.Println("Sleeping...")

		time.Sleep(3 * time.Second)

		fmt.Fprintf(w, "H.e.l.l.o %s", string(body))
	}

	router.Handler(http.MethodPost, "/slow",
		http.TimeoutHandler(http.HandlerFunc(slowHandler), 2*time.Second, "Request took too long"))

	s := &http.Server{
		Addr:        ":8080",
		Handler:     router,
		ReadTimeout: 100 * time.Millisecond,
	}

	log.Fatal(s.ListenAndServe())
}
