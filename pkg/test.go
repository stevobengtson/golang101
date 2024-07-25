package pkg

type MyInterface interface {
	Increment()
	SetInt(int)
	GetInt() int
}

type MyStruct struct /*implements MyInterface*/ {
	myInt    int    // 0
	MyString string // nil
}

func (ms *MyStruct) Increment() {
	ms.myInt++
}

func (ms *MyStruct) SetInt(value int) {
	ms.myInt = value
}

func (ms *MyStruct) GetInt() int {
	return ms.myInt
}

type MySecondStruct struct /*implements MyInterface*/ {
	myInt    int    // 0
	MyString string // nil
}

func (ms *MySecondStruct) Increment() {
	ms.myInt *= 2
}

func (ms *MySecondStruct) SetInt(value int) {
	ms.myInt = value
}

func (ms *MySecondStruct) GetInt() int {
	return ms.myInt
}
