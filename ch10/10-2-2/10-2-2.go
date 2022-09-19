package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()    //取得當下時間
	day := now.Weekday() //取得星期幾
	// day = time.Monday
	hour := now.Hour() //取得小時
	// hour = 1
	fmt.Println("Day:", day, "/ hour:", hour)
	if day.String() == "Monday" && (hour >= 0 && hour < 2) {
		fmt.Println("執行全功能測試")
	} else {
		fmt.Println("執行簡易測試")
	}
}
