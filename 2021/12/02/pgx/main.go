package main

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"net/url"
	"os"
	"path"
	"time"

	"github.com/jackc/pgx/v4"
)

func main() {
	conn, err := newConn()
	if err != nil {
		fmt.Println("newDB", err)
		return
	}

	defer func() {
		_ = conn.Close(context.Background())
		fmt.Println("closed")
	}()

	now := time.Now()

	if err := insertsUsers(conn); err != nil {
		fmt.Println("failed", err)
		return
	}

	fmt.Println("total", time.Since(now))
}

type RowSrc struct {
	cr     *csv.Reader
	values []interface{}
	err    error
}

func (r *RowSrc) Next() bool {
	record, err := r.cr.Read()
	if errors.Is(err, io.EOF) {
		return false
	}

	if err != nil {
		r.err = err
		return false
	}

	r.values = make([]interface{}, 2)
	r.values[0] = record[0]
	r.values[1] = record[1]

	return true
}

func (r *RowSrc) Values() ([]interface{}, error) {
	return r.values, r.err
}

func (r *RowSrc) Err() error {
	return r.err
}

func insertsUsers(conn *pgx.Conn) error {
	f, err := os.Open(path.Join("..", "data.csv"))
	if err != nil {
		return fmt.Errorf("os.Open %w", err)
	}

	rowSrc := RowSrc{
		cr: csv.NewReader(f),
	}

	count, err := conn.CopyFrom(context.Background(), pgx.Identifier{"users"}, []string{"first_name", "last_name"}, &rowSrc)
	if err != nil {
		return fmt.Errorf("conn.CopyFrom %w", err)
	}

	fmt.Println("rows", count)

	return nil
}

func newConn() (*pgx.Conn, error) {
	dsn := url.URL{
		Scheme: "postgres",
		Host:   "localhost:5432",
		User:   url.UserPassword("user", "password"),
		Path:   "dbname",
	}

	q := dsn.Query()
	q.Add("sslmode", "disable")

	dsn.RawQuery = q.Encode()

	conn, err := pgx.Connect(context.Background(), dsn.String())
	if err != nil {
		return nil, fmt.Errorf("pgx.Connect%w", err)
	}

	return conn, nil
}
