package mathfuncs

import "math"

// Range calculates the lower and upper limits for the given data
// based on the mean and standard deviation of the last four elements (or fewer if less than four elements exist).
func Range(data []float64) (int, int) {
	start := len(data) - 4  // Determine the starting index for the last four elements
	if start < 0 {          // If the starting index is negative, set it to 0
		start = 0
	}

	window := data[start:]   // Create a slice containing the last four (or fewer) elements
	mean := Average(window)  // Calculate the mean of the window slice
	stdDev := StdDeviation(window)  // Calculate the standard deviation of the window slice

	// Calculate the lower and upper limits based on the mean and standard deviation
	lowerLimit := mean - (2 * stdDev)
	upperLimit := mean + (2 * stdDev)

	// Round the limits to the nearest integer and return them
	return int(math.Round(lowerLimit)), int(math.Round(upperLimit))
}
