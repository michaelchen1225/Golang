package main

import (
	"os"
	"time"
)

func main() {
	//建立開啟檔案
	f, err := os.OpenFile("junk.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.Write([]byte(time.Now().String() + "\n"))
}
