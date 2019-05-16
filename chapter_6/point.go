package main

import "fmt"

func main() {
	var x int
	x = 11
	y := &x
	fmt.Printf("Memory address: %p\n", y)
	fmt.Printf("Memory address: %x\n", y)
	fmt.Println("Dereferenced:", *y)
	*y = 44
	fmt.Println("x is now:", x)
}
