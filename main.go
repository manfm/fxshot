package main

import (
	"fmt"
	"github.com/manfm/fxshot/analytics"
	"github.com/manfm/fxshot/fx"
	"github.com/manfm/fxshot/numbers"
	"github.com/manfm/fxshot/tm"
)

func main() {
	records := fx.Import("data/EURUSD5.csv")

	prices := make([]float64, len(records))
	for k, r := range records {
		prices[k] = r.Close
	}

	ma := analytics.MA(prices, 14)
	// pm := analytics.PercentageMovement(prices)
	tr := analytics.IdentifyTrend(records)
	// sv := analytics.SpikeVolatility(records)
	// hlv := analytics.HighLowVolatility(records)

	for i, r := range records {
		from := i - 5
		if from < 0 {
			continue
		}

		fmt.Println(r, numbers.RoundPlus(ma[i], 5), tr[i], tm.StrongTrendSignal(records[from:i+1], ma[from:i+1], tr[from:i+1]))
	}
}
