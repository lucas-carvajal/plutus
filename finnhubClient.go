package main

import (
	"context"
	"fmt"

	finnhub "github.com/Finnhub-Stock-API/finnhub-go/v2"
)

type FinnhubClient struct {
	client *finnhub.DefaultApiService
}

func NewFinnhubClient(apiKey string) *FinnhubClient {
	cfg := finnhub.NewConfiguration()
	cfg.AddDefaultHeader("X-Finnhub-Token", apiKey)
	client := finnhub.NewAPIClient(cfg).DefaultApi
	return &FinnhubClient{client: client}
}

func (s *FinnhubClient) GetStockPrice(symbol string) (float32, error) {
	quote, response, err := s.client.Quote(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		return 0, fmt.Errorf("error fetching stock quote: %v, response: %v", err, response)
	}
	if quote.C == nil {
		return 0, fmt.Errorf("no current price available for %s", symbol)
	}
	return *quote.C, nil
}

// GetStockQuote retrieves detailed quote data for a given stock symbol
func (s *FinnhubClient) GetStockQuote(symbol string) (*StockQuote, error) {
	quote, response, err := s.client.Quote(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		return nil, fmt.Errorf("error fetching stock quote: %v, response: %v", err, response)
	}
	if quote.C == nil || quote.H == nil || quote.L == nil || quote.O == nil || quote.Pc == nil {
		return nil, fmt.Errorf("incomplete quote data for %s", symbol)
	}
	return &StockQuote{
		CurrentPrice:  *quote.C,
		TodaysHigh:    *quote.H,
		TodaysLow:     *quote.L,
		TodaysOpen:    *quote.O,
		PreviousClose: *quote.Pc,
	}, nil
}
