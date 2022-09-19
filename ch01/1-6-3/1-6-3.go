package main

import (
	"fmt"
	"time"
)

func main() {
	var count1 *int
	count2 := new(int)
	countTemp := 5
	count3 := &countTemp
	t := &time.Time{}

	if count1 != nil {
		fmt.Printf("count1: %#v\n", *count1) //用*取得指標的值
	}
	if count2 != nil {
		fmt.Printf("count2: %#v\n", *count2)
	}
	if count3 != nil {
		fmt.Printf("count3: %#v\n", *count3)
	}
	if t != nil {
		fmt.Printf("time  : %#v\n", *t)
		//存取t(time.Time結構)自身的函示時不需要寫成*t來解除參照
		fmt.Printf("time  : %#v\n", t.String())
	}
}
