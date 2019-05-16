package main

import (
	"fmt"
	"strings"
)

func main() {
	// Contains(s, substr string) bool
	fmt.Println(strings.Contains("test", "es"))

	fmt.Println("found string 'es'", strings.Count("test", "es"), "time(s) string 'test'")
	fmt.Println("found string 't'", strings.Count("test", "t"), "time(s) string 'test'")
	fmt.Println("string test has prefix 'tes':", strings.HasPrefix("test", "tes"))
	fmt.Println("string test has suffix 'tes':", strings.HasSuffix("test", "tes"))

	fmt.Println("string letter 't' index:", strings.Index("test", "t"))
	fmt.Println("string letter 'e' index:", strings.Index("test", "e"))
	fmt.Println("string letter 's' index:", strings.Index("test", "s"))
	fmt.Println("string letter 't' index:", strings.Index("test", "t"))
	fmt.Println("'test' to upper case:", strings.ToUpper("test"))
	fmt.Println("'test' to lower case:", strings.ToLower(strings.ToUpper("test")))
	fmt.Println("'test' REPEAT REPEAT:", strings.Repeat("test", 5))

	letters := []string{}
	letters = strings.Split("test", "")
	fmt.Println(letters[0])
	fmt.Println(letters[1])
	fmt.Println(letters[2])
	fmt.Println(letters[3])

	slices := []byte("test")
	fmt.Println(slices)
}
