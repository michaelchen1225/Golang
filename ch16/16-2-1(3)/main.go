package main

import (
	"fmt"
	"time"
)

func sum(from, to int) int {
	res := 0
	for i := from; i <= to; i++ {
		res += i
	}
	return res
}

func main() {
	var s1, s2 int

	go func() { //執行匿名goroutine, 它會執行s1
		s1 = sum(1, 100)
	}()

	s2 = sum(1, 10) //s2仍在main()內執行

	time.Sleep(time.Second) //等待1秒
	fmt.Println(s1, s2)
}
