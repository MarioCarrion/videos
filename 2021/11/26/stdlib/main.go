package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type User struct {
	Name      *string
	BirthYear *int64
}

func main() {
	db, err := newDB()
	if err != nil {
		fmt.Println("newDB", err)
		return
	}

	defer func() {
		_ = db.Close()
		fmt.Println("closed")
	}()

	newName := func(s string) *string {
		return &s
	}

	newBirthYear := func(b int64) *int64 {
		return &b
	}

	if err := insertUsers(db, []User{
		{
			Name:      newName("mario1"),
			BirthYear: newBirthYear(1900),
		},
		{
			Name:      nil,
			BirthYear: newBirthYear(1901),
		},
	}); err != nil {
		fmt.Println("insertUsers", err)
	}
}

func insertUsers(db *sql.DB, users []User) error {
	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		return fmt.Errorf("db.BeginTx %w", err)
	}
	defer func() {
		_ = tx.Rollback()
	}()

	for _, user := range users {
		_, err := tx.ExecContext(context.Background(), "INSERT INTO users(name, birth_year) VALUES($1, $2)", user.Name, user.BirthYear)
		if err != nil {
			return fmt.Errorf("tx.ExecContext %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("tx.Commit %w", err)
	}

	return nil
}

func newDB() (*sql.DB, error) {
	dsn := url.URL{
		Scheme: "postgres",
		Host:   "localhost:5432",
		User:   url.UserPassword("user", "password"),
		Path:   "dbname",
	}

	q := dsn.Query()
	q.Add("sslmode", "disable")

	dsn.RawQuery = q.Encode()

	db, err := sql.Open("pgx", dsn.String())
	if err != nil {
		return nil, fmt.Errorf("sql.Open %w", err)
	}

	return db, nil
}
