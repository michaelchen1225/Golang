package main

import "fmt"

func main() {
	var message *string //沒有初始值的指標變數會是nil

	if message == nil {
		fmt.Println("錯誤, 非預期的 nil 值")
		return
	}
	fmt.Println(&message)
}
