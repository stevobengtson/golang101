package main

import (
	"fmt"
	"os"
)

func myReadFile(filename string) ([]byte, error) {
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("my error reading file: %w", err)
	}

	return fileContents, nil
}

func main() {
	var fileContents []byte

	fmt.Println("Hello, World!")

	// Read in file and print out the contents
	fileContents, err := myReadFile("test.txt") // ([]byte, error)
	if err != nil {
		panic(err)
	}

	// Print out as a string
	fmt.Println(string(fileContents))
}
