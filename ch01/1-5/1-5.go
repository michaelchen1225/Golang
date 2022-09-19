package main

import (
	"fmt"
	"time"
)

func main() {
	var count int                          //整數
	fmt.Printf("Count    : %#v \n", count) //count的值會帶入%#v

	var discount float64 //64位元浮點
	fmt.Printf("Discount : %#v \n", discount)

	var debug bool //布林值
	fmt.Printf("Debug    : %#v \n", debug)

	var message string //字串
	fmt.Printf("Message  : %#v \n", message)

	var emails []string //字串切片
	fmt.Printf("Emails   : %#v \n", emails)

	var startTime time.Time //time.Time結構
	fmt.Printf("Start    : %#v \n", startTime)
}
