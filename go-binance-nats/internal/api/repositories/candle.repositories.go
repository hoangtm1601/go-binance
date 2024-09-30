package repositories

import (
	"errors"
	"time"

	"github.com/hoangtm1601/go-binance-nats/internal/models"
	"gorm.io/gorm"
)

type CandleRepository struct {
	db *gorm.DB
}

func NewCandleRepository(db *gorm.DB) *CandleRepository {
	return &CandleRepository{db: db}
}

func (r *CandleRepository) Create(candle *models.Candle) error {
	return r.db.Create(candle).Error
}

func (r *CandleRepository) InsertMany(candles []*models.Candle) error {
	return r.db.Create(candles).Error
}

func (r *CandleRepository) GetByID(id uint) (*models.Candle, error) {
	var candle models.Candle
	err := r.db.First(&candle, id).Error
	return &candle, err
}

func (r *CandleRepository) GetBySymbol(symbol string) ([]models.Candle, error) {
	var candles []models.Candle
	err := r.db.Where("symbol = ?", symbol).Find(&candles).Error
	return candles, err
}

func (r *CandleRepository) GetLatestCandleByInterval(symbol string, interval models.CandleInterval) (*models.Candle, error) {
	var candle models.Candle
	err := r.db.Where("symbol = ? and interval = ?", symbol, interval).Order("start DESC").First(&candle).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &candle, err
}

func (r *CandleRepository) GetCandlesBySymbolIntervalAndDateRange(symbol string, interval models.CandleInterval, startDate, endDate time.Time) ([]models.Candle, error) {
	var candles []models.Candle
	err := r.db.Where("symbol = ? AND interval = ? AND start >= ? AND end <= ?", symbol, interval, startDate, endDate).
		Order("start").
		Find(&candles).Error
	return candles, err
}

func (r *CandleRepository) GetCandlesByTimeRange(startDate, endDate int64, symbol string, interval models.CandleInterval) ([]models.Candle, error) {
	var candles []models.Candle
	err := r.db.Select("symbol, interval, start, \"end\", op, hi, lo, cl, bv, qv, tbv, tqv, cnt").Where("symbol = ? AND interval = ? AND start >= ? AND \"end\" <= ?", symbol, interval, startDate, endDate).
		Find(&candles).Error
	return candles, err
}
