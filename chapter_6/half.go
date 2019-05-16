package main

import (
	"fmt"
)

func half(number int) (int, bool) {

	var remainder int

	remainder = (number / 2) % 2

	if remainder == 0 {
		return (number / 2), true
	} else {
		return (number / 2), false
	}
}

func main() {
	var halfnum int
	var success bool

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	for _, number := range numbers {
		halfnum, success = half(number)
		fmt.Println(number, halfnum, success)
	}
}
