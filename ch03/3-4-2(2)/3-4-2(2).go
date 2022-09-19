package main

import "fmt"

func main() {
	username := "Sir_King_Über"

	for i := 0; i < len(username); i++ {
		fmt.Print(string(username[i]), " ") //用string()把字元轉成文字印出來
	}
}
