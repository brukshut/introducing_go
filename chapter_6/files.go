package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// slice of bytes, our string
	words := []byte("oh my gosh newfile1\n")
	err := ioutil.WriteFile("newfile1", words, 0755)
	check(err)

	f, err := os.Create("newfile2")
	check(err)
	defer f.Close()

	n1, err := f.WriteString("holy moly newfile2\n")
	fmt.Printf("wrote %d bytes\n", n1)

	f.Sync()
}
