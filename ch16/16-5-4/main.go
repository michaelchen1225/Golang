package main

import "fmt"

func doSomething() (chan int, chan bool) {
	//建立通道
	in, out := make(chan int), make(chan bool)
	//以匿名函式啟動goroutine
	go func() {
		for i := range in {
			fmt.Println(i)
		}
		out <- true //通知作業結束
	}()
	return in, out //傳回通道
}

func main() {
	in, out := doSomething() //從函式取得通道
	in <- 1
	in <- 2
	in <- 3
	close(in)
	<-out //等待goroutine結束
}
