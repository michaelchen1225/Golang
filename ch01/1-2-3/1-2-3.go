package main

import (
	"fmt"
	"time"
)

var (
	debug       bool         //省略初始值
	longlevel   = "info"     //省略型別
	startUpTime = time.Now() //省略型別
)

func main() {
	fmt.Println(debug, longlevel, startUpTime)
}
