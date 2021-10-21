package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/url"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/jackc/pgx/v4/stdlib" // To initialize `pgx` DB driver
)

func main() {
	isAdmin := flag.Bool("admin", false, "Delete all admins")
	birthYear := flag.Int("birthyear", 0, "Maximum Birth Year for deleting users")

	flag.Parse()

	db, err := newDB()
	if err != nil {
		log.Fatalln("Couldn't connect to DB", err)
	}

	//-

	psql := sq.Delete("users").Where("is_admin = ?", *isAdmin)

	if *birthYear > 0 {
		psql = psql.Where("birth_year > ?", *birthYear)
	}

	sql, args, err := psql.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		log.Fatalln("Couldn't create SQL statement", err)
	}

	fmt.Println("query", sql, "args", args)

	//-

	stmt, err := db.PrepareContext(context.Background(), sql)
	if err != nil {
		log.Fatalln("Couldn't prepare", err)
	}

	if _, err := stmt.ExecContext(context.Background(), args...); err != nil {
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
