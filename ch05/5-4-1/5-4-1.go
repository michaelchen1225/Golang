package main

import (
	"fmt"
)

func main() {
	f := func() { //f會變成func()型別
		fmt.Println("透過變數呼叫一個匿名函式")
	}
	fmt.Println("匿名函式宣告的下一行")
	f() //透過變數f呼叫匿名函式
}
