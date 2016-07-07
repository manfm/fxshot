package tm //trading model

import (
	"github.com/manfm/fxshot/fx"
)

func StrongTrendSignal(data []fx.Record, ma []float64, trend []bool) int {
	history_view := 3
	bull := 0
	bear := 0
	l := len(data)
	// check last 3 records
	for i := (l - history_view); i < l; i++ {
		if data[i].Low >= ma[i] && trend[i] == true {
			bull++
		} else if data[i].High <= ma[i] && trend[i] == false {
			bear++
		}
	}

	if bull == history_view {
		return 1
	} else if bear == history_view {
		return -1
	} else {
		return 0
	}
}
