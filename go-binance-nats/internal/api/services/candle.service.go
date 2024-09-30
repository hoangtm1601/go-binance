package services

import (
	"github.com/hoangtm1601/go-binance-nats/internal/api/repositories"
	"github.com/hoangtm1601/go-binance-nats/internal/models"
)

type CandleService struct {
	repo *repositories.CandleRepository
}

func NewCandleService(repo *repositories.CandleRepository) *CandleService {
	return &CandleService{repo: repo}
}

func (s *CandleService) GetCandleByID(id uint) (*models.Candle, error) {
	return s.repo.GetByID(id)
}

func (s *CandleService) GetCandlesBySymbol(symbol string) ([]models.Candle, error) {
	return s.repo.GetBySymbol(symbol)
}

func (s *CandleService) GetLatestCandleByInterval(symbol string, interval models.CandleInterval) (*models.Candle, error) {
	return s.repo.GetLatestCandleByInterval(symbol, interval)
}

func (s *CandleService) GetCandlesByTimeRange(startDate, endDate int64, symbol string, interval models.CandleInterval) ([]models.Candle, error) {
	return s.repo.GetCandlesByTimeRange(startDate, endDate, symbol, interval)
}
