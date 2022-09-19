package main

import (
	"fmt"
	"time"
)

func push(from, to int, out chan int) {
	for i := from; i <= to; i++ {
		out <- i
		time.Sleep(time.Microsecond) //等待1毫秒(模擬資料處理時間)
	}
}

func main() {
	s1 := 0
	ch := make(chan int, 100) //通道緩衝區大小為100

	go push(1, 25, ch)
	go push(26, 50, ch)
	go push(51, 75, ch)
	go push(76, 100, ch)

	for c := 0; c < 100; c++ {
		i := <-ch //從通道讀取數字
		fmt.Println(i)
		s1 += i //累加數字
	}

	fmt.Println("Result:", s1)
}
