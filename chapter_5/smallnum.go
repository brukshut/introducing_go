package main

// find smallest number in slice

import "fmt"

func main() {
	x := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}

	var smallest int
	for n := 0; n <= (len(x) - 1); n++ {
		if smallest == 0 {
			smallest = x[n]
		} else if n < smallest {
			smallest = x[n]
		}
	}
	fmt.Print(smallest)
	fmt.Println(" is the smallest number")
}
