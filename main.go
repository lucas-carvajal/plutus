package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	twelveData := NewTwelveDataClient(TWELVE_DATA_API_KEY)

	price, volume, timestamp, err := twelveData.GetLatestPriceAndVolume("MSTR", "5min")
	if err != nil {
		log.Fatalf("Failed: %v", err)
	}
	fmt.Printf("MSTR: Price %.2f, Volume %d, Time %s\n", price, volume, timestamp.Format(time.RFC3339))
	fmt.Println("=== *** === *** === *** === ***")

	finnhub := NewFinnhubClient(FINNHUB_API_KEY)

	// Example: Get the full quote
	symbol := "AAPL"
	quote, err := finnhub.GetStockQuote(symbol)
	if err != nil {
		log.Fatalf("Failed to get stock quote: %v", err)
	}

	fmt.Printf("Quote for %s: Current: %.2f, High: %.2f, Low: %.2f, Open: %.2f, Previous Close: %.2f\n",
		symbol, quote.CurrentPrice, quote.TodaysHigh, quote.TodaysLow, quote.TodaysOpen, quote.PreviousClose)

	candles, err := finnhub.GetLatestVolume(symbol)
	if err != nil {
		log.Fatalf("Failed to get stock candles: %v", err)
	}

	fmt.Printf("Comparing latest quotes for %s from quotes and candles:\n", symbol)
	fmt.Printf("Quote's price: %.2f\n", quote.CurrentPrice)
	fmt.Printf("Candles's high and low price: %.2f - %.2f\n", candles.ClosePrices[0], candles.ClosePrices[len(candles.ClosePrices)-1])

	fmt.Println("=== === ===")

	fmt.Println("Candles data from the last 10 minutes:")
	for i, _ := range candles.ClosePrices {
		timeStr := time.Unix(candles.Timestamps[i], 0).Format("2006-01-02 15:04:05")
		fmt.Printf("Time: %s\n", timeStr)
		fmt.Printf("High Price: %.2f\n", candles.HighPrices[i])
		fmt.Printf("Low Price: %.2f\n", candles.LowPrices[i])
		fmt.Printf("Volume: %.2f\n", candles.Volume[i])
	}

	// type StockCandles struct {
	// 	ClosePrices []float32
	// 	HighPrices  []float32
	// 	LowPrices   []float32
	// 	OpenPrices  []float32
	// 	Volume      []float32
	// 	Timestamps  []int64
	// 	Status      string
	// }
	// ===

	// Start the logging goroutine
	go startLogging()

	setUpApi()
}

// startLogging runs a ticker that logs "hello" every 10 minutes
func startLogging() {
	ticker := time.NewTicker(10 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		log.Println("hello")
	}
}

func setUpApi() {
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello from the server",
		})
	})

	// Run the server
	if err := r.Run(":8080"); err != nil {
		log.Printf("Server failed to start: %v", err)
	}
}
