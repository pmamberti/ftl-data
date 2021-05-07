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
	book1 := bookstore.Book{Title: "This is Book 1"}
	book2 := bookstore.Book{Title: "This is Book 2"}
	bookstore.Books = []bookstore.Book{book1, book2}
	want := bookstore.Books
	got := bookstore.GetAllBooks([]bookstore.Book{book1, book2})

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestNewID(t *testing.T) {
}
