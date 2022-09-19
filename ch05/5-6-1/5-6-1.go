package main

import "fmt"

func main() {
	defer done()
	fmt.Println("main()開始")
	fmt.Println("main()結束")
}

func done() {
	fmt.Println("換我結束!!")
}
