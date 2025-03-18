package repository

import (
	"fmt"
	"plutus/domain"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// TODO TEST ALL

type QuoteRepository struct {
	db *gorm.DB
}

func NewQuoteRepository(connStr string) (*QuoteRepository, error) {
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("gorm open error: %v", err)
	}
	// TODO: GORM's AutoMigrate handles additive schema changes (e.g., new columns) but not rollbacks (e.g., dropping columns).
	// Later, consider integrating a tool like golang-migrate for full migration support with up/down scripts.
	if err := db.AutoMigrate(&QuoteEntity{}); err != nil {
		return nil, fmt.Errorf("auto migrate error: %v", err)
	}
	return &QuoteRepository{db: db}, nil
}

func (r *QuoteRepository) SaveOrUpdate(quote domain.Quote) (uint, error) {
	entity := ToQuoteEntity(&quote)
	result := r.db.Save(&entity) // Save upserts based on ID or unique constraints
	if result.Error != nil {
		return 0, fmt.Errorf("create error: %v", result.Error)
	}
	return entity.ID, nil
}

func (r *QuoteRepository) Read(symbol string, datetime time.Time) (domain.Quote, error) {
	var entity QuoteEntity
	result := r.db.Where("symbol = ? AND datetime = ?", symbol, datetime).First(&entity)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return domain.Quote{}, fmt.Errorf("quote not found for symbol %s at %s", symbol, datetime)
		}
		return domain.Quote{}, fmt.Errorf("read error: %v", result.Error)
	}
	return entity.ToQuote(), nil
}

func (r *QuoteRepository) Update(quote domain.Quote) error {
	entity := ToQuoteEntity(&quote)
	result := r.db.Model(&QuoteEntity{}).
		Where("symbol = ? AND datetime = ?", quote.Symbol, quote.Datetime).
		Updates(entity)
	if result.Error != nil {
		return fmt.Errorf("update error: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no quote updated for symbol %s at %s", quote.Symbol, quote.Datetime)
	}
	return nil
}

func (r *QuoteRepository) Delete(symbol string, datetime time.Time) error {
	result := r.db.Where("symbol = ? AND datetime = ?", symbol, datetime).Delete(&QuoteEntity{})
	if result.Error != nil {
		return fmt.Errorf("delete error: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no quote deleted for symbol %s at %s", symbol, datetime)
	}
	return nil
}

func (r *QuoteRepository) GetAll(symbol string) ([]domain.Quote, error) {
	var entities []QuoteEntity
	result := r.db.Where("symbol = ?", symbol).Order("datetime DESC").Find(&entities)
	if result.Error != nil {
		return nil, fmt.Errorf("list error: %v", result.Error)
	}
	quotes := make([]domain.Quote, len(entities))
	for i, entity := range entities {
		quotes[i] = entity.ToQuote()
	}
	return quotes, nil
}
