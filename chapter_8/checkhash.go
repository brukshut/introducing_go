package main

import (
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

func getHash(filename string) (uint32, error) {
	// open the file
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	// remember to alway close open files
	defer file.Close()

	// create a hasher
	hasher := crc32.NewIEEE()

	var bytes int64
	// copy file into the hasher
	// (dst src) (bytesWritten, err)
	bytes, err = io.Copy(hasher, file)
	fmt.Println(bytes)

	// we don't care about bytes, but we want error
	if err != nil {
		return 0, err
	}
	return hasher.Sum32(), nil
}

func main() {
	hash1, err := getHash("test1.txt")
	if err != nil {
		return
	}

	hash2, err := getHash("test2.txt")
	if err != nil {
		return
	}
	fmt.Println(hash1, hash2)
}
