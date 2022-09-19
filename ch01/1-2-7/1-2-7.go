package main

import (
	"fmt"
	"time"
)

func getConfig() (bool, string, time.Time) {
	return false, "info", time.Now()
}
func main() {
	//只提供型別
	var start, middle, end float32
	fmt.Println(start, middle, end)
	//初始值的型別各異
	var name, left, right, top, bottom = "one", 1, 1.5, 2, 2.5
	fmt.Println(name, left, right, top, bottom)
	//對函式一樣適用
	var Debug, logLevel, startUpTime = getConfig()
	fmt.Println(Debug, logLevel, startUpTime)
}
