package main

import (
	"fmt"
	"strings"
)

func readThem(in chan string, done chan bool) {
	for i := range in { //讀取通道in直到它關閉
		fmt.Println(strings.ToUpper(i)) //把字母轉大寫印出
	}
	done <- true //傳送結束訊號(值不重要)
}

func main() {
	strs := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	workers := 4
	in := make(chan string, len(strs))
	done := make(chan bool, workers)

	for i := 0; i < workers; i++ {
		go readThem(in, done) //建立4個goroutine
	}

	for _, s := range strs { //將字母傳入通道in
		in <- s
	}
	close(in) //關閉通道in

	for i := 0; i < workers; i++ {
		<-done //等待收到4個停止訊號
	}
}
