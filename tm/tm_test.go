package tm

import (
  "testing"
  "github.com/manfm/fxshot/fx"
  "github.com/manfm/fxshot/analytics"
)

func TestStrongTrendSignal(t *testing.T) {
  bull := []fx.Record{
		{"", "", 1, 2, 1, 2, 0},
		{"", "", 2, 3, 2, 3, 0},
		{"", "", 3, 4, 3, 4, 0},
		{"", "", 4, 5, 4, 5, 0},
		{"", "", 5, 6, 5, 6, 0},
		{"", "", 6, 7, 6, 7, 0},
	}

  ma := []float64{1, 2, 3, 4, 5, 6}
  trend := analytics.IdentifyTrend(bull)

  result := StrongTrendSignal(bull, ma, trend)
	if result != 1 {
		t.Error("StrongTrendSignal failed", result)
	}

  bear := []fx.Record{
		{"", "", 7, 7, 6, 6, 0},
		{"", "", 6, 6, 5, 5, 0},
		{"", "", 5, 5, 4, 4, 0},
		{"", "", 3, 3, 2, 2, 0},
	}

  ma = []float64{7, 6, 5, 4}
  trend = analytics.IdentifyTrend(bear)

  result = StrongTrendSignal(bear, ma, trend)
	if result != -1 {
		t.Error("StrongTrendSignal failed", result)
	}

  records := []fx.Record{
    {"", "", 6, 6, 4, 4, 0},
  	{"", "", 4, 4, 2, 2, 0},
    {"", "", 2, 2, 1, 1, 0},
		{"", "", 1, 2, 1, 2, 0},
		{"", "", 2, 3, 2, 3, 0},
		{"", "", 3, 4, 3, 4, 0},
		{"", "", 4, 5, 4, 5, 0},
	}

  ma = []float64{2.1, 2.0, 1.9, 1.95, 2.0, 2.2, 2.3}
  trend = analytics.IdentifyTrend(records)

  result = StrongTrendSignal(records, ma, trend)
	if result != 1 {
		t.Error("StrongTrendSignal failed", result)
	}
}
