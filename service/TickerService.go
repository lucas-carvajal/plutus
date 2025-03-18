package service

import (
	"fmt"
	"log"
	"time"
)

type TickerService struct {
	dataIngestionService *DataIngestionService
}

func NewTickerService(dataIngestionService *DataIngestionService) *TickerService {
	return &TickerService{dataIngestionService: dataIngestionService}
}

func (s *TickerService) Start() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		currentMinute := time.Now().Minute()

		fmt.Println("currentMinute", currentMinute)

		if currentMinute%10 == 0 {
			log.Println("Running scheduled data ingestion")
			err := s.dataIngestionService.IngestNewData("MSTR")
			if err != nil {
				log.Printf("Data ingestion failed: %v", err)
			}
		}
	}
}
