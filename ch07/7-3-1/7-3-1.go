package main

import "fmt"

type Speaker interface {
	Speak() string
}

type cat struct {
}

func (c cat) Speak() string { //cat實作Speaker介面
	return "Purr Meow"
}

func chatter(s Speaker) { //接收Speaker介面型別的引數
	fmt.Println(s.Speak())
}

func main() {
	c := cat{}
	chatter(c)
}
