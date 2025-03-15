package main

import (
	"fmt"
	"plutus/domain"
	"strconv"
	"time"
)

type QuoteResponse struct {
	Symbol        string `json:"symbol"`
	Name          string `json:"name"`
	Exchange      string `json:"exchange"`
	MicCode       string `json:"mic_code"`
	Currency      string `json:"currency"`
	Datetime      string `json:"datetime"`
	Timestamp     int64  `json:"timestamp"`
	Open          string `json:"open"`
	High          string `json:"high"`
	Low           string `json:"low"`
	Close         string `json:"close"`
	Volume        string `json:"volume"`
	PreviousClose string `json:"previous_close"`
	Change        string `json:"change"`
	PercentChange string `json:"percent_change"`
	AverageVolume string `json:"average_volume"`
	IsMarketOpen  bool   `json:"is_market_open"`
	FiftyTwoWeek  struct {
		High string `json:"high"`
		Low  string `json:"low"`
	} `json:"fifty_two_week"`
}

func (qr *QuoteResponse) ToQuote() (domain.Quote, error) {
	open, err := strconv.ParseFloat(qr.Open, 64)
	if err != nil {
		return domain.Quote{}, fmt.Errorf("open parse error: %v", err)
	}
	high, err := strconv.ParseFloat(qr.High, 64)
	if err != nil {
		return domain.Quote{}, fmt.Errorf("high parse error: %v", err)
	}
	low, err := strconv.ParseFloat(qr.Low, 64)
	if err != nil {
		return domain.Quote{}, fmt.Errorf("low parse error: %v", err)
	}
	closePrice, err := strconv.ParseFloat(qr.Close, 64)
	if err != nil {
		return domain.Quote{}, fmt.Errorf("close parse error: %v", err)
	}
	volume, err := strconv.ParseInt(qr.Volume, 10, 64) // Parse string to int64
	if err != nil {
		return domain.Quote{}, fmt.Errorf("volume parse error: %v", err)
	}
	previousClose, err := strconv.ParseFloat(qr.PreviousClose, 64)
	if err != nil {
		return domain.Quote{}, fmt.Errorf("previous_close parse error: %v", err)
	}
	change, err := strconv.ParseFloat(qr.Change, 64)
	if err != nil {
		return domain.Quote{}, fmt.Errorf("change parse error: %v", err)
	}
	percentChange, err := strconv.ParseFloat(qr.PercentChange, 64)
	if err != nil {
		return domain.Quote{}, fmt.Errorf("percent_change parse error: %v", err)
	}
	averageVolume, err := strconv.ParseInt(qr.AverageVolume, 10, 64) // Parse string to int64
	if err != nil {
		return domain.Quote{}, fmt.Errorf("average_volume parse error: %v", err)
	}
	fiftyTwoWeekHigh, err := strconv.ParseFloat(qr.FiftyTwoWeek.High, 64)
	if err != nil {
		return domain.Quote{}, fmt.Errorf("fifty_two_week.high parse error: %v", err)
	}
	fiftyTwoWeekLow, err := strconv.ParseFloat(qr.FiftyTwoWeek.Low, 64)
	if err != nil {
		return domain.Quote{}, fmt.Errorf("fifty_two_week.low parse error: %v", err)
	}

	datetime := time.Unix(qr.Timestamp, 0)

	return domain.Quote{
		Symbol:        qr.Symbol,
		Name:          qr.Name,
		Exchange:      qr.Exchange,
		MicCode:       qr.MicCode,
		Currency:      qr.Currency,
		Datetime:      datetime,
		Timestamp:     qr.Timestamp,
		Open:          open,
		High:          high,
		Low:           low,
		Close:         closePrice,
		Volume:        volume, // Now int64 from parsed string
		PreviousClose: previousClose,
		Change:        change,
		PercentChange: percentChange,
		AverageVolume: averageVolume, // Now int64 from parsed string
		IsMarketOpen:  qr.IsMarketOpen,
		FiftyTwoWeek: struct {
			High float64
			Low  float64
		}{
			High: fiftyTwoWeekHigh,
			Low:  fiftyTwoWeekLow,
		},
	}, nil
}

type TimeSeriesResponse struct {
	Meta struct {
		Symbol   string `json:"symbol"`
		Interval string `json:"interval"`
	} `json:"meta"`
	Values []struct {
		Datetime string `json:"datetime"`
		Open     string `json:"open"`
		High     string `json:"high"`
		Low      string `json:"low"`
		Close    string `json:"close"`
		Volume   string `json:"volume"`
	} `json:"values"`
	Status string `json:"status"`
}
