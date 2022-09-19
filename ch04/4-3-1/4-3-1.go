package main

import "fmt"

func defineArray() [10]int { //函式的回傳值微陣列
	var arr [10]int
	return arr
}

func main() {
	fmt.Printf("%#v\n", defineArray()) //印出陣列的型別與值
}
