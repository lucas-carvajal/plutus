package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	twelveData := NewTwelvedataClient(TWELVE_DATA_API_KEY)

	// price, volume, timestamp, err := twelveData.GetLatestPriceAndVolume("MSTR", "5min")
	// if err != nil {
	// 	log.Fatalf("Failed: %v", err)
	// }
	// fmt.Printf("MSTR: Price %.2f, Volume %d, Time %s\n", price, volume, timestamp.Format(time.RFC3339))
	// fmt.Println("=== *** === *** === *** === ***")

	latestQuote, err := twelveData.GetLatestQuote("MSTR")
	if err != nil {
		log.Fatalf("Failed: %v", err)
	}
	fmt.Println(latestQuote.OutputFormatted())

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
