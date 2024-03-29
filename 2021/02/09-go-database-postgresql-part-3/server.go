package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Name struct {
	NConst    string `json:"nconst"`
	Name      string `json:"name"`
	BirthYear string `json:"birthYear"`
	DeathYear string `json:"deathYear"`
}

type Error struct {
	Message string `json:"error"`
}

func main() {
	router := mux.NewRouter()

	renderJSON := func(w http.ResponseWriter, val interface{}, statusCode int) {
		w.WriteHeader(statusCode)
		_ = json.NewEncoder(w).Encode(val)
	}

	//---- Part 1

	dbSQL, err := NewPostgreSQLsql()
	if err != nil {
		log.Fatalf("Could not initialize Database connection using sqlx %s", err)
	}
	defer dbSQL.Close()

	router.HandleFunc("/names/sql/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

		name, err := dbSQL.FindByNConst(id)
		if err != nil {
			renderJSON(w, &Error{Message: err.Error()}, http.StatusInternalServerError)
			return
		}

		renderJSON(w, &name, http.StatusOK)
	})

	//-

	dbSQLX, err := NewPostgreSQLsqlx()
	if err != nil {
		log.Fatalf("Could not initialize Database connection using sqlx %s", err)
	}
	defer dbSQLX.Close()

	router.HandleFunc("/names/sqlx/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

		name, err := dbSQLX.FindByNConst(id)
		if err != nil {
			renderJSON(w, &Error{Message: err.Error()}, http.StatusInternalServerError)
			return
		}

		renderJSON(w, &name, http.StatusOK)
	})

	//-

	pgxDB, err := NewPostgreSQLpgx()
	if err != nil {
		log.Fatalf("Could not initialize Database connection using pgx %s", err)
	}
	defer pgxDB.Close()

	router.HandleFunc("/names/pgx/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

		name, err := pgxDB.FindByNConst(id)
		if err != nil {
			renderJSON(w, &Error{Message: err.Error()}, http.StatusInternalServerError)
			return
		}

		renderJSON(w, &name, http.StatusOK)
	})

	//---- Part 2

	dbGorm, err := NewPostgreSQLgorm()
	if err != nil {
		log.Fatalf("Could not initialize Database connection using gorm %s", err)
	}

	router.HandleFunc("/names/gorm/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

		name, err := dbGorm.FindByNConst(id)
		if err != nil {
			renderJSON(w, &Error{Message: err.Error()}, http.StatusInternalServerError)
			return
		}

		renderJSON(w, &name, http.StatusOK)
	})

	//-

	dbBoiler, err := NewPostgreSQLboiler()
	if err != nil {
		log.Fatalf("Could not initialize Database connection using boiler %s", err)
	}
	defer dbBoiler.Close()

	router.HandleFunc("/names/sqlboiler/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

		name, err := dbBoiler.FindByNConst(id)
		if err != nil {
			renderJSON(w, &Error{Message: err.Error()}, http.StatusInternalServerError)
			return
		}

		renderJSON(w, &name, http.StatusOK)
	})

	//---- Part 3

	dbSquirrel, err := NewPostgreSQLsquirrel()
	if err != nil {
		log.Fatalf("Could not initialize Database connection using squirrel %s", err)
	}
	defer dbSquirrel.Close()

	router.HandleFunc("/names/squirrel/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

		name, err := dbSquirrel.FindByNConst(id)
		if err != nil {
			renderJSON(w, &Error{Message: err.Error()}, http.StatusInternalServerError)
			return
		}

		renderJSON(w, &name, http.StatusOK)
	})

	//-

	dbSQLC, err := NewPostgreSQLC()
	if err != nil {
		log.Fatalf("Could not initialize Database connection using sqlc %s", err)
	}
	defer dbSQLC.Close()

	router.HandleFunc("/names/sqlc/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

		name, err := dbSQLC.FindByNConst(id)
		if err != nil {
			renderJSON(w, &Error{Message: err.Error()}, http.StatusInternalServerError)
			return
		}

		renderJSON(w, &name, http.StatusOK)
	})

	//-

	fmt.Println("Starting server :8080")

	srv := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
