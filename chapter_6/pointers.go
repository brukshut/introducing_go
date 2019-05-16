package main

import "fmt"

func pzero(x *int) {
	*x = 0
}

func zero(x int) {
	x = 0
}

func main() {
	x := 5
	zero(x)
	fmt.Println(x) // x is still 5

	pzero(&x)
	fmt.Println(x) // x is 0
}
