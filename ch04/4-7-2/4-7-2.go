package main

import "fmt"

type point struct { //具名結構
	x int
	y int
}

func compare() (bool, bool) {
	point1 := struct { //匿名結構(有給初始值)
		x int
		y int
	}{
		10,
		10,
	}
	point2 := struct { //匿名結構(無初始值)
		x int
		y int
	}{}
	point2.x = 10 //設定欄位值
	point2.y = 5
	point3 := point{10, 10}
	return point1 == point2, point1 == point3
}

func main() {
	a, b := compare()
	fmt.Println("point1 == point2:", a)
	fmt.Println("point1 == point3:", b)
}
