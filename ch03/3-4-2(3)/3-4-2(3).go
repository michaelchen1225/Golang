package main

import "fmt"

func main() {
	username := "Sir_King_Über"
	runes := []rune(username) //將字串轉成rune切片

	for i := 0; i < len(runes); i++ {
		fmt.Print(string(runes[i]), " ") //將rune轉字串印出
	}
}
