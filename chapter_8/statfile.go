package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	// error is <nil> if something is wrong
	fmt.Println(err)
	if err != nil {
		return
	}
	fmt.Println(file, err)
	defer file.Close()

	// get the file size
	stat, err := file.Stat()
	if err != nil {
		return
	}
	fmt.Println(stat, err)

	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}

	fmt.Println(stat.Size())

	str := string(bs)
	fmt.Println(str)
}
