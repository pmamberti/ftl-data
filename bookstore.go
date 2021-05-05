package bookstore

import "fmt"

// Book represents information about a Book
type Book struct {
	Title           string
	Author          []string
	Description     string
	Series          int
	SeriesTitle     string
	PriceCents      int
	DiscountPercent int
	Copies          int
	NetPrice        int
	Featured        bool
}

var books []Book

func netPrice(b Book) int {
	discount := float64(
		b.PriceCents,
	) * float64(b.DiscountPercent) / 100
	return b.PriceCents - int(discount)
}

func AddBook(c []Book, b Book) []Book {
	c = append(c, b)
	return c
}

func FeaturedBooks(c []Book) []Book {
	var featured []Book
	for _, b := range c {
		if b.Featured {
			featured = append(featured, b)
		}
	}
	return featured
}

func main() {
	books = []Book{}

	b := Book{
		Title:           "An Elegant Puzzle",
		Author:          []string{"Will Larson"},
		PriceCents:      2999,
		DiscountPercent: 10,
	}

	b1 := Book{
		Title:           "Staff Engineer",
		Author:          []string{"Will Larson"},
		PriceCents:      2499,
		DiscountPercent: 10,
		Featured:        true,
	}

	b.NetPrice = netPrice(b)
	b1.NetPrice = netPrice(b1)

	books = AddBook(books, b)
	books = AddBook(books, b1)

	for _, b := range books {
		fmt.Printf(
			"Author(s): %v \nTitle: %v \n ============= \n",
			b.Author,
			b.Title,
		)
	}

	fmt.Printf("Featured: %v", FeaturedBooks(books))
}
