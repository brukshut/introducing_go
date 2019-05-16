package main

import "fmt"
import m "golang-book/chapter8/math"

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := m.Average(xs)
	max := m.Maximum(xs)
	min := m.Minimum(xs)
	fmt.Println("Average is", avg)
	fmt.Println("Maximum is", max)
	fmt.Println("Minimum is", min)
}
