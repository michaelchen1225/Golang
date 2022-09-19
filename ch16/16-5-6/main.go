package main

import (
	"fmt"
	"sync"
)

type Workers struct { //Worker結構
	in, out   chan int   //輸入和輸出通道
	workerNum int        //最大goroutine數
	mtx       sync.Mutex //互斥鎖(不需用指標)
}

//初始化Workers結構, 建立通道及互斥鎖 , 啟動goroutine
func (w *Workers) init(maxWorkers, maxData int) {
	//建立通道
	w.in, w.out = make(chan int, maxData), make(chan int)
	//建立互斥鎖
	w.mtx = sync.Mutex{}
	for i := 0; i < maxWorkers; i++ {
		w.mtx.Lock()
		w.workerNum++ //紀錄
		w.mtx.Unlock()
		go w.readThem() //啟動goroutine
	}
}

//輸入資料
func (w *Workers) addData(data int) {
	w.in <- data
}

//讀出資料
func (w *Workers) readThem() {
	sum := 0
	for i := range w.in { //讀取通道in直到關閉和無值
		sum += i
	}
	w.out <- sum //將自己部份的加總值傳給通道out

	//任務結束 , 減少goroutine的紀錄數量
	w.mtx.Lock()
	w.workerNum--
	w.mtx.Unlock()
	if w.workerNum <= 0 { //減到0時關閉通道out
		close(w.out)
	}
}

//取得結果
func (w *Workers) gatherResult() int {
	close(w.in) //關閉通道in
	total := 0
	for i := range w.out { //讀取通道out直到關閉和無值
		total += i
	}
	return total
}

func main() {
	maxWorkers := 10
	maxData := 100
	workers := Workers{}              //建立Workers節構
	workers.init(maxWorkers, maxData) //初始化Workers

	for i := 1; i <= maxData; i++ {
		workers.addData(i) //新增資料
	}
	res := workers.gatherResult() //取得結果
	fmt.Println(res)
}
