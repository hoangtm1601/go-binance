package services

import (
	"fmt"
	"github.com/hoangtm1601/go-binance-rest/internal/models/dto"
	"math"
	"sync"
)

type CandleService struct {
}

func NewCandleService() *CandleService {
	return &CandleService{}
}

func (s *CandleService) CalculateDEMA(candles []dto.CandleResponseDTO, period int) []float64 {
	if len(candles) < period {
		return nil
	}

	// Extract closing prices
	closePrices := make([]float64, len(candles))
	for i, candle := range candles {
		closePrices[i] = candle.Cl
	}

	// Calculate first EMA
	ema1 := s.calculateEMA(closePrices, period)

	// Calculate EMA of EMA
	ema2 := s.calculateEMA(ema1, period)

	// Calculate DEMA
	dema := make([]float64, len(ema1))
	for i := range dema {
		if i < period-1 {
			dema[i] = 0 // DEMA is undefined for the first (period-1) elements
		} else {
			dema[i] = 2*ema1[i] - ema2[i]
		}
	}

	return dema
}

func (s *CandleService) calculateEMA(data []float64, period int) []float64 {
	ema := make([]float64, len(data))
	k := 2.0 / float64(period+1)

	// Initialize EMA with SMA for the first 'period' elements
	sum := 0.0
	for i := 0; i < period; i++ {
		sum += data[i]
		ema[i] = 0 // EMA is undefined for the first (period-1) elements
	}
	ema[period-1] = sum / float64(period)

	// Calculate EMA for the rest of the data
	for i := period; i < len(data); i++ {
		ema[i] = data[i]*k + ema[i-1]*(1-k)
	}

	return ema
}

func (s *CandleService) CalculateMA(candles []dto.CandleResponseDTO, period int) []float64 {
	if len(candles) < period {
		return nil
	}

	ma := make([]float64, len(candles))

	// Calculate the sum for the first period
	sum := 0.0
	for i := 0; i < period; i++ {
		sum += candles[i].Cl
	}

	// Calculate MA for the first complete period
	ma[period-1] = sum / float64(period)

	// Calculate MA for the rest of the data
	for i := period; i < len(candles); i++ {
		sum = sum - candles[i-period].Cl + candles[i].Cl
		ma[i] = sum / float64(period)
	}

	return ma
}

func (s *CandleService) CalculateSMA(candles []dto.CandleResponseDTO, period int) []float64 {
	if len(candles) < period {
		return nil
	}

	sma := make([]float64, len(candles))

	// Calculate the sum for the first period
	sum := 0.0
	for i := 0; i < period; i++ {
		sum += candles[i].Cl
	}

	// Calculate SMA for the first complete period
	sma[period-1] = sum / float64(period)

	// Calculate SMA for the rest of the data
	for i := period; i < len(candles); i++ {
		sum = sum - candles[i-period].Cl + candles[i].Cl
		sma[i] = sum / float64(period)
	}

	return sma
}

func (s *CandleService) CalculateRSI(candles []dto.CandleResponseDTO, period int) []float64 {
	if len(candles) < period+1 {
		return nil
	}

	rsi := make([]float64, len(candles))
	gains := make([]float64, len(candles))
	losses := make([]float64, len(candles))

	// Calculate initial gains and losses
	for i := 1; i < len(candles); i++ {
		change := candles[i].Cl - candles[i-1].Cl
		if change > 0 {
			gains[i] = change
		} else {
			losses[i] = -change
		}
	}

	// Calculate average gains and losses
	gainCandles := make([]dto.CandleResponseDTO, period)
	lossCandles := make([]dto.CandleResponseDTO, period)
	for i := 0; i < period; i++ {
		gainCandles[i] = dto.CandleResponseDTO{Cl: gains[i+1]}
		lossCandles[i] = dto.CandleResponseDTO{Cl: losses[i+1]}
	}
	avgGain := s.CalculateSMA(gainCandles, period)[period-1]
	avgLoss := s.CalculateSMA(lossCandles, period)[period-1]

	// Calculate initial RSI
	rs := avgGain / avgLoss
	rsi[period] = 100 - (100 / (1 + rs))

	// Calculate RSI for the rest of the data
	for i := period + 1; i < len(candles); i++ {
		avgGain = (avgGain*(float64(period)-1) + gains[i]) / float64(period)
		avgLoss = (avgLoss*(float64(period)-1) + losses[i]) / float64(period)
		rs = avgGain / avgLoss
		rsi[i] = 100 - (100 / (1 + rs))
	}

	return rsi
}

