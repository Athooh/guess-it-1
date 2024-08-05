package mathfuncs

func Sorted(s []float64) []float64 {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] > s[j] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
	return s
}

func Median(s []float64) float64 {
	var med float64

	Sorted(s)
	if len(s)%2 == 1 {
		idx := len(s) / 2
		med = s[idx]
	} else if len(s)%2 == 0 {
		medTot := s[len(s)/2] + s[len(s)/2-1]
		med = medTot / 2
	}

	return med
}
