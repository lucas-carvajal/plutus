package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"plutus/domain"
	"time"
)

// CONTINUE HERE: https://grok.com/chat/6514b4c4-6882-487b-88e4-9b7ff8692ed8

type TwelvedataClient struct {
	apiKey string
	client *http.Client
}

func NewTwelvedataClient(apiKey string) *TwelvedataClient {
	return &TwelvedataClient{
		apiKey: apiKey,
		client: &http.Client{Timeout: 10 * time.Second},
	}
}

func (c *TwelvedataClient) GetLatestPriceAndVolume(symbol string, interval string) (float64, int64, time.Time, error) {
	url := fmt.Sprintf(
		"https://api.twelvedata.com/time_series?symbol=%s&interval=%s&outputsize=1&apikey=%s",
		symbol, interval, c.apiKey,
	)

	req, err := http.NewRequestWithContext(context.Background(), "GET", url, nil)
	if err != nil {
		return 0, 0, time.Time{}, fmt.Errorf("request error: %v", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return 0, 0, time.Time{}, fmt.Errorf("fetch error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, 0, time.Time{}, fmt.Errorf("status: %s", resp.Status)
	}

	var ts TimeSeriesResponse
	if err := json.NewDecoder(resp.Body).Decode(&ts); err != nil {
		return 0, 0, time.Time{}, fmt.Errorf("decode error: %v", err)
	}

	if ts.Status != "ok" || len(ts.Values) == 0 {
		return 0, 0, time.Time{}, fmt.Errorf("no data: %+v", ts)
	}

	latest := ts.Values[0]
	closePrice, err := parseFloat(latest.Close)
	if err != nil {
		return 0, 0, time.Time{}, fmt.Errorf("close parse error: %v", err)
	}
	volume, err := parseInt(latest.Volume)
	if err != nil {
		return 0, 0, time.Time{}, fmt.Errorf("volume parse error: %v", err)
	}
	timestamp, err := time.Parse("2006-01-02 15:04:05", latest.Datetime)
	if err != nil {
		return 0, 0, time.Time{}, fmt.Errorf("time parse error: %v", err)
	}

	return closePrice, volume, timestamp, nil
}

func (c *TwelvedataClient) GetLatestQuote(symbol string) (domain.Quote, error) {
	url := fmt.Sprintf(
		"https://api.twelvedata.com/quote?symbol=%s&apikey=%s",
		symbol, c.apiKey,
	)

	req, err := http.NewRequestWithContext(context.Background(), "GET", url, nil)
	if err != nil {
		return domain.Quote{}, fmt.Errorf("request error: %v", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return domain.Quote{}, fmt.Errorf("fetch error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body := make([]byte, 1024)
		resp.Body.Read(body)
		return domain.Quote{}, fmt.Errorf("status: %s, body: %s", resp.Status, string(body))
	}

	var quoteResponse QuoteResponse
	if err := json.NewDecoder(resp.Body).Decode(&quoteResponse); err != nil {
		return domain.Quote{}, fmt.Errorf("decode error: %v", err)
	}

	if quoteResponse.Symbol == "" {
		return domain.Quote{}, fmt.Errorf("no valid data returned: %+v", quoteResponse)
	}

	quote, err := quoteResponse.ToQuote()
	if err != nil {
		return domain.Quote{}, fmt.Errorf("conversion error: %v", err)
	}

	return quote, nil
}

func parseFloat(s string) (float64, error) {
	var f float64
	_, err := fmt.Sscanf(s, "%f", &f)
	return f, err
}

func parseInt(s string) (int64, error) {
	var i int64
	_, err := fmt.Sscanf(s, "%d", &i)
	return i, err
}
