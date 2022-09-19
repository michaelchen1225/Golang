package main

import (
	"fmt"
)

type point struct {
	x int
	y int
}

func (p *point) setPoint(x, y int) {
	p.x = x
	p.y = y
}

func (p point) getPoint() string {
	return fmt.Sprintf("(%v, %v)", p.x, p.y)
}

func main() {
	a := point{}  //a是結構
	b := &point{} //b是指向結構變數的指標
	a.setPoint(10, 10)
	b.setPoint(10, 5)
	fmt.Println("point1:", a.getPoint())
	fmt.Println("point2:", b.getPoint())
}
