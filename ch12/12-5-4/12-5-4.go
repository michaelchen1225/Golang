package main

import (
	"errors"
	"fmt"
	"os"
)

//檢查檔案是否存在的自訂函式
func checkFile(filename string) {
	finfo, err := os.Stat(filename) //取得檔案描述資訊
	if err != nil {
		if errors.Is(err, os.ErrNotExist) { //若error中包含檔案不存在的錯誤
			fmt.Printf("%v: 檔案不存在!\n\n", filename)
			return //退出函式
		}
	}

	//若檔案正確開啟 , 印出其檔案資訊
	fmt.Printf("檔名: %s\n是目錄: %t\n修改時間: %v\n權限: %v\n大小: %d\n\n",
		finfo.Name(), finfo.IsDir(), finfo.ModTime(), finfo.Mode(), finfo.Size())
}

func main() {
	checkFile("junk.txt")
	checkFile("test.txt")
}
