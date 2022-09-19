package main

import (
	"fmt"
	"sync"
)

func sum(from, to int, wg *sync.WaitGroup, res *int) {
	*res = 0
	for i := from; i <= to; i++ {
		*res += i
	}
	if wg != nil {
		wg.Done() //回報goroutine結束
	}
}

func main() {
	var s1, s2 int

	wg := &sync.WaitGroup{} //建立WaitGroup
	wg.Add(1)               //要等待一個goroutine
	go sum(1, 100, wg, &s1) //以goroutine形式執行sum()
	sum(1, 10, nil, &s2)    //以正常方式執行sum()
	wg.Wait()               //等待goroutine結束

	fmt.Println(s1, s2)
}
