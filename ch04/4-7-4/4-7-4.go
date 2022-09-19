package main

import (
	"fmt"
)

type name string //name型別

type point struct { //point結構型別
	x int
	y int
}

func (n name) printName() { //name型別方法(值接收器)
	fmt.Println("name:", n)
}

func (p point) setPoint(x, y int) { //point結構方法(指標接收器)
	p.x = x //存取結構欄位
	p.y = y
}

func (p point) getPoint() string { //point結構方法
	return fmt.Sprintf("(%v, %v)", p.x, p.y)
}

func main() {
	var n name = "Golang"
	n.printName() //呼叫 n 的方法 printName()

	a, b := point{}, point{}
	a.setPoint(10, 10)                   //呼叫 a 的setPoint方法
	b.setPoint(10, 5)                    //呼叫 b 的setPoint方法
	fmt.Println("point1:", a.getPoint()) //呼叫 a 的getPoint() 方法
	fmt.Println("point2:", b.getPoint()) //呼叫 b 的getPoint() 方法
}
