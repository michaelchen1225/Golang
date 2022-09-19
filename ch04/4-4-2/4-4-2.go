package main

import (
	"fmt"
	"os"
)

func getPassedArgs() []string {
	//回傳包含使用者參數的切片
	//由於得去掉地一個參數 , 這裡仍得用迴圈
	var args []string
	for i := 1; i < len(os.Args); i++ {
		args = append(args, os.Args[i])
	}
	return args
}

func getLocals(extraLocals []string) []string {
	//回傳包含語系名稱切片

	var locales []string
	//加入兩個預設"語系元素"
	locales = append(locales, "en_US", "fr_FR")
	//加入使用者提供的數量不定參數 (extraLocals 可能為空切片)
	locales = append(locales, extraLocals...)
	return locales
}

func main() {
	locales := getLocals(getPassedArgs())
	fmt.Println("要使用的語系:", locales)
}
