package main

import "fmt"

func hello() {
	fmt.Println("Hello World")
}

func main() {
	fmt.Println("開始")
	go hello() //產生一個goroutine
	fmt.Println("結束")
}
