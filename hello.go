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

	prices := make([]float64, len(records))
	for k, r := range records {
		prices[k] = float64(r.Close)
	}

	ma := analytics.MA(prices, 16)
	pm := analytics.PercentageMovement(prices)
	it := analytics.IdentifyTrend(records)
	v := analytics.Volatility(records)

	for i, r := range records {
		fmt.Println(r, numbers.RoundPlus(ma[i], 5), numbers.RoundPlus(pm[i], 4), numbers.RoundPlus(v[i], 5), it[i])
	}
}
