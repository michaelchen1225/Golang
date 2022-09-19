package main

import "fmt"

func main() {
	max := 4

	counter := decrement(max) //取得閉包函式
	fmt.Println(counter())    //呼叫閉包函式
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}

func decrement(i int) func() int {

	return func() int {
		//閉包函式會記住父函式的的參數i
		if i > 0 {
			i--
		}
		return i
	}
}
