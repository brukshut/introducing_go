package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func ppath(path string, info os.FileInfo, err error) error {
	fmt.Println(path)
	//fmt.Println(info)
	return nil
}

func main() {
	//	filepath.Walk("/usr/local/src", func(path string, info os.FileInfo, err error) error {
	//		fmt.Println(path)
	//fmt.Println(info)
	//		return nil
	//	})

	filepath.Walk("/usr/local/src", ppath)
}
