package mathfuncs

import (
	"math"
	"reflect"
	"testing"
)

func TestAverage(t *testing.T) {
	tests := []struct {
		data     []float64
		expected float64
	}{
		{[]float64{1, 2, 3, 4, 5}, 3.0},
		{[]float64{-1, -2, -3, -4, -5}, -3.0},
		{[]float64{-1, 0, 1, 2}, 0.5},
		{[]float64{}, math.NaN()},
	}
	for _, tt := range tests {
		result := Average(tt.data)
		if math.IsNaN(tt.expected) && !math.IsNaN(result) {
			t.Errorf("Average calculated is incorrect, got: %f, wants: %f.", result, tt.expected)
		} else if !math.IsNaN(tt.expected) && result != tt.expected {
			t.Errorf("Average calculated is incorrect, got: %f, wants: %f.", result, tt.expected)
		}
	}
}

func TestSorted(t *testing.T) {
	tests := []struct {
		data     []float64
		expected []float64
	}{
		{[]float64{1, 2, 3, 4, 5}, []float64{1, 2, 3, 4, 5}},
		{[]float64{5, 4, 3, 2, 1}, []float64{1, 2, 3, 4, 5}},
		{[]float64{3, 1, 4, 5, 2}, []float64{1, 2, 3, 4, 5}},
		{[]float64{2, 2, 2, 2}, []float64{2, 2, 2, 2}},
		{[]float64{1}, []float64{1}},
		{[]float64{}, []float64{}},
	}
	for _, tt := range tests {
		result := Sorted(tt.data)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("Sorted calculated is incorrect, got: %v, wants: %v.", result, tt.expected)
		}
	}
}

func TestMedian(t *testing.T) {
	tests := []struct {
		data     []float64
		expected float64
	}{
		{[]float64{1, 3, 5}, 3.0},
		{[]float64{1, 2, 3, 4}, 2.5},
		{[]float64{3, 1, 4, 2}, 2.5},
		{[]float64{-5, -1, -3}, -3.0},
		{[]float64{1}, 1.0},
		//{[]float64{}, math.NaN()},
	}
	for _, tt := range tests {
		result := Median(tt.data)
		if math.IsNaN(tt.expected) && !math.IsNaN(result) {
			t.Errorf("Median calculated is incorrect, got: %f, wants: %f.", result, tt.expected)
		} else if !math.IsNaN(tt.expected) && result != tt.expected {
			t.Errorf("Median calculated is incorrect, got: %f, wants: %f.", result, tt.expected)
		}
	}
}

func TestVariance(t *testing.T) {
	tests := []struct {
		data     []float64
		expected float64
	}{
		{[]float64{1, 2, 3, 4, 5}, 2.0},
		{[]float64{-1, -2, -3, -4, -5}, 2.0},
		{[]float64{-1, 0, 1, 2}, 1.25},
		{[]float64{2, 2, 2, 2}, 0.0},
		{[]float64{1}, 0.0},
		{[]float64{}, math.NaN()},
	}
	for _, tt := range tests {
		result := Variance(tt.data)
		if math.IsNaN(tt.expected) && !math.IsNaN(result) {
			t.Errorf("Variance calculated is incorrect, got: %f, wants: %f.", result, tt.expected)
		} else if !math.IsNaN(tt.expected) && math.Abs(result-tt.expected) > 1e-9 {
			t.Errorf("Variance calculated is incorrect, got: %f, wants: %f.", result, tt.expected)
		}
	}
}

func TestStdDeviation(t *testing.T) {
	tests := []struct {
		data     []float64
		expected float64
	}{
		{[]float64{1, 2, 3, 4, 5}, math.Sqrt(2.0)},
		{[]float64{-1, -2, -3, -4, -5}, math.Sqrt(2.0)},
		{[]float64{-1, 0, 1, 2}, math.Sqrt(1.25)},
		{[]float64{2, 2, 2, 2}, 0.0},
		{[]float64{1}, 0.0},
		{[]float64{}, math.NaN()},
	}
	for _, tt := range tests {
		result := StdDeviation(tt.data)
		if math.IsNaN(tt.expected) && !math.IsNaN(result) {
			t.Errorf("StdDeviation calculated is incorrect, got: %f, wants: %f.", result, tt.expected)
		} else if !math.IsNaN(tt.expected) && math.Abs(result-tt.expected) > 1e-9 {
			t.Errorf("StdDeviation calculated is incorrect, got: %f, wants: %f.", result, tt.expected)
		}
	}
}
