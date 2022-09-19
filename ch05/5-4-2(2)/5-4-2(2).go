package main

import (
	"fmt"
)

func main() {
	increment := incrementor() //接收回傳的函式

	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}

func incrementor() func() int {
	i := 0              //定義在匿名函式之父函式內的變數
	return func() int { //閉包 : 匿名函式包住外部變數, 且能夠記得外部變數的值
		i++
		return i
	}
}
