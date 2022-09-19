package main

import "fmt"

type playerName interface {
	giveName(n string)
}
type player struct {
	name string
}

func (p *player) giveName(n string) { //使用指標接收器
	p.name = n
}

func goat(p playerName) {
	fmt.Println(p, "is the greatest of all time.")
}

func main() {
	p1 := &player{} //要宣告為指標
	p1.giveName("Michael Jordan")
	goat(p1)
	//印出 &{Michael Jordan} is the greatest of all time.
}
