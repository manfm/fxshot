package main

import (
  "fmt"
	"github.com/manfm/fxshot/fx"
  "github.com/manfm/fxshot/numbers"
)

func main() {
	 records := fx.Import("data/EURUSD5.csv")

	for _, r := range records {
		fmt.Printf("%f - %f = %f \n", r.Open, r.Close, numbers.RoundPlus(float64(r.Open-r.Close), 5))
	}
}
