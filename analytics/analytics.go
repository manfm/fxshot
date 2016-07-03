package analytics

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
