package main

import (
	"fmt"
	"os"
)

func main() {
	//讀取整個檔案的內容
	content, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("檔案內容:")
	fmt.Println(string(content))
}
