package library

import (
	"fmt"
	"strings"
)

//go:generate counterfeiter -o librarytesting/fake_book_store.gen.go . BookStore
//go:generate counterfeiter -o librarytesting/fake_book_details_store.gen.go . BookDetailsStore

type Book struct {
	GUID string

	ISBN string
	Details
}

type Details struct {
	Author string
	Name   string
}

type BookService struct {
	Book    BookStore
	Details BookDetailsStore
}

type BookStore interface {
	Create(isbn string, details Details) (Book, error)
}

type BookDetailsStore interface {
	Find(isbn string) (Details, error)
}

func (b *BookService) Create(isbn string) (Book, error) {
	details, err := b.Details.Find(isbn)
	if err != nil {
		return Book{}, fmt.Errorf("details store %v: %w", isbn, err)
	}

	book, err := b.Book.Create(strings.ToUpper(isbn), details)
	if err != nil {
		return Book{}, fmt.Errorf("book store %v: %w", isbn, err)
	}

	return book, nil
}
