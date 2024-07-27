package mathfuncs

func Average(s []float64) float64 {
	var total float64
	for _, c := range s {
		total += c
	}
	return total / float64(len(s))
}
