package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	name string
	age  int
}

type Employee struct {
	name string
	age  int
}

func main() {
	a := User{"John", 42}
	fmt.Printf("a = %#v\n", a)

	//把a從User轉成Employee結構型別
	b := *(*Employee)(unsafe.Pointer(&a))
	fmt.Printf("b = %#v\n", b)
}
