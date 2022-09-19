package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//建立訊號通道
	sigs := make(chan os.Signal, 1)
	//註冊要通過通道接收的訊號
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	defer cleanUp() //延後執行的清理作業
	fmt.Println("程式執行中 (按下 Ctrl + C 來中斷)")

MainLoop: //一個標籤, 用來代表以下這個無窮for迴圈
	for {
		s := <-sigs
		switch s {
		case syscall.SIGINT:
			fmt.Println("程序中斷:", s)
			break MainLoop
		case syscall.SIGTERM:
			fmt.Println("程序終止:", s)
			break MainLoop
		}
	}

	fmt.Println("程式結果")
}

//模擬程式終止後的清理作業
func cleanUp() {
	fmt.Println("進行清理作業...")
	for i := 0; i <= 10; i++ {
		fmt.Printf("刪除檔案 %v...(僅模擬)\n", i)
		time.Sleep(time.Millisecond * 100)
	}
}
