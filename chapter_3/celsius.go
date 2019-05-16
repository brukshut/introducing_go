package main

import "fmt"

func main() {
	// our intput variable
	var input float64

	// prompt for fahrenheit
	fmt.Print("Convert Fahrenheit Temperature: ")
	fmt.Scanf("%f", &input)

	// convert to celsius
	output := (((input - 32) * 5) / 9)
	fmt.Println(output)
}
