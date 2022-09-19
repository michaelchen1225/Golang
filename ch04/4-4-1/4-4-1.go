package main

import (
	"fmt"
	"os"
)

//將os.Args 參數的每一個元素放進一個切片後回傳
func getPassedArgs(minArgs int) []string {
	if len(os.Args) < minArgs { //如果使用者提供的參數不足就結束
		fmt.Printf("至少需要輸入 %v 個參數\n", minArgs)
		os.Exit(1) //強制結束程式
	}
	var args []string //建立空切片
	//os.Args的第一個參數是執行的程式名稱, 不是參數
	//所以要從索引1開始走訪:
	for i := 1; i < len(os.Args); i++ {
		args = append(args, os.Args[i])
	}
	return args //回傳切片
}

func findLongest(args []string) string {
	var longest string
	for i := 0; i < len(args); i++ {
		if len(args[i]) > len(longest) {
			longest = args[i]
		}
	}
	return longest
}

func main() {
	if longest := findLongest(getPassedArgs(3)); len(longest) > 0 {
		fmt.Println("傳入的最長單字:", longest)
	} else {
		fmt.Println("發生錯誤")
		os.Exit(1)
	}
}
