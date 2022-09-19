package main

import (
	"context"
	"fmt"
	"time"
)

func countNumbers(c context.Context, out chan int) {
	i := 0
	for { //無窮迴圈
		select {
		case <-c.Done(): //收到取消信號, 船值給out
			out <- i
			return
		default: //正常情況下, 每100毫秒讓計數器+1
			time.Sleep(time.Millisecond * 100)
			i++
		}
	}
}

func main() {
	out := make(chan int)

	//建立一個空context結構
	c := context.TODO()
	//延伸一個可取消的context並取得取消函式
	cl, cancel := context.WithCancel(c)

	go countNumbers(cl, out) //將context傳給goroutine

	time.Sleep(time.Millisecond * 100 * 5) //等待500毫秒
	cancel()                               //呼叫context提供的取消函式

	fmt.Println(<-out) //印出out內的值
}
