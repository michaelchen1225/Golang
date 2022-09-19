package main

import "fmt"

func main() {
	username := "Sir_King_Über" //建立含有多重位元組字元的字串

	for i := 0; i < len(username); i++ { //走訪字串中的每一個位元
		fmt.Print(username[i], " ") //印出一個字元加一個空格
	}
}
