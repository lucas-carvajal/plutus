package marketHours

import "time"

type MarketHours struct {
	OpenTime  time.Duration // Hours since midnight in local time
	CloseTime time.Duration // Hours since midnight in local time
	TimeZone  *time.Location
	Days      map[time.Weekday]bool // Which days the market is open
}

type Market string

const (
	NASDAQ Market = "NASDAQ"
)

var MarketHoursMap = map[Market]MarketHours{
	NASDAQ: {
		OpenTime:  9*time.Hour + 30*time.Minute,  // 9:30 AM
		CloseTime: 16 * time.Hour,                // 4:00 PM
		TimeZone:  time.FixedZone("ET", -5*3600), // Eastern Time (-5 UTC)
		Days: map[time.Weekday]bool{
			time.Monday:    true,
			time.Tuesday:   true,
			time.Wednesday: true,
			time.Thursday:  true,
			time.Friday:    true,
			time.Saturday:  false,
			time.Sunday:    false,
		},
	},
}
