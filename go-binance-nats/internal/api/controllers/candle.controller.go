package controllers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/nats-io/nats.go"

	"github.com/hoangtm1601/go-binance-nats/internal/api/services"
	"github.com/hoangtm1601/go-binance-nats/internal/middleware"
	"github.com/hoangtm1601/go-binance-nats/internal/models/dto"
)

type CandleController struct {
	service *services.CandleService
	nc      *nats.Conn
}

func NewCandleController(service *services.CandleService, nc *nats.Conn) *CandleController {
	return &CandleController{
		service: service,
		nc:      nc,
	}
}

// SetupSubscriptions sets up NATS subscriptions for the CandleController
func (cc *CandleController) SetupSubscriptions() error {
	_, err := cc.nc.Subscribe("BINANCE.CANDLE.INDEX", middleware.LoggingMiddleware(
		middleware.CompressAndEncodeMiddleware(cc.GetCandlesWithIndicators),
	))
	if err != nil {
		return fmt.Errorf("failed to subscribe to BINANCE.CANDLE.INDEX: %w", err)
	}
	log.Println("Subscribed to BINANCE.CANDLE.INDEX")
	return nil
}

func (cc *CandleController) GetCandlesWithIndicators(msg *nats.Msg) (interface{}, error) {
	// Decode the NATS message payload
	var query dto.GetCandlesDTO

	if err := json.Unmarshal(msg.Data, &query); err != nil {
		return nil, fmt.Errorf("error unmarshaling query: %w", err)
	}

	candles, err := cc.service.GetCandlesByTimeRange(
		query.StartDate.UnixMilli(),
		query.EndDate.UnixMilli(),
		query.Symbol,
		query.Interval,
	)
	if err != nil {
		return nil, fmt.Errorf("error getting candles with indicators: %w", err)
	}

	return candles, nil
}
