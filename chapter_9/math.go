package math

// Finds the average of a series of numbers
func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

// Finds the minimum value in a slice of numbers
func Minimum(numbers []float64) float64 {
	minimum := Maximum(numbers)
	for _, num := range numbers {
		if num < minimum {
			minimum = num
		}
	}
	return minimum
}

// Finds the maximum value in a slice of numbers
func Maximum(numbers []float64) float64 {
	maximum := float64(0)
	for _, num := range numbers {
		if num > maximum {
			maximum = num
		}
	}
	return maximum
}
