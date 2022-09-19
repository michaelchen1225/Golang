package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello World")
}

func main() {
	fmt.Println("開始")
	go hello()
	time.Sleep(time.Second) //等待一秒
	fmt.Println("結束")
}
