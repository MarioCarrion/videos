package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	recordsC, err := readCSV("file1.csv")
	if err != nil {
		log.Fatalf("Could not read csv %v", err)
	}

	for val := range sanitize(titleize(recordsC)) {
		fmt.Printf("%v\n", val)
	}
}

// Read values
func readCSV(file string) (<-chan []string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("opening file %w", err)
	}

	ch := make(chan []string)

	go func() {
		cr := csv.NewReader(f)
		cr.FieldsPerRecord = 3

		for {
			record, err := cr.Read()
			if errors.Is(err, io.EOF) {
				close(ch)

				return
			}

			ch <- record
		}
	}()

	return ch, nil
}

// Remove "invalid" records
func sanitize(strC <-chan []string) <-chan []string {
	ch := make(chan []string)

	go func() {
		for val := range strC {
			if len(val[0]) > 3 {
				fmt.Println("skipped ", val)
				continue
			}

			ch <- val
		}

		close(ch)
	}()

	return ch
}

// Modify received values
func titleize(strC <-chan []string) <-chan []string {
	ch := make(chan []string)

	go func() {
		for val := range strC {
			val[0] = strings.Title(val[0])
			val[1], val[2] = val[2], val[1]

			ch <- val
		}

		close(ch)
	}()

	return ch
}
