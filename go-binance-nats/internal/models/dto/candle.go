package dto

import (
	"time"

	"github.com/hoangtm1601/go-binance-nats/internal/models"
)

// GetCandlesDTO represents the query parameters for getting candles with indicators
type GetCandlesDTO struct {
	StartDate time.Time             `form:"start_date" binding:"required" time_format:"2006-01-02T15:04:05Z07:00" example:"2024-09-20T00:00:00Z"`
	EndDate   time.Time             `form:"end_date" binding:"required" time_format:"2006-01-02T15:04:05Z07:00" example:"2024-09-25T00:00:00Z"`
	Symbol    string                `form:"symbol" binding:"required" example:"BTCUSDT"`
	Interval  models.CandleInterval `form:"interval" binding:"required,oneof=1min 5min 15min 30min 60min 240min 720min 1440min" example:"1min"`
}

// CandleResponseDTO represents a candle
type CandleResponseDTO struct {
	Symbol   string                `json:"symbol"`
	Interval models.CandleInterval `json:"interval"`
	Start    int64                 `json:"start"`
	End      int64                 `json:"end"`
	LastEnd  int64                 `json:"lastEnd"`
	Op       float64               `json:"op"`
	Hi       float64               `json:"hi"`
	Lo       float64               `json:"lo"`
	Cl       float64               `json:"cl"`
	Bv       float64               `json:"bv"`
	Qv       float64               `json:"qv"`
	Tbv      float64               `json:"tbv"`
	Tqv      float64               `json:"tqv"`
	Cnt      int64                 `json:"cnt"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error"`
}
