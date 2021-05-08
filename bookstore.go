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

var Books = []Book{
	{
		ID:     "Book1",
		Title:  "This is my first Book",
		Author: []string{"The Author"},
	},
	{
		ID:     "Book2",
		Title:  "This is my second Book",
		Author: []string{"Another Author"},
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

func GetAllBooks(c []Book) []Book {
	return c
}

func NewID() (id string) {
	rand.Seed(time.Now().UnixNano())
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	for len(id) < 3 {
		id += string(charset[rand.Intn(len(charset))])
	}

	return id
}

func GetBookDetails(id string) string {
	for _, b := range Books {
		if b.ID == id {
			return fmt.Sprintf(
				"ID: %v, Title: %v, Author: %v",
				b.ID,
				b.Title,
				b.Author[0],
			)
		}
	}

	return "Book not found"
}

func main() {
	fmt.Println("Hello! This is Main")
}
