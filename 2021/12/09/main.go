package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
)

func main() {
	c := mysql.Config{
		User:      "root",
		Passwd:    "password",
		Addr:      "localhost:3306",
		Net:       "tcp",
		DBName:    "dbname",
		ParseTime: true,
	}

	fmt.Println(c.FormatDSN())

	db, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		fmt.Println("sql.Open", err)
		return
	}
	defer func() {
		_ = db.Close()
		fmt.Println("closed")
	}()

	if err := db.PingContext(context.Background()); err != nil {
		fmt.Println("db.PingContext", err)
		return
	}

	row := db.QueryRowContext(context.Background(),
		"SELECT updated_at FROM users WHERE name = ? AND birth_year = ?",
		"mario", 1900)
	if err := row.Err(); err != nil {
		fmt.Println("db.QueryRowContext", err)
		return
	}

	var updatedAt time.Time

	if err := row.Scan(&updatedAt); err != nil {
		fmt.Println("row.Scan", err)
		return
	}

	fmt.Println("updated_at", updatedAt)
}
