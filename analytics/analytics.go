package analytics

import (
	"github.com/manfm/fxshot/fx"
	"math"
)

func MA(data []float64, period int) []float64 {
	result := make([]float64, len(data))
	for k, _ := range data {
		pos := k + 1
		if pos < period {
			result[k] = 0.0
			continue
		}
		start_i := pos - period
		end_i := pos
		sub := data[start_i:end_i]
		result[k] = Average(sub)
	}
	return result
}

func Average(data []float64) float64 {
	sum := 0.0
	for _, v := range data {
		sum = sum + v
	}
	average := sum / float64(len(data))

	return average
}

func PercentageMovement(data []float64) []float64 {
	result := make([]float64, len(data))
	for k, _ := range data {
		if k == 0 {
			result[k] = 0.0
			continue
		}
		result[k] = (data[k]/data[k-1])*100 - 100
	}
	return result
}

func IdentifyTrend(data []fx.Record) []bool {
	result := make([]bool, len(data))
	for k, _ := range data {
		if k == 0 {
			result[k] = false
			continue
		}
		trend := data[k].Close >= data[k-1].Open
		result[k] = trend
	}
	return result
}

func Volatility(data []fx.Record) []float64 {
	result := make([]float64, len(data))
	for k, r := range data {
		result[k] = math.Abs(float64(r.High-r.Low)) - math.Abs(float64(r.Open-r.Close))
	}
	return result
}
