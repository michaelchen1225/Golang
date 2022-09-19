package main

import (
	"fmt"
	"sync"
)

func readThem(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()     //在結束時對WaitGroup回報
	for i := range ch { //一直讀取ch, 直到他被關閉且無值可取
		fmt.Println(i)
	}
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	ch := make(chan int)
	go readThem(ch, wg)
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	close(ch) //值傳完就關閉通道
	wg.Wait() //等待一個goroutine結束
}