func (s *CandleService) CalculateBollingerBands(candles []dto.CandleResponseDTO, period int, stdDev float64) ([]float64, []float64, []float64) {
	upper := make([]float64, len(candles))
	middle := make([]float64, len(candles))
	lower := make([]float64, len(candles))

	// Calculate SMA for all candles
	sma := s.CalculateSMA(candles, period)

	for i := period - 1; i < len(candles); i++ {
		if i >= period-1 {
			sum := 0.0
			for j := 0; j < period; j++ {
				diff := candles[i-j].Cl - sma[i]
				sum += diff * diff
			}
			stdDeviation := math.Sqrt(sum / float64(period))

			upper[i] = sma[i] + stdDev*stdDeviation
			middle[i] = sma[i]
			lower[i] = sma[i] - stdDev*stdDeviation
			continue
		}
		upper[i] = 0
		middle[i] = sma[i]
		lower[i] = 0
	}

	return upper, middle, lower
}

func (s *CandleService) CalculateMACD(candles []dto.CandleResponseDTO, fastPeriod, slowPeriod, signalPeriod int) ([]float64, []float64, []float64) {
	closePrices := make([]float64, len(candles))
	for i, candle := range candles {
		closePrices[i] = candle.Cl
	}

	fastEMA := s.calculateEMA(closePrices, fastPeriod)
	slowEMA := s.calculateEMA(closePrices, slowPeriod)

	macdLine := make([]float64, len(candles))
	for i := 0; i < len(candles); i++ {
		macdLine[i] = fastEMA[i] - slowEMA[i]
	}

	signalLine := s.calculateEMA(macdLine, signalPeriod)

	histogram := make([]float64, len(candles))
	for i := 0; i < len(candles); i++ {
		histogram[i] = macdLine[i] - signalLine[i]
	}

	return macdLine, signalLine, histogram
}

func (s *CandleService) CalculateStochasticOscillator(candles []dto.CandleResponseDTO, period, smoothK, smoothD int) ([]float64, []float64) {
	if len(candles) < period {
		return nil, nil
	}

	k := make([]float64, len(candles))
	d := make([]float64, len(candles))

	for i := period - 1; i < len(candles); i++ {
		highestHigh := candles[i].Hi
		lowestLow := candles[i].Lo

		for j := 0; j < period; j++ {
			if candles[i-j].Hi > highestHigh {
				highestHigh = candles[i-j].Hi
			}
			if candles[i-j].Lo < lowestLow {
				lowestLow = candles[i-j].Lo
			}
		}

		k[i] = (candles[i].Cl - lowestLow) / (highestHigh - lowestLow) * 100
	}

	// Smooth K values
	kCandles := make([]dto.CandleResponseDTO, len(k))
	for i, v := range k {
		kCandles[i] = dto.CandleResponseDTO{Cl: v}
	}
	k = s.CalculateSMA(kCandles[period-1:], smoothK)

	// Calculate D values (SMA of K)
	dCandles := make([]dto.CandleResponseDTO, len(k))
	for i, v := range k {
		dCandles[i] = dto.CandleResponseDTO{Cl: v}
	}
	d = s.CalculateSMA(dCandles, smoothD)

	slippingArr := make([]float64, period-1)
	for i := 0; i < period-2; i++ {
		slippingArr[i] = 0
	}
	// Initialize the first period-1 elements to 0
	k = append(slippingArr, k...)
	d = append(slippingArr, d...)

	return k, d
}

func (s *CandleService) GetCandlesWithIndicators(candles []dto.CandleResponseDTO, period int) ([]dto.CandleWithIndicators, error) {
	if len(candles) < period {
		return nil, fmt.Errorf("not enough candles for the specified period")
	}

	var wg sync.WaitGroup

	var sma, dema, rsi, macd, signal []float64
	//var sma, dema, rsi []float64
	var bollingerUpper, bollingerMiddle, bollingerLower []float64
	var stochasticK, stochasticD []float64

	wg.Add(1)
	go func() {
		defer wg.Done()
		sma = s.CalculateSMA(candles, period)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		dema = s.CalculateDEMA(candles, period)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		rsi = s.CalculateRSI(candles, period)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		macd, signal, _ = s.CalculateMACD(candles, 12, 26, 9)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		bollingerUpper, bollingerMiddle, bollingerLower = s.CalculateBollingerBands(candles, period, 2)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		stochasticK, stochasticD = s.CalculateStochasticOscillator(candles, period, 3, 3)
	}()

	// Wait for all goroutines to finish
	wg.Wait()

	// Prepare the result
	result := make([]dto.CandleWithIndicators, len(candles))
	for i, candle := range candles {
		result[i] = dto.CandleWithIndicators{
			Candle: candle,
			Indicators: dto.IndicatorResult{
				SMA:             sma[i],
				DEMA:            dema[i],
				RSI:             rsi[i],
				MACD:            macd[i],
				Signal:          signal[i],
				BollingerUpper:  bollingerUpper[i],
				BollingerMiddle: bollingerMiddle[i],
				BollingerLower:  bollingerLower[i],
				StochasticK:     stochasticK[i],
				StochasticD:     stochasticD[i],
			},
		}
	}

	return result, nil
}
