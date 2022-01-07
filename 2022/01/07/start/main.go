package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// User represents the user for this application
type User struct {
	// The name for this user
	Name string `json:"name"`

	// The birth year for this user
	BirthYear int `json:"birth_year"`
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		users := []User{
			{"Mario", 1990},
			{"Wario", 1980},
		}

		res, _ := json.Marshal(&users)

		_, _ = w.Write(res)
	}).Methods(http.MethodGet)

	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		// XXX: Imagine validation is implemented here

		var user User

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			_, _ = w.Write([]byte("decoding failed"))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, _ = w.Write([]byte(fmt.Sprintf("created %+v", user)))
	}).Methods(http.MethodPost)

	s := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Fatal(s.ListenAndServe())
}
