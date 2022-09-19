package main

import "fmt"

type Speaker interface {
	Speak() string
}

type cat struct { //加入欄位
	name string
	age  int
}

func (c *cat) Speak() string { //用指標接收器來指向cat結構變數
	return "Purr Meow"
}

func (c *cat) String() string { //用指標接收器來指向cat結構變數
	return fmt.Sprintf("%v(%v years old)", c.name, c.age)
}

func main() {
	c := cat{name: "Oreo", age: 9} //解決方法 : c := &cat{name: "Oreo", age: 9}
	fmt.Println(c.Speak())
	fmt.Println(c)
}
