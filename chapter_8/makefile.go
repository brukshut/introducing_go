package main

import "os"

func main() {
	file, err := os.Create("newfile.txt")
	if err != nil {
		// handle error
		return
	}
	defer file.Close()
	file.WriteString("This is a new file.\n")
}
