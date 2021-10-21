package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/url"

	_ "github.com/jackc/pgx/v4/stdlib" // To initialize `pgx` DB driver
)

// Passing in an `id` like:
//
// "33ca99ce-f1d1-46c2-b26e-a0ff45011b18' OR ''='"
//
// The value "33ca..." is irrelevant because the `OR` is used instead.

func main() {
	var id string

	flag.StringVar(&id, "id", "", "User ID to delete")

	flag.Parse()

	if id == "" {
		fmt.Println("id is blank, exiting")
		return
	}

	db, err := newDB()
	if err != nil {
		log.Fatalln("Couldn't connect to DB", err)
	}

	//-

	query := fmt.Sprintf("DELETE FROM users WHERE id = '%s'", id)

	if _, err := db.Exec(query); err != nil {
		log.Fatalln("Couldn't delete", err)
	}
}

func newDB() (*sql.DB, error) {
	dsn := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword("user", "password"),
		Host:   "localhost:5432",
		Path:   "dbname",
	}

	q := dsn.Query()
	q.Add("sslmode", "disable")

	dsn.RawQuery = q.Encode()

	db, err := sql.Open("pgx", dsn.String())
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}

	if err := db.PingContext(context.Background()); err != nil {
		return nil, fmt.Errorf("db.Ping: %w", err)
	}

	return db, nil
}
