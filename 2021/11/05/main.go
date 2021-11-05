package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	ServerDefaultReadTimeout  time.Duration = 100 * time.Millisecond
	ServerDefaultWriteTimeout time.Duration = 100 * time.Millisecond
	ServerDefaultAddress      string        = ":8080"
)

type Option func(*http.Server) error

func WithAddr(addr string) func(*http.Server) error {
	return func(s *http.Server) error {
		s.Addr = addr
		return nil
	}
}

func WithReadTimeout(t time.Duration) func(*http.Server) error {
	return func(s *http.Server) error {
		if t > time.Second {
			return errors.New("timeout value not allowed")
		}

		s.ReadTimeout = t
		return nil
	}
}

func WithWriteTimeout(t time.Duration) func(*http.Server) error {
	return func(s *http.Server) error {
		if t > time.Second {
			return errors.New("timeout value not allowed")
		}

		s.WriteTimeout = t

		return nil
	}
}

func NewServer(opts ...Option) (http.Server, error) {
	s := http.Server{
		Addr:         ServerDefaultAddress,
		ReadTimeout:  ServerDefaultReadTimeout,
		WriteTimeout: ServerDefaultWriteTimeout,
	}

	for _, opt := range opts {
		if err := opt(&s); err != nil {
			return http.Server{}, fmt.Errorf("option failed %w", err)
		}
	}

	return s, nil
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello")
	})

	s, err := NewServer(
		WithAddr(":9191"),
		WithReadTimeout(100*time.Millisecond),
		WithWriteTimeout(100*time.Millisecond),
	)

	if err != nil {
		log.Fatalln("Couldn't initialize server", err)
	}

	s.Handler = mux

	if err := s.ListenAndServe(); err != nil {
		log.Fatalln("Couldn't listen and serve", err)
	}
}
