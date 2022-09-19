package main

import (
	"fmt"
	"time"
)

var ( //一次宣告多個變數時要用小括號括起來
	Debug      bool      = false      //布林值
	loglevel   string    = "info"     //字串
	starUpTime time.Time = time.Now() //time.Time結構
)

func main() {
	fmt.Println(Debug, loglevel, starUpTime)
}
