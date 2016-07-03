package main

import (
	"fmt"
	"github.com/manfm/fxshot/analytics"
	"github.com/manfm/fxshot/fx"
	"github.com/manfm/fxshot/numbers"
)

func main() {
	records := fx.Import("data/EURUSD5.csv")

	for _, r := range records {
		fmt.Printf("%f - %f = %f \n", r.Open, r.Close, numbers.RoundPlus(float64(r.Open-r.Close), 5))
	}

	for _, r := range ma16(records) {
		fmt.Println(r)
	}
}

func ma16(records []fx.Record) []float64 {
	prices := make([]float64, len(records))

	for k, r := range records {
		prices[k] = float64(r.Close)
	}

	return analytics.MA(prices, 16)
}
