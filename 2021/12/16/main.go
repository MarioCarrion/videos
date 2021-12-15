package main

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

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

	db, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		fmt.Println("sql.Open", err)
		return
	}
	defer func() {
		_ = db.Close()
		fmt.Println("closed")
	}()

	// Example 1: using the csv file as-is

	mysql.RegisterLocalFile("data.csv")

	_, err = db.ExecContext(context.Background(),
		`LOAD DATA LOCAL INFILE 'data.csv' INTO TABLE users
		FIELDS TERMINATED BY ','
		ENCLOSED BY '"'
		(first_name, last_name)
		`,
	)
	if err != nil {
		fmt.Println("db.ExecContext", err)
		return
	}

	// Example 2: augmenting the csv file

	mysql.RegisterReaderHandler("data", func() io.Reader {
		file, err := os.Open("data.csv")
		if err != nil {
			fmt.Println("os.Open", err)
			return nil
		}
		defer func() {
			_ = file.Close()
		}()

		buffer := bytes.Buffer{}

		csvWriter := csv.NewWriter(&buffer)

		csvReader := csv.NewReader(file)
		for {
			record, err := csvReader.Read()
			if errors.Is(err, io.EOF) {
				break
			}

			if err != nil {
				fmt.Println("csvR.Read", err)
				return nil
			}

			record[0] = strings.ToUpper(record[0])

			csvWriter.Write(record)
		}

		csvWriter.Flush()

		return &buffer
	})

	_, err = db.ExecContext(context.Background(),
		`LOAD DATA LOCAL INFILE 'Reader::data' INTO TABLE users
		FIELDS TERMINATED BY ','
		ENCLOSED BY '"'
		(first_name, last_name)
		`,
	)
	if err != nil {
		fmt.Println("db.ExecContext", err)
		return
	}
}
