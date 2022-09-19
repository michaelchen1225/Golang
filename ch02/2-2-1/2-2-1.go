package main

import (
	"fmt"
)

func main() {
	input := 5 //定義一個型別為int的變數

	if input%2 == 0 { //檢查除以2的餘數是否為0
		fmt.Println(input, "是偶數")
	}
	if input%2 == 1 {
		fmt.Println(input, "是奇數")
	}
}
