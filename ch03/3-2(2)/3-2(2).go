package main

import (
	"fmt"
	"unicode" //使用unicode函式庫
)

func passwordChecker(pw string) bool {
	pwR := []rune(pw) //把密碼轉乘rune型別 , 以便接收 UTF-8 字串
	if len(pwR) < 8 { //若密碼長度不足8 , 等於檢查失敗
		return false
	}
	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSymbol := false

	for _, v := range pwR { //用for range 走訪字串的每個字元, 忽略其索引
		if unicode.IsUpper(v) { //檢查是否有大寫字元
			hasUpper = true
		}
		if unicode.IsLower(v) { //檢查是否有小寫字元
			hasLower = true
		}
		// 是否有數字
		if unicode.IsNumber(v) {
			hasNumber = true
		}

		// 是否有符號
		if unicode.IsPunct(v) || unicode.IsSymbol(v) {
			hasSymbol = true
		}
	}
	return hasUpper && hasLower && hasNumber && hasSymbol
}

func main() {
	if passwordChecker("") {
		fmt.Println("密碼格式良好")
	} else {
		fmt.Println("密碼格式不正確")
	}

	if passwordChecker("This!I5A") {
		fmt.Println("密碼格式良好")
	} else {
		fmt.Println("密碼格式不正確")
	}
}
