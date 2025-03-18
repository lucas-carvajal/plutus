package marketHours

import "time"

func IsActiveNow(market Market) bool {
	mh := MarketHoursMap[market]

	now := time.Now()

	nowAtMarket := now.In(mh.TimeZone)

	if !mh.Days[nowAtMarket.Weekday()] {
		return false
	}

	midnight := time.Date(nowAtMarket.Year(), nowAtMarket.Month(), nowAtMarket.Day(), 0, 0, 0, 0, mh.TimeZone)
	timeSinceMidnight := nowAtMarket.Sub(midnight)

	// Check if current time is within trading hours
	return timeSinceMidnight >= mh.OpenTime && timeSinceMidnight < mh.CloseTime
}
