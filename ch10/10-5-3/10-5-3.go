package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()            //第一次取得系統時間
	time.Sleep(time.Second * 2)    //等待2秒
	end := time.Now()              //第二次取得系統時間
	duration1 := end.Sub(start)    // 計算兩指之間的長度
	duration2 := time.Since(start) //計算start到time.Now()的時間長度

	fmt.Println("Duration1:", duration1)
	fmt.Println("Duration2:", duration2)
	if duration1 < time.Duration(time.Millisecond*2500) {
		fmt.Println("程式執行時間符合預期")
	} else {
		fmt.Println("程式執行時間超出預期")
	}
}
