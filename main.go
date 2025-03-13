package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	stockService := NewFinnhubClient(FINNHUB_API_KEY)

	// === Example: Get just the current price
	symbol := "AAPL"
	price, err := stockService.GetStockPrice(symbol)
	if err != nil {
		log.Fatalf("Failed to get stock price: %v", err)
	}
	fmt.Printf("Current price for %s: %.2f\n", symbol, price)

	// Example: Get the full quote
	quote, err := stockService.GetStockQuote(symbol)
	if err != nil {
		log.Fatalf("Failed to get stock quote: %v", err)
	}
	fmt.Printf("Quote for %s: Current: %.2f, High: %.2f, Low: %.2f, Open: %.2f, Previous Close: %.2f\n",
		symbol, quote.CurrentPrice, quote.TodaysHigh, quote.TodaysLow, quote.TodaysOpen, quote.PreviousClose)
	// ===

	// Create a new Gin router
	r := gin.Default()

	// Set up API endpoints
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello from the server",
		})
	})

	// Start the logging goroutine
	go startLogging()

	// Run the server
	if err := r.Run(":8080"); err != nil {
		log.Printf("Server failed to start: %v", err)
	}
}

// startLogging runs a ticker that logs "hello" every 10 minutes
func startLogging() {
	ticker := time.NewTicker(10 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			log.Println("hello")
		}
	}
}
