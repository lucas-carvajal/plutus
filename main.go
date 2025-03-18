package main

import (
	"fmt"
	"log"
	"time"

	"plutus/clients"
	"plutus/repository"
	"plutus/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// PostgreSQL connection string
	// TODO provide real one
	connStr := "postgres://postgres:yourpassword@localhost:5432/quotes_db?sslmode=disable"

	// Initialize repository
	repo, err := repository.NewQuoteRepository(connStr)
	if err != nil {
		fmt.Printf("Repository init error: %v\n", err)
		return
	}

	fmt.Printf("Repository: %v\n", repo)

	twelveData := clients.NewTwelvedataClient(TWELVE_DATA_API_KEY)

	dataIngestionService := service.NewDataIngestionService(repo, twelveData)

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

	go startLogging(dataIngestionService)

	setUpApi()
}

func startLogging(dataIngestionService *service.DataIngestionService) {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		currentMinute := time.Now().Minute()

		if currentMinute%10 == 0 {
			log.Println("Running scheduled data ingestion")
			err := dataIngestionService.IngestNewData("MSTR")
			if err != nil {
				log.Printf("Data ingestion failed: %v", err)
			}
		} else {
			log.Printf("Skipping data ingestion - current minute (%d) is not divisible by 10", currentMinute)
		}
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
