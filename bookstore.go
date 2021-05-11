package bookstore

import (
	"fmt"
	"math/rand"
	"time"
)

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

type Catalog map[string]Book

var Books = Catalog{
	"Book1": {
		ID:     "Book1",
		Title:  "This is my first Book",
		Author: []string{"The Author"},
	},
	"Book2": {
		ID:     "Book2",
		Title:  "This is my second Book",
		Author: []string{"Another Author"},
	},
	"Book3": {
		ID:     "Book3",
		Title:  "This is our third Book",
		Author: []string{"Another Author", "The Author"},
	},
}

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

func GetAllBooks(c Catalog) Catalog {
	return c
}

func NewID() (id string) {
	rand.Seed(time.Now().UnixNano())
	charset := "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "0123456789"
	for len(id) < 3 {
		id += string(charset[rand.Intn(len(charset))])
	}
	return id
}

func GetBookDetails(id string, c Catalog) (Book, bool) {
	b, ok := c[id]
	return b, ok
}

func isIn(s string, l []string) bool {
	for _, i := range l {
		if s == i {
			return true
		}
	}
	return false
}

func GetAllByAuthor(a string) []Book {
	results := []Book{}
	for _, b := range Books {
		if isIn(a, b.Author) {
			results = append(results, b)
		}
	}
	return results
}

func main() {
	fmt.Println("Hello! This is Main")
}
