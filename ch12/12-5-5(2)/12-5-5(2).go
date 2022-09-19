package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("test.txt") //開啟檔案
	if err != nil {
		panic(err)
	}
	defer f.Close()
	content, err := io.ReadAll(f) //讀取檔案的整個內容
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("檔案內容:")
	fmt.Println(string(content))
}
