package bookstore

import "fmt"

// Book represents information about a Book
type Book struct {
	ID              string
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

var Books = []Book{}

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

func GetAllBooks(c []Book) []Book {
	return c
}

func main() {
	fmt.Println("Hello! This is Main")
}
