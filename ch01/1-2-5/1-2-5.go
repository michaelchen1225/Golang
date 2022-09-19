package main

import (
	"fmt"
	"time"
)

func main() {
	//以簡式寫法逐一宣告變數
	Debug := false
	logLevel := "info"
	startUpTime := time.Now()
	fmt.Println(Debug, logLevel, startUpTime)
}
