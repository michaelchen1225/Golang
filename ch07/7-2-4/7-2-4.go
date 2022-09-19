package main

import "fmt"

type Speaker interface {
	Speak() string
}

type cat struct { //加入欄位
	name string
	age  int
}

func (c cat) Speak() string {
	return "Purr Meow"
}

func (c cat) String() string { //String()方法
	return fmt.Sprintf("%v(%v years old)", c.name, c.age)
}

func main() {
	c := cat{name: "Oreo", age: 9}
	fmt.Println(c.Speak())
	fmt.Println(c) //用fmt套件直接印出cat
}
