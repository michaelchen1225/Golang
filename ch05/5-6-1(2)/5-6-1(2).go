package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("換我結束!!")
	}() //要記得最後的小括號!!

	fmt.Println("main()開始")
	fmt.Println("main()結束")
}
