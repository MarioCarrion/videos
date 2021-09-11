package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sync"
)

// What is the Fan-In concurrency pattern?
// Consolidation of multiple channels into one channel by multiplexing each recieved value.
func main() {
	ch1, err := readCSV("file1.csv")
	if err != nil {
		panic(fmt.Errorf("Could not read file1 %v\n", err))
	}

	ch2, err := readCSV("file2.csv")
	if err != nil {
		panic(fmt.Errorf("Could not read file2 %v\n", err))
	}

	//-

	exit := make(chan struct{})

	// chM := merge1(ch1, ch2)
	chM := merge2(ch1, ch2)

	go func() {
		for v := range chM {
			fmt.Println(v)
		}

		close(exit)
	}()

	<-exit

	fmt.Println("All completed, exiting")
}

func merge1(cs ...<-chan []string) <-chan []string {
	var wg sync.WaitGroup

	out := make(chan []string)

	send := func(c <-chan []string) {
		for n := range c {
			out <- n
		}

		wg.Done()
	}

	wg.Add(len(cs))

	for _, c := range cs {
		go send(c)
	}

	go func() {
		wg.Wait()

		close(out)
	}()

	return out
}

func merge2(cs ...<-chan []string) <-chan []string {
	chans := len(cs)
	wait := make(chan struct{}, chans)

	out := make(chan []string)

	send := func(c <-chan []string) {
		defer func() { wait <- struct{}{} }()

		for n := range c {
			out <- n
		}
	}

	for _, c := range cs {
		go send(c)
	}

	go func() {
		for range wait {
			chans--
			if chans == 0 {
				break
			}
		}

		close(out)
	}()

	return out
}

func readCSV(file string) (<-chan []string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("opening file %v\n", err)
	}

	ch := make(chan []string)

	cr := csv.NewReader(f)

	go func() {
		for {
			record, err := cr.Read()
			if err == io.EOF {
				close(ch)

				return
			}

			ch <- record
		}
	}()

	return ch, nil
}
