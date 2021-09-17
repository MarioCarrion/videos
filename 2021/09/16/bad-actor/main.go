package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type slowReader struct {
	Value string
	index int
}

func (s *slowReader) Read(p []byte) (n int, err error) {
	if s.index == len(s.Value) {
		fmt.Println("")

		return 0, io.EOF
	}

	p[0] = byte(s.Value[s.index])

	fmt.Printf("%s", string(s.Value[s.index]))

	time.Sleep(500 * time.Millisecond)

	s.index++

	return 1, nil
}

func main() {
	reader := slowReader{Value: "Wario"}

	req, err := http.NewRequestWithContext(context.Background(),
		http.MethodPost,
		"http://localhost:8080/hello",
		&reader)
	if err != nil {
		log.Fatalln(err)
	}

	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	io.Copy(os.Stdout, res.Body)

	fmt.Println("\nexiting")
}
