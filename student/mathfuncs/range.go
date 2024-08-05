package mathfuncs

import "math"

func Range(data []float64) (int, int) {
	start := len(data) - 4
	if start < 0 {
		start = 0
	}

	window := data[start:]
	mean := Average(window)
	stdDev := StdDeviation(window)

	lowerLimit := mean - (2 * stdDev)
	upperLimit := mean + (2 * stdDev)

	return int(math.Round(lowerLimit)), int(math.Round(upperLimit))
}
