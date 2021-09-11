package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"sync"

	"golang.org/x/sync/errgroup"
)

/**
Concurrency Pattern using errgroup package

What is in errgroup?
Includes ways to do Synchronization, Error propagation and Context cancelation.

Works for groups of goroutines working on a common task.
*/
func main() {
	wait := waitGroups()
	// wait := errGroup()

	<-wait
}

func waitGroups() <-chan struct{} {
	ch := make(chan struct{}, 1)

	var wg sync.WaitGroup

	for _, file := range []string{"file1.csv", "file2.csv", "file3.csv"} {
		file := file

		wg.Add(1)

		go func() {
			defer wg.Done()

			ch, err := read(file)
			if err != nil {
				fmt.Printf("error reading %v\n", err)
			}

			for line := range ch {
				fmt.Println(line)
			}
		}()
	}

	go func() {
		wg.Wait()

		close(ch)
	}()

	return ch
}

func errGroup() <-chan struct{} {
	ch := make(chan struct{}, 1)

	var g errgroup.Group

	for _, file := range []string{"file1.csv", "file2.csv", "file3.csv"} {
		file := file

		g.Go(func() error {
			ch, err := read(file)
			if err != nil {
				return fmt.Errorf("error reading %w", err)
			}

			for line := range ch {
				fmt.Println(line)
			}

			return nil
		})
	}

	go func() {
		if err := g.Wait(); err != nil {
			fmt.Printf("Error reading files %v", err)
		}

		close(ch)
	}()

	return ch
}

func read(file string) (<-chan []string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("opening file %w", err)
	}

	ch := make(chan []string)

	go func() {
		cr := csv.NewReader(f)

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
