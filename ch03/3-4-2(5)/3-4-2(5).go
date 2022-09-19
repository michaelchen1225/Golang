package main

import "fmt"

func main() {
	username := "Sir_King_Über"

	fmt.Println("Bytes :", len(username))         //取得字串長度 (位元組長度)
	fmt.Println("Runes :", len([]rune(username))) //取得rune集合長度
	//用切片讀取字串的前10個元素 , 立論上剛好到Ü
	fmt.Println(string(username[:10]))
	fmt.Println(string([]rune(username)[:10]))
}
