package main

import "fmt"

func main() {
	var feet float64

	fmt.Print("Enter number of feet (convert to meters): ")
	fmt.Scanf("%f", &feet)

	output := (feet * 0.3048)
	fmt.Println(output)

}
