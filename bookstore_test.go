package bookstore_test

import (
	"bookstore"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBook(t *testing.T) {
	_ = bookstore.Book{
		Title:           "Devops For Dummies",
		Author:          []string{"Emily Freeman"},
		Description:     "A gentle intro for everyone to the DevOps philosophy.",
		SeriesTitle:     "For Dummies",
		Series:          1,
		PriceCents:      1999,
		DiscountPercent: 15,
		Copies:          10,
		Featured:        true,
		ID:              "0000",
	}
}

func TestGetAllBooks(t *testing.T) {
	t.Parallel()
	book1 := bookstore.Book{
		ID:     "Book1",
		Title:  "This is my first Book",
		Author: []string{"The Author"},
	}
	book2 := bookstore.Book{
		ID:     "Book2",
		Title:  "This is my second Book",
		Author: []string{"Another Author"},
	}

	bookstore.Books = []bookstore.Book{book1, book2}
	want := bookstore.Books
	got := bookstore.GetAllBooks([]bookstore.Book{book1, book2})

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func isIn(s string, l []string) bool {
	for _, i := range l {
		if s == i {
			return true
		}
	}
	return false
}

func TestNewID(t *testing.T) {
	idSet := []string{}

	for len(idSet) < 1000 {
		idSet = append(idSet, bookstore.NewID())
	}
	got := bookstore.NewID()

	if isIn(got, idSet) {
		t.Errorf("ID %v is not unique", got)
	}
}

func TestGetBookDetails(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		id, want string
	}{
		{
			id:   "Book1",
			want: "ID: Book1, Title: This is my first Book, Author: The Author",
		},
		{id: "BookX", want: "Book not found"},
	}

	for _, tc := range testCases {
		got := bookstore.GetBookDetails(tc.id)
		if tc.want != got {
			t.Errorf("want (%v) got (%v)", tc.want, got)
		}

	}
}
