package clients

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

// type StockCandles struct {
// 	C  *[]float32 `json:"c,omitempty"` 	-> Close prices: [175.50, 175.60, 175.45] - An array of closing prices for each candlestick period, representing the final trade price of each period (e.g., each minute or day).
// 	H  *[]float32 `json:"h,omitempty"` 	-> High prices: [175.70, 175.80, 175.50] - An array of the highest prices reached during each candlestick period (e.g., the peak price in a minute or day).
// 	L  *[]float32 `json:"l,omitempty"` 	-> Low prices: [175.40, 175.55, 175.30] - An array of the lowest prices reached during each candlestick period (e.g., the bottom price in a minute or day).
// 	O  *[]float32 `json:"o,omitempty"` 	-> Open prices: [175.60, 175.50, 175.65] - An array of opening prices for each candlestick period, representing the first trade price of each period (e.g., start of a minute or day).
// 	V  *[]int64   `json:"v,omitempty"` 	-> Volume: [125000, 130000, 115000] - An array of trading volumes (number of shares traded) for each candlestick period.
// 	T  *[]int64   `json:"t,omitempty"` 	-> Timestamps: [1678886400, 1678886460, 1678886520] - An array of timestamps marking the end of each candlestick period, expressed as Unix timestamps (seconds since January 1, 1970 UTC).
// 	S  *string    `json:"s,omitempty"` 	-> Status: "ok" - Indicates the success of the API response, either "ok" (data returned) or "no_data" (no data available for the requested range or symbol).
// }

type StockCandles struct {
	ClosePrices []float32
	HighPrices  []float32
	LowPrices   []float32
	OpenPrices  []float32
	Volume      []float32
	Timestamps  []int64
	Status      string
}
