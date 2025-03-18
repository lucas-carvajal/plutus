package service

import (
	"plutus/clients"
	"plutus/repository"
)

type DataIngestionService struct {
	repo       *repository.QuoteRepository
	twelveData *clients.TwelvedataClient
}

func NewDataIngestionService(repo *repository.QuoteRepository, twelveData *clients.TwelvedataClient) *DataIngestionService {
	return &DataIngestionService{repo: repo, twelveData: twelveData}
}

func (s *DataIngestionService) IngestNewData(symbol string) error {
	quote, err := s.twelveData.GetLatestQuote(symbol)
	if err != nil {
		return err
	}

	_, err = s.repo.SaveOrUpdate(quote)
	return err
}
