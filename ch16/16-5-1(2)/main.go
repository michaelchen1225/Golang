package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(in, out chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for i := range in { //讀取通道直到它被關閉
		sum += i
		time.Sleep(time.Millisecond) //模擬資料處理時間
	}
	out <- sum
}

//負責彙整woker()計算結果的goroutine
//注意sum()接收的in和out不會跟worker()一樣
func sum(in, out chan int) {
	sum := 0
	for i := range in { //讀取通道直到它被關閉
		sum += i
	}
	out <- sum
}

func work(workers, from, to int) int {
	wg := &sync.WaitGroup{}
	wg.Add(workers)

	in := make(chan int, (to-from)+1)
	out := make(chan int, workers)
	res := make(chan int, 1)

	for i := 0; i < workers; i++ {
		go worker(in, out, wg) //產生指定數量的worker()
	}
	go sum(out, res) //執行sum()

	for i := from; i <= to; i++ {
		in <- i //提供資料給各worker()
	}
	close(in)  //關閉in(通知所有woker停止讀取)
	wg.Wait()  //等待所有worker結束
	close(out) //關閉out(通知sum停止讀值)

	return <-res //讀取並傳回最終加總
}

func main() {
	res := work(4, 1, 100) //建立4個worker, 計算1~100加總
	fmt.Println(res)
}
