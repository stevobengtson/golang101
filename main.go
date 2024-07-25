package main

import (
	"fmt"
	"os"

	"github.com/stevobengtson/golang101/pkg"
)

func myReadFile(filename string) ([]byte, error) {
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("my error reading file: %w", err)
	}

	return fileContents, nil
}

func doSomething(m pkg.MyInterface) {
	m.Increment()
	fmt.Println("mStruct.myInt:", m.GetInt())
}

func main() {
	var fileContents []byte // nil
	var count int64         // 0
	var myFloat float64     // 0.00
	var strings []string    // nil

	var myInt int = 10

	// Factory
	//mStruct := makeSomething.GetMyStructInstance() // MyInterface
	mStruct := pkg.MySecondStruct{MyString: "Hello, World!"}
	mStruct.SetInt(10)

	doSomething(&mStruct)

	var anotherStructInstance pkg.MyStruct
	anotherStructInstance.SetInt(100)
	anotherStructInstance.MyString = "Hello, World!"

	var myStructInstance pkg.MySecondStruct
	myStructInstance.SetInt(myInt)
	myStructInstance.Increment()

	fmt.Println("myStructInstance.myInt:", myStructInstance.GetInt())
	fmt.Println("myStructInstance.myString:", myStructInstance.MyString)

	fmt.Println("myInt: ", &myInt)
	fmt.Println("myStructInstance.myInt: ", myStructInstance.GetInt())

	myInt = 20
	fmt.Println("myStructInstance.myInt:", myStructInstance.GetInt())

	myStructInstance.SetInt(30)
	fmt.Println("myStructInstance.myInt:", myStructInstance.GetInt())
	fmt.Println("myInt: ", myInt)

	strings = make([]string, 10)

	strings[0] = "Hello"
	strings[1] = "Hello"
	strings[2] = "Hello"
	strings[3] = "Hello"
	strings[4] = "Hello"
	strings[5] = "Hello"

	count = 20
	myFloat = 3.14
	count = int64(myFloat * 100)
	fmt.Println("count:", count)

	fmt.Println("Hello, World!")

	// Read in file and print out the contents
	fileContents, err := myReadFile("test.txt") // ([]byte, error)
	if err != nil {
		panic(err)
	}

	// Print out as a string
	fmt.Println(string(fileContents))
}
