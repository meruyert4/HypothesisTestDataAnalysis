package main

import (
	"fmt"
	"log"
	"path/filepath"

	"hypothesistestanalysis/internal"
)

func main() {
	filePath := filepath.Join("bestsellers with categories.csv")

	books, err := internal.LoadBooksFromCSV(filePath)
	if err != nil {
		log.Fatalf("Failed to load books: %v", err)
	}

	z, meanFiction, meanNonFiction, nF, nNF := internal.CompareFictionVsNonFictionPrices(books)

	fmt.Printf("📚 Fiction Avg Price: $%.2f (%d books)\n", meanFiction, nF)
	fmt.Printf("📘 Non Fiction Avg Price: $%.2f (%d books)\n", meanNonFiction, nNF)
	fmt.Println(internal.InterpretPriceZScore(z))
}
