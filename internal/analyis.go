package internal

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Book struct {
	Title       string
	Rating      float64
	Reviews     int
	Price       float64
	TitleLength int
	Genre       string
}

func LoadBooksFromCSV(filePath string) ([]Book, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var books []Book

	for i, row := range records {
		if i == 0 {
			continue // skip header
		}

		userRating, _ := strconv.ParseFloat(row[2], 64)
		reviews, _ := strconv.Atoi(row[3])
		price, _ := strconv.ParseFloat(row[4], 64)
		genre := row[6]

		books = append(books, Book{
			Title:       row[0],
			Rating:      userRating,
			Reviews:     reviews,
			Price:       price,
			TitleLength: len(strings.Fields(row[0])),
			Genre:       genre,
		})
	}

	return books, nil
}

// Calculates mean prices and z-score to compare Fiction vs Non Fiction
func CompareFictionVsNonFictionPrices(books []Book) (z, meanFiction, meanNonFiction float64, nFiction, nNonFiction int) {
	var fictionPrices, nonFictionPrices []float64

	for _, b := range books {
		if b.Genre == "Fiction" {
			fictionPrices = append(fictionPrices, b.Price)
		} else {
			nonFictionPrices = append(nonFictionPrices, b.Price)
		}
	}

	nFiction = len(fictionPrices)
	nNonFiction = len(nonFictionPrices)

	meanFiction = mean(fictionPrices)
	meanNonFiction = mean(nonFictionPrices)
	stdFiction := stdDev(fictionPrices, meanFiction)
	stdNonFiction := stdDev(nonFictionPrices, meanNonFiction)

	// Standard error
	se := math.Sqrt((stdFiction*stdFiction)/float64(nFiction) + (stdNonFiction*stdNonFiction)/float64(nNonFiction))

	// z = (mean1 - mean2) / SE
	z = (meanFiction - meanNonFiction) / se

	return z, meanFiction, meanNonFiction, nFiction, nNonFiction
}

func InterpretPriceZScore(z float64) string {
	// One-tailed test at 95% confidence → z < -1.645 to be significant
	if z < -1.645 {
		return fmt.Sprintf("Z-score: %.2f → Reject H₀. Fiction books are significantly cheaper. ✅", z)
	}
	return fmt.Sprintf("Z-score: %.2f → Fail to reject H₀. No strong evidence Fiction books are cheaper. ❌", z)
}
