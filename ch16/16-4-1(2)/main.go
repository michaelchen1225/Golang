package main

import "fmt"

func greet(ch chan string) {
	ch <- "Hello" //對通道傳入訊息
}

func main() {
	ch := make(chan string) //建立無緩衝區的字串通道
	go greet(ch)            //將通道傳給goroutine
	fmt.Println(<-ch)       //從通道接收訊息
}
