package main

import "fmt"

// simple function example
func average(numbers []float64) float64 {
	total := 0.0
	for _, value := range numbers {
		total += value
	}
	return total / float64(len(numbers))
}

func main() {
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(xs))
}
