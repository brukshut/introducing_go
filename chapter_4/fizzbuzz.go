package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Print(i)
			fmt.Print(" ==> Fizz")
			if i%5 == 0 {
				fmt.Print("Buzz")
			}
			fmt.Print("\n")
		} else if i%5 == 0 {
			fmt.Print(i)
			fmt.Println(" ==> Buzz")
		}
	}
}
