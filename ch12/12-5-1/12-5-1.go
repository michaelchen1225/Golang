package main

import "os"

func main() {
	f, err := os.Create("test.txt") //建立test.txt
	if err != nil {                 //檢查建立檔案時是否遇到錯誤
		panic(err)
	}
	defer f.Close() //確保在main()結束時關閉檔案
}
