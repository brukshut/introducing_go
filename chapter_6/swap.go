package main

// write a program that can swap two integers
// (x :=1, y := 2; swap(&x, &y)

import "fmt"

func swap(x *int, y *int) (int, int) {
	// return deferenced pointers
	return *y, *x
}

func main() {
	x := 1
	y := 2
	fmt.Println("x is", x)
	fmt.Println("y is", y)
	fmt.Printf("swap x and y\n")
	x, y = swap(&x, &y)
	fmt.Printf("x is now %d\n", x)
	fmt.Printf("y is now %d\n", y)
}
