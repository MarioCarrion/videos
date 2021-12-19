package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"

	"golang.org/x/sync/errgroup"
)

func main() {
	wait := errGroupWithErrChannel() // Using a function that returns an error channel
	// wait := errGroupWithType() // Using a type that saves Error as state

	<-wait
}

func errGroupWithErrChannel() <-chan struct{} {
	waitC := make(chan struct{}, 1)

	var g errgroup.Group

	for _, file := range []string{"file1.csv", "file2.csv", "file3.csv"} {
		file := file

		g.Go(func() error {
			readC, errC, err := readWithErrChannel(file)
			if err != nil {
				return fmt.Errorf("error reading %w", err)
			}

			loop := true

			for loop {
				select {
				case line, ok := <-readC:
					if ok {
						fmt.Println(line)
					} else {
						loop = false
					}
				case err = <-errC:
					loop = false
				}
			}

			return err
		})
	}

	go func() {
		if err := g.Wait(); err != nil {
			fmt.Printf("Error reading: files %v\n", err)
		}

		close(waitC)
	}()

	return waitC
}

func readWithErrChannel(file string) (<-chan []string, <-chan error, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, nil, fmt.Errorf("opening file %w", err)
	}

	ch := make(chan []string)
	errC := make(chan error, 1)

	go func() {
		defer func() {
			close(errC)
			close(ch)
		}()

		cr := csv.NewReader(f)

		for {
			record, err := cr.Read()
			if errors.Is(err, io.EOF) {
				return
			}

			if err != nil {
				errC <- err
				return
			}

			ch <- record
		}
	}()

	return ch, errC, nil
}

//-

type Reader struct {
	cr    *csv.Reader
	Error error
}

func newReader(file string) (*Reader, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("opening file %w", err)
	}

	return &Reader{
		cr: csv.NewReader(f),
	}, nil
}

func (r *Reader) Read() <-chan []string {
	ch := make(chan []string)

	go func() {
		defer close(ch)

		for {
			record, err := r.cr.Read()
			if errors.Is(err, io.EOF) {
				return
			}

			if err != nil {
				r.Error = err
				return
			}

			ch <- record
		}
	}()

	return ch
}

func errGroupWithType() <-chan struct{} {
	waitC := make(chan struct{}, 1)

	var g errgroup.Group

	for _, file := range []string{"file1.csv", "file2.csv", "file3.csv"} {
		file := file

		g.Go(func() error {
			reader, err := newReader(file)
			if err != nil {
				return fmt.Errorf("error reading %w", err)
			}

			for line := range reader.Read() {
				fmt.Println(line)
			}

			return reader.Error
		})
	}

	go func() {
		if err := g.Wait(); err != nil {
			fmt.Printf("Error reading: files %v\n", err)
		}

		close(waitC)
	}()

	return waitC
}
