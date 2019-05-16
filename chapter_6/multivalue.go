package main

import "fmt"

func f() (r int, e int) {
	e = 5
	r = 1
	return
}

func main() {
	x, y := f()
	fmt.Println(x)
	fmt.Println(y)
}
