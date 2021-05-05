package bookstore_test

import (
	"bookstore"
	"testing"
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
	}
}
