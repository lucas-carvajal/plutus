package domain

import (
	"fmt"
	"time"
)

// Quote represents stock quote data from the Twelvedata /quote endpoint.
// Fields correspond directly to the API response structure.
type Quote struct {
	Symbol        string    // Stock ticker symbol (e.g., AAPL, MSFT)
	Name          string    // Full company name
	Exchange      string    // Exchange where the stock is traded (e.g., NASDAQ)
	MicCode       string    // Market Identifier Code for the exchange
	Currency      string    // Currency in which the stock is traded (e.g., USD)
	Datetime      time.Time // Date and time of the quote
	Timestamp     int64     // Unix timestamp of the quote
	Open          float64   // Opening price for the current/most recent trading day
	High          float64   // Highest price for the current/most recent trading day
	Low           float64   // Lowest price for the current/most recent trading day
	Close         float64   // Closing/current price
	Volume        int64     // Trading volume for the current/most recent trading day
	PreviousClose float64   // Previous day's closing price
	Change        float64   // Absolute change from previous close
	PercentChange float64   // Percentage change from previous close
	AverageVolume int64     // Average daily trading volume
	IsMarketOpen  bool      // Whether the market is currently open for trading
	FiftyTwoWeek  struct {
		High float64 // Highest price in the last 52 weeks
		Low  float64 // Lowest price in the last 52 weeks
	}
}

func (q *Quote) MidQuote() float64 {
	return (q.Open + q.Close) / 2
}

func (q *Quote) OutputFormatted() string {
	// No need to format the timestamp separately since we now have Datetime as time.Time
	dateTimeFormatted := q.Datetime.Format("2006-01-02 15:04:05")

	return fmt.Sprintf(
		"Quote Information for %s (%s):\n"+
			"----------------------------------------\n"+
			"Company Name:       %s\n"+
			"Exchange:           %s (%s)\n"+
			"Currency:           %s\n"+
			"Date/Time:          %s\n"+
			"Timestamp:          %d\n"+
			"----------------------------------------\n"+
			"Price Information:\n"+
			"Open:               %.2f %s\n"+
			"High:               %.2f %s\n"+
			"Low:                %.2f %s\n"+
			"Close:              %.2f %s\n"+
			"Previous Close:     %.2f %s\n"+
			"Change:             %.2f (%.2f%%)\n"+
			"Mid Quote:          %.2f %s\n"+
			"----------------------------------------\n"+
			"Volume Information:\n"+
			"Volume:             %d\n"+
			"Average Volume:     %d\n"+
			"----------------------------------------\n"+
			"52-Week Range:      %.2f - %.2f %s\n"+
			"Market Status:      %s\n",
		q.Symbol, q.Name,
		q.Name,
		q.Exchange, q.MicCode,
		q.Currency,
		dateTimeFormatted,
		q.Timestamp,
		q.Open, q.Currency,
		q.High, q.Currency,
		q.Low, q.Currency,
		q.Close, q.Currency,
		q.PreviousClose, q.Currency,
		q.Change, q.PercentChange,
		q.MidQuote(), q.Currency,
		q.Volume,
		q.AverageVolume,
		q.FiftyTwoWeek.Low, q.FiftyTwoWeek.High, q.Currency,
		marketStatus(q.IsMarketOpen),
	)
}

// Helper function to return market status as a string
func marketStatus(isOpen bool) string {
	if isOpen {
		return "Market Open"
	}
	return "Market Closed"
}
