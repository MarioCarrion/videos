package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/url"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib" // To initialize `pgx` DB driver
)

func main() {
	isAdmin := flag.Bool("admin", false, "Delete all admins")
	birthYear := flag.Int("birthyear", 0, "Maximum Birth Year for deleting users")

	flag.Parse()

	conn, err := newConn()
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

	if _, err := conn.Prepare(context.Background(), "deleteWithCondition", sql); err != nil {
		log.Fatalln("Couldn't prepare", err)
	}

	if _, err := conn.Exec(context.Background(), "deleteWithCondition", args...); err != nil {
		log.Fatalln("Couldn't delete", err)
	}
}

func newConn() (*pgx.Conn, error) {
	dsn := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword("user", "password"),
		Host:   "localhost:5432",
		Path:   "dbname",
	}

	q := dsn.Query()
	q.Add("sslmode", "disable")

	dsn.RawQuery = q.Encode()

	conn, err := pgx.Connect(context.Background(), dsn.String())
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}

	if err := conn.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("db.Ping: %w", err)
	}

	return conn, nil
}
