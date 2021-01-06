package library_test

import (
	"errors"
	"testing"

	library "github.com/MarioCarrion/videos/counterfeiter"
	"github.com/MarioCarrion/videos/counterfeiter/librarytesting"

	"github.com/google/go-cmp/cmp"
)

func TestBook_Create(t *testing.T) {
	type output struct {
		book    library.Book
		withErr bool
	}

	type spyCreate struct {
		isbn    string
		details library.Details
	}

	tests := []struct {
		name   string
		setup  func(*librarytesting.FakeBookStore, *librarytesting.FakeBookDetailsStore)
		input  string
		output output
		args   spyCreate
	}{
		{
			"Dummy=BookDetailsStore, Stub=BookStore, Fake=both!",
			func(b *librarytesting.FakeBookStore, d *librarytesting.FakeBookDetailsStore) {
				b.CreateReturns(library.Book{
					GUID: "guid",
					ISBN: "isbn",
					Details: library.Details{
						Author: "author",
						Name:   "name",
					},
				}, nil)
			},
			"isbn",
			output{
				book: library.Book{
					GUID: "guid",
					ISBN: "isbn",
					Details: library.Details{
						Author: "author",
						Name:   "name",
					},
				},
			},
			spyCreate{
				isbn: "ISBN",
			},
		},
		{
			"Mock=BookDetailsStore",
			func(b *librarytesting.FakeBookStore, d *librarytesting.FakeBookDetailsStore) {
				d.FindReturns(library.Details{}, errors.New("details store"))
			},
			"isbn",
			output{
				withErr: true,
			},
			spyCreate{},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			book := &librarytesting.FakeBookStore{}
			details := &librarytesting.FakeBookDetailsStore{}

			test.setup(book, details)

			svc := library.BookService{Book: book, Details: details}
			actual, err := svc.Create(test.input)

			if diff := cmp.Diff(test.output.book, actual); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}

			if test.output.withErr != (err != nil) {
				t.Errorf("expected err %T, got %s", test.output.withErr, err)
			}

			if test.args.isbn != "" { // Spy test double
				argsISBN, argsDetails := book.CreateArgsForCall(0)

				if diff := cmp.Diff(test.args.isbn, argsISBN); diff != "" {
					t.Errorf("mismatch (-want +got):\n%s", diff)
				}

				if diff := cmp.Diff(test.args.details, argsDetails); diff != "" {
					t.Errorf("mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
