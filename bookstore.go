package main

import "fmt"

// Book represents information about a Book
type Book struct {
	Title           string
	Author          string
	Series          int
	SeriesTitle     string
	PriceCents      int
	DiscountPercent int
	Copies          int
	NetPrice        int
}

func netPrice(b Book) int {
	discount := float64(b.PriceCents) * float64(b.DiscountPercent) / 100
	return b.PriceCents - int(discount)
}

func main() {
	b := Book{
		Title:           "An Elegant Puzzle",
		Author:          "Will Larson",
		PriceCents:      2999,
		DiscountPercent: 10,
	}

	b.NetPrice = netPrice(b)

	fmt.Printf(
		`Author: %v
Title: %v
Price: %v
Discount: %v%%`, b.Author, b.Title, b.PriceCents/100, b.DiscountPercent)
}
