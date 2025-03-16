package repository

import (
	"plutus/domain"
	"time"
)

type QuoteEntity struct {
	ID               uint      `gorm:"primaryKey"`
	Symbol           string    `gorm:"type:text;not null;uniqueIndex:idx_symbol_datetime"`
	Name             string    `gorm:"type:text"`
	Exchange         string    `gorm:"type:text"`
	MicCode          string    `gorm:"type:text"`
	Currency         string    `gorm:"type:text"`
	Datetime         time.Time `gorm:"type:timestamp;not null;uniqueIndex:idx_symbol_datetime"`
	Timestamp        int64     `gorm:"type:bigint;not null"`
	Open             float64   `gorm:"type:double precision;not null"`
	High             float64   `gorm:"type:double precision;not null"`
	Low              float64   `gorm:"type:double precision;not null"`
	Close            float64   `gorm:"type:double precision;not null"`
	Volume           int64     `gorm:"type:bigint;not null"`
	PreviousClose    float64   `gorm:"type:double precision;not null"`
	Change           float64   `gorm:"type:double precision;not null"`
	PercentChange    float64   `gorm:"type:double precision;not null"`
	AverageVolume    int64     `gorm:"type:bigint;not null"`
	IsMarketOpen     bool      `gorm:"type:boolean;not null"`
	FiftyTwoWeekHigh float64   `gorm:"column:fifty_two_week_high;type:double precision;not null"`
	FiftyTwoWeekLow  float64   `gorm:"column:fifty_two_week_low;type:double precision;not null"`
}

func (qe *QuoteEntity) ToQuote() domain.Quote {
	return domain.Quote{
		Symbol:        qe.Symbol,
		Name:          qe.Name,
		Exchange:      qe.Exchange,
		MicCode:       qe.MicCode,
		Currency:      qe.Currency,
		Datetime:      qe.Datetime,
		Timestamp:     qe.Timestamp,
		Open:          qe.Open,
		High:          qe.High,
		Low:           qe.Low,
		Close:         qe.Close,
		Volume:        qe.Volume,
		PreviousClose: qe.PreviousClose,
		Change:        qe.Change,
		PercentChange: qe.PercentChange,
		AverageVolume: qe.AverageVolume,
		IsMarketOpen:  qe.IsMarketOpen,
		FiftyTwoWeek: struct {
			High float64
			Low  float64
		}{
			High: qe.FiftyTwoWeekHigh,
			Low:  qe.FiftyTwoWeekLow,
		},
	}
}

func ToQuoteEntity(q *domain.Quote) QuoteEntity {
	return QuoteEntity{
		Symbol:           q.Symbol,
		Name:             q.Name,
		Exchange:         q.Exchange,
		MicCode:          q.MicCode,
		Currency:         q.Currency,
		Datetime:         q.Datetime,
		Timestamp:        q.Timestamp,
		Open:             q.Open,
		High:             q.High,
		Low:              q.Low,
		Close:            q.Close,
		Volume:           q.Volume,
		PreviousClose:    q.PreviousClose,
		Change:           q.Change,
		PercentChange:    q.PercentChange,
		AverageVolume:    q.AverageVolume,
		IsMarketOpen:     q.IsMarketOpen,
		FiftyTwoWeekHigh: q.FiftyTwoWeek.High,
		FiftyTwoWeekLow:  q.FiftyTwoWeek.Low,
	}
}
