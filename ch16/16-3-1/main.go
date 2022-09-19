package main

import (
	"log"
	"sync"
)

//計算累加的函式
func sum(from, to int, wg *sync.WaitGroup, res *int32) {
	for i := from; i <= to; i++ {
		//atomic.AddInt32(res, int32(i))
		*res += int32(i)
	}
	wg.Done()
}

func main() {
	s1 := int32(0)
	wg := &sync.WaitGroup{}

	wg.Add(4) //新增4個goroutine
	go sum(1, 25, wg, &s1)
	go sum(26, 50, wg, &s1)
	go sum(51, 75, wg, &s1)
	go sum(76, 100, wg, &s1)
	wg.Wait() //等待所有goroutine結束

	log.SetFlags(0) //設定log輸出實不帶其他資訊(時間/程式名稱等)
	log.Println(s1)
}
