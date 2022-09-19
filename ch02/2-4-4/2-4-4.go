package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for {
		r := rand.Intn(8) //產生0~8整數亂數
		if r%3 == 0 {     //若亂數是三的倍數則跳過這輪迴圈
			fmt.Println("略過")
			continue
		} else if r%2 == 0 { //若亂數是偶數則跳出迴圈
			fmt.Println("跳出")
			break
		}
		fmt.Println(r)
	}
}
