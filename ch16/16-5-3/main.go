package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(id int, out chan int, done chan bool) {
	for {
		n := rand.Intn(100) //產生0~99的隨機整數
		select {
		case out <- n: //把隨機整數傳給通道out
			fmt.Printf("ID %d 傳送 %d\n", id, n)
		case <-done: //若通道done被關閉而變成可讀
			fmt.Printf("ID %d 結束\n", id)
			return //結束goroutine
		}
		time.Sleep(time.Millisecond) //模擬運算時間
	}
}

func work(workers, from, to int) int {
	out := make(chan int, workers)
	done := make(chan bool)
	defer close(done) //等work()函式結束時關閉done通道(送出取消訊號)

	for i := 0; i < workers; i++ {
		go worker(i, out, done) //建立若干goroutine
	}

	res := 0
	for i := range out {
		if i >= 90 { //若有goroutine回傳值大於90
			res = i
			break //結束work()
		}
	}
	return res //回傳答案給main
}

func main() {
	rand.Seed(time.Now().UnixNano())
	res := work(4, 1, 100)
	fmt.Println("答案:", res)
	time.Sleep(time.Second) //等待1秒, 好看到所有goroutine跑完
}
