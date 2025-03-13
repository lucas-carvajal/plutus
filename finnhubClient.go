package main

import (
	"context"
	"fmt"
	"time"

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

func (client *FinnhubClient) GetLatestVolume(symbol string) (StockCandles, error) {
	now := time.Now().Unix()
	from := now - 10*60 // 10 minutes ago
	resolution := "1"   // 1-minute candles

	candles, response, err := client.client.StockCandles(context.Background()).
		Symbol(symbol).
		Resolution(resolution).
		From(from).
		To(now).
		Execute()
	if err != nil {
		return StockCandles{}, fmt.Errorf("error fetching candles: %v, response: %v", err, response)
	}

	if candles.S == nil || *candles.S != "ok" || len(*candles.V) == 0 {
		return StockCandles{}, fmt.Errorf("no volume data available for %s", symbol)
	}

	return StockCandles{
		ClosePrices: *candles.C,
		HighPrices:  *candles.H,
		LowPrices:   *candles.L,
		OpenPrices:  *candles.O,
		Volume:      *candles.V,
		Timestamps:  *candles.T,
		Status:      *candles.S,
	}, nil
}
