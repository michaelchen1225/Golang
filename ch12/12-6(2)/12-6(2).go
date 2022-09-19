package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

//自訂error
var ErrWorkingFileNotFound = errors.New("查無工作檔案")

func main() {
	workFileName := "note.txt"
	backupFileName := "backup.txt"
	err := writeBackup(workFileName, backupFileName)
	if err != nil {
		panic(err)
	}
}

//備分檔案的函式
func writeBackup(work, backup string) error {
	workFile, err := os.Open(work) //開啟工作檔
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return ErrWorkingFileNotFound //查無工作檔則傳回自訂error
		}
		return err
	}
	defer workFile.Close() //備份結束後關閉工作檔

	//開啟備份檔, 沒有就新增一個 , 資料附加到結尾
	backFile, err := os.OpenFile(backup, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer backFile.Close() //備分結束後關閉工作檔

	content, err := io.ReadAll(workFile) //讀取工作檔內容
	if err != nil {
		return err
	}

	//把一行日期和工作檔內容寫入備分檔
	_, err = backFile.WriteString(fmt.Sprintf("[%v]\n%v", time.Now().String(), string(content)))
	if err != nil {
		return err
	}

	return nil
}
