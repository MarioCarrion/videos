package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

// What is the Fan-Out concurrency pattern?
// Breakup of one channel into multiple ones by distributing each value.
func main() {
	ch1, err := readCSV("file1.csv")
	if err != nil {
		panic(fmt.Errorf("Could not read file1 %v\n", err))
	}

	//-

	br1 := breakup("1", ch1)
	br2 := breakup("2", ch1)
	br3 := breakup("3", ch1)

	for {
		if br1 == nil && br2 == nil && br3 == nil {
			break
		}

		select {
		case _, ok := <-br1:
			if !ok {
				br1 = nil
			}
		case _, ok := <-br2:
			if !ok {
				br2 = nil
			}
		case _, ok := <-br3:
			if !ok {
				br3 = nil
			}
		}
	}

	fmt.Println("All completed, exiting")
}

func breakup(worker string, ch <-chan []string) chan struct{} {
	chE := make(chan struct{})

	go func() {
		for v := range ch {
			fmt.Println(worker, v)
		}

		close(chE)
	}()

	return chE
}

func merge(cs ...<-chan []string) <-chan []string {
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
			fmt.Println(chans)
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
