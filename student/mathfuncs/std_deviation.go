package mathfuncs

import "math"

func StdDeviation(s []float64) float64 {
	return math.Sqrt(Variance(s))
}
