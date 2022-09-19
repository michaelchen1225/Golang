package main

import (
	"fmt"
)

type Speaker interface {
	Speak() string
}

type person struct {
	name      string
	age       int
	isMarried bool
}

func main() {
	p := person{name: "Cailyn", age: 44, isMarried: false}
	fmt.Println(p.Speak())
	fmt.Println(p)
}

func (p person) String() string { //實作Stringer介面
	return fmt.Sprintf("%v (%v 歲)\n已婚: %v ", p.name, p.age, p.isMarried)
}

func (p person) Speak() string { //實作Speaker介面
	return "各位好, 我的名字是 " + p.name
}
