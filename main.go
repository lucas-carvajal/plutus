package main

import (
	"fmt"
	"log"

	"plutus/clients"
	"plutus/repository"
	"plutus/service"

	"github.com/gin-gonic/gin"
)

func main() {
	repo, err := repository.NewQuoteRepository(POSTGRES_CONNECTION_STRING)
	if err != nil {
		fmt.Printf("Repository init error: %v\n", err)
		return
	}

	twelveData := clients.NewTwelvedataClient(TWELVE_DATA_API_KEY)

	dataIngestionService := service.NewDataIngestionService(repo, twelveData)

	tickerService := service.NewTickerService(dataIngestionService)

	go tickerService.Start()

	latestQuote, err := twelveData.GetLatestQuote("MSTR")
	if err != nil {
		log.Fatalf("Failed: %v", err)
	}
	fmt.Println(latestQuote.OutputFormatted())

	setUpApi()
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
