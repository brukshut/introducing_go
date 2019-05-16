package main

import "fmt"

func greatest(args ...int) int {
	greatest := 0
	for _, v := range args {
		if v > greatest {
			greatest = v
		}
	}
	return greatest
}

func main() {
	fmt.Println(greatest(11, 22, 33))
	fmt.Println(greatest(101, 111, 103, 11, 12, 13))
}
