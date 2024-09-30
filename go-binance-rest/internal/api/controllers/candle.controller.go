package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hoangtm1601/go-binance-rest/internal/api/services"
	"github.com/hoangtm1601/go-binance-rest/internal/initializers"
	"github.com/hoangtm1601/go-binance-rest/internal/models/dto"
	"github.com/hoangtm1601/go-binance-rest/utils"
)

type CandleController struct {
	service *services.CandleService
}

func NewCandleController(service *services.CandleService) *CandleController {
	return &CandleController{service: service}
}

// GetCandlesWithIndicators godoc
// @Summary Get candles with indicators
// @Description Retrieve candles with calculated indicators for a given time range and symbol
// @Tags candles
// @Accept json
// @Produce json
// @Param			payload	query		dto.GetCandlesWithIndicatorsDTO			false	"GetCandlesWithIndicators payload"
// @Success 200 {array} dto.CandleWithIndicators
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security		Bearer
// @Router /candles/indicators [get]
func (cc *CandleController) GetCandlesWithIndicators(c *gin.Context) {
	var query dto.GetCandlesWithIndicatorsDTO
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	queryBytes, err := json.Marshal(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal query"})
		return
	}

	nc := initializers.GetNatsConnection()
	rep, err := nc.Request("BINANCE.CANDLE.INDEX", queryBytes, 60*time.Second)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get response from NATS"})
		return
	}

	candles, err := utils.DecodeNatsResponse[[]dto.CandleResponseDTO](rep.Data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	candleWithIndicators, _ := cc.service.GetCandlesWithIndicators(candles, query.Period)

	c.JSON(http.StatusOK, candleWithIndicators)
}
