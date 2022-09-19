package main

import (
	"errors"
	"fmt"
)

//接收一個interface{}型別的函數
func doubler(v interface{}) (string, error) {
	if i, ok := v.(int); ok { //嘗試轉成int
		return fmt.Sprint(i * 2), nil
	}
	if s, ok := v.(string); ok { //嘗試轉成string
		return s + s, nil
	}
	//型別不符前面的檢查 , 傳回錯誤
	return "", errors.New("傳入了未支援的值")
}

func main() {
	res, _ := doubler(5)
	fmt.Println("5   :", res)
	res, _ = doubler("yum")
	fmt.Println("yum :", res)
	_, err := doubler(true)
	fmt.Println("true:", err)
}
