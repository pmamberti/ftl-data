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
	testCatalog := bookstore.Catalog{
		"book1": bookstore.Book{
			ID:     "Book1",
			Title:  "This is my first Book",
			Author: []string{"The Author"},
		},
		"book2": bookstore.Book{
			ID:     "Book2",
			Title:  "This is my second Book",
			Author: []string{"Another Author"},
		},
	}
	want := testCatalog
	got := bookstore.GetAllBooks(testCatalog)

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
	testCatalog := bookstore.Catalog{
		"Book1": {
			ID:     "Book1",
			Title:  "This is my First Book",
			Author: []string{"The Author"},
		},
	}
	testCases := []struct {
		name        string
		catalog     bookstore.Catalog
		id          string
		want        bookstore.Book
		errExpected bool
	}{
		{
			name:        "id exists in catalog, returns book by id",
			catalog:     testCatalog,
			id:          "Book1",
			want:        testCatalog["Book1"],
			errExpected: false,
		},
		{
			name:        "id does not exists in catalog, returns false",
			catalog:     testCatalog,
			id:          "Book2",
			want:        testCatalog["Book2"],
			errExpected: true,
		},
	}

	for _, tc := range testCases {
		got, err := bookstore.GetBookDetails(tc.id, tc.catalog)
		errReceived := err != true

		if errReceived != tc.errExpected {
			t.Errorf(
				"unexpected error \n \t(%v) expected \n\t(%v) received",
				tc.errExpected,
				errReceived,
			)
		}

		if !cmp.Equal(tc.want, got) {
			t.Errorf(cmp.Diff(tc.want, got))
		}
	}
}

// func TestGetAllByAuthor(t *testing.T) {
// 	want := []bookstore.Book{
// 		{
// 			ID:     "Book1",
// 			Title:  "This is my first Book",
// 			Author: []string{"The Author"},
// 		},
// 		{
// 			ID:     "Book3",
// 			Title:  "This is our third Book",
// 			Author: []string{"Another Author", "The Author"},
// 		},
// 	}
// 	got := bookstore.GetAllByAuthor("The Author")

// 	if !cmp.Equal(want, got) {
// 		t.Errorf(cmp.Diff(want, got))
// 	}
// }
