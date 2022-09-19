package main

import (
	"fmt"
)

type name string //自訂型別

type location struct { //結構型別
	x int
	y int
}

type size struct { //結構型別
	width  int
	height int
}

type dot struct {
	name     //內嵌自訂型別
	location //內嵌結構
	size     //內嵌結構
}

func getDots() []dot {
	var dot1 dot

	dot2 := dot{}
	dot2.name = "A" //存取內嵌的自訂型別
	dot2.x = 5      //內嵌結構型別的欄位被提升了
	dot2.y = 6
	dot2.width = 10
	dot2.height = 20

	dot3 := dot{
		name: "B",
		location: location{ //透過嵌入型別賦予初始值
			x: 13,
			y: 27,
		},
		size: size{
			width:  5,
			height: 7,
		},
	}

	dot4 := dot{}
	dot4.name = "C"
	dot4.location.x = 101
	dot4.location.y = 209
	dot4.size.width = 87
	dot4.size.height = 43

	return []dot{dot1, dot2, dot3, dot4}
}

func main() {
	dots := getDots()
	for i := 0; i < len(dots); i++ {
		fmt.Printf("dot%v: %#v\n", i+1, dots[i])
	}

}
