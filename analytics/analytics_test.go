package analytics

import (
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
