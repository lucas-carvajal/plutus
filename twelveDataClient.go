package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type TwelveDataClient struct {
	apiKey string
	client *http.Client
}

func NewTwelveDataClient(apiKey string) *TwelveDataClient {
	return &TwelveDataClient{
		apiKey: apiKey,
		client: &http.Client{Timeout: 10 * time.Second},
	}
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

func (c *TwelveDataClient) GetLatestPriceAndVolume(symbol string, interval string) (float64, int64, time.Time, error) {
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
