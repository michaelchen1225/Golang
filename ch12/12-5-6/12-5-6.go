package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("test.txt") //開啟檔案取得os.File結構
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Println("檔案內容:")
	//建立一個bufio.Reader結構, 緩衝區大小10
	reader := bufio.NewReaderSize(file, 10)
	for {
		//讀取reader直到碰到換行符號為止
		line, err := reader.ReadString('\n')
		fmt.Print(line)
		if err == io.EOF { //若讀取到檔案結尾就結束
			break
		}
	}
}
