package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("data.csv") //開啟CSV檔案
	if err != nil {
		panic(err)
	}

	reader := csv.NewReader(file) //取得csv.Reader節構
	for {
		record, err := reader.Read() //從csv.Reader讀取一行資料
		if err == io.EOF {           //遇到檔案結尾錯誤, 就離開迴圈
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(record) //印出該行結果
	}
}
