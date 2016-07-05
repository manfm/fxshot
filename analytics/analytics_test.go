package analytics

import (
	"github.com/manfm/fxshot/fx"
	"reflect"
	"testing"
)

func TestMA(t *testing.T) {
	data := []float64{1, 2, 4, 8}
	expected := []float64{0.0, 1.5, 3, 6}
	result := MA(data, 2)
	if !reflect.DeepEqual(result, expected) {
		t.Error("MA failed", result)
	}

	data = []float64{1.2, 3.456, 3, 7}
	expected = []float64{0.0, 0.0, 2.552, 4.485333333333333}
	result = MA(data, 3)
	if !reflect.DeepEqual(result, expected) {
		t.Error("MA failed", result)
	}
}

func TestAverage(t *testing.T) {
	data := []float64{1, 2}
	expected := 1.5
	if Average(data) != expected {
		t.Error("Average failed")
	}
}

func TestPercentageMovement(t *testing.T) {
	data := []float64{1.5, 3, 2}
	expected := []float64{0, 100, -33.33333333333334}
	result := PercentageMovement(data)
	if !reflect.DeepEqual(result, expected) {
		t.Error("PercentageMovement failed", result)
	}
}

func TestIdentifyTrend(t *testing.T) {
	data := []fx.Record{
		{"", "", 1, 0, 0, 2, 0},
		{"", "", 2, 0, 0, 3, 0},
		{"", "", 3, 0, 0, 2, 0},
		{"", "", 2, 0, 0, 4, 0},
		{"", "", 4, 0, 0, 3, 0},
		{"", "", 3, 0, 0, 1, 0},
	}
	expected := []bool{false, true, true, true, true, false}
	result := IdentifyTrend(data)
	if !reflect.DeepEqual(result, expected) {
		t.Error("PercentageMovement failed", result)
	}
}

func TestSpikeVolatility(t *testing.T) {
	if SpikeVolatility([]fx.Record{{"", "", 1, 2, 1, 2, 0}})[0] != 0 {
		t.Error("Volatility failed")
	}

	if SpikeVolatility([]fx.Record{{"", "", 1, 2, 1, 1, 0}})[0] != 1.0 {
		t.Error("Volatility failed")
	}

	if SpikeVolatility([]fx.Record{{"", "", 1, 3, 0, 2, 0}})[0] != 2.0 {
		t.Error("Volatility failed")
	}

	if SpikeVolatility([]fx.Record{{"", "", 2, 2, 1, 1, 0}})[0] != 0.0 {
		t.Error("Volatility failed")
	}

	if SpikeVolatility([]fx.Record{{"", "", 2, 3, 0, 1, 0}})[0] != 2.0 {
		t.Error("Volatility failed")
	}
}

func TestHighLowVolatility(t *testing.T) {
	if HighLowVolatility([]fx.Record{{"", "", 0, 2, 1, 0, 0}})[0] != 1.0 {
		t.Error("Volatility failed")
	}

	if HighLowVolatility([]fx.Record{{"", "", 0, 2, 1, 0, 0}})[0] != 1.0 {
		t.Error("Volatility failed")
	}

	if HighLowVolatility([]fx.Record{{"", "", 0, 3, 0, 0, 0}})[0] != 3.0 {
		t.Error("Volatility failed")
	}
}
