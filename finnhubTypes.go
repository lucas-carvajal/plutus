package main

// type Quote struct {
// 	C   *float32 `json:"c,omitempty"` 	-> Current price: 100.00
// 	D   *float32 `json:"d,omitempty"` 	-> The absolute change in price since the previous close, calculated as C - Pc (Current Price minus Previous Close)
// 	Dp  *float32 `json:"dp,omitempty"` 	-> The percentage change in price since the previous close, calculated as ((C - Pc) / Pc) * 100.
// 	H   *float32 `json:"h,omitempty"` 	-> The highest price the stock reached during the current trading session (usually the current trading day).
// 	L   *float32 `json:"l,omitempty"` 	-> The lowest price the stock reached during the current trading session (usually the current trading day).
// 	O   *float32 `json:"o,omitempty"` 	-> The price of the stock at the opening of the current trading session (the first trade of the day).
// 	Pc  *float32 `json:"pc,omitempty"`	-> The closing price of the stock from the previous trading session (typically the prior day's close).
// 	T   *int64   `json:"t,omitempty"` 	-> The time of the last trade or quote update, expressed as a Unix timestamp: 1715558400
// }

type StockQuote struct {
	CurrentPrice  float32
	TodaysHigh    float32
	TodaysLow     float32
	TodaysOpen    float32
	PreviousClose float32
}
