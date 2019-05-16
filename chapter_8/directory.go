package main

import (
	"fmt"
	"os"
)

func main() {
	dir, err := os.Open("/usr/local/src")
	if err != nil {
		return
	}
	defer dir.Close()

	//fileInfos, err := dir.Readdir(-1)
	fileInfos, err := dir.Readdir(2)
	if err != nil {
		return
	}

	for _, file := range fileInfos {
		fmt.Println(file.Name())
	}
}
