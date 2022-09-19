package main

import (
	"errors"
	"fmt"
)

func validate(input int) error { //會回傳error的檢查函式
	if input < 0 {
		return errors.New("輸入值不得為負")
	} else if input > 100 {
		return errors.New("輸入值不得超過 100")
	} else if input%7 == 0 {
		return errors.New("輸入值不得為 7 的倍數")
	} else {
		return nil //檢查都通過時回傳nil回傳nil
	}
}

func main() {
	input := 21
	if err := validate(input); err != nil { //接收error並檢查是否有錯誤
		fmt.Println(err)
	} else if input%2 == 0 {
		fmt.Println(input, "是偶數")
	} else {
		fmt.Println(input, "是奇數")
	}
}
