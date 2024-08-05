package mathfuncs

func Variance(s []float64) float64 {
	var sqDiff float64
	mean := Average(s)
	for _, n := range s {
		diff := n - mean
		sqDiff += diff * diff
	}
	return sqDiff / float64(len(s))
}
