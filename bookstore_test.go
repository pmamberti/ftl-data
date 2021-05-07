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

func TestNewID(t *testing.T) {
	// TODO: Add test, once I know what it looks like
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
