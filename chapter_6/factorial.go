package main

import "fmt"

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func main() {
	x := uint(0)
	z := uint(3)
	fmt.Println(factorial(x))
	fmt.Println(factorial(z))
	fmt.Println(factorial(5))
	fmt.Println(factorial(7))
}
