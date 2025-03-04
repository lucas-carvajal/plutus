package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
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
