package main

import (
	"errors"
	"fmt"
)

var ( //自己定義error值
	ErrHourlyRate  = errors.New("無效的時薪")
	ErrHoursWorked = errors.New("無效的一周工時")
)

func main() {
	pay, err := payDay(81, 50)
	if err != nil { //若payDay傳回錯誤就印出
		fmt.Println(err)
	}
	fmt.Println(pay)

	pay, err = payDay(80, 5)
	if err != nil { //同上
		fmt.Println(err)
	}
	fmt.Println(pay)

	pay, err = payDay(80, 50)
	if err != nil { //同上
		fmt.Println(err)
	}
	fmt.Println(pay)
}

func payDay(hoursWorked, hourlyRate int) (int, error) {
	if hourlyRate < 10 || hourlyRate > 75 {
		//用error值得Error()取得內容字串 , 產生新的error值後回傳
		return 0, fmt.Errorf("無效的時薪: %w",ErrHourlyRate)
	}
	if hoursWorked < 0 || hoursWorked > 80 {
		return 0, fmt.Errorf("無效的一周工時: %w", ErrHoursWorked)
	}
	//計算加班費
	if hoursWorked > 40 {
		hoursOver := hoursWorked - 40
		hoursRegular := hoursWorked - hoursOver
		return hoursRegular*hourlyRate + hoursOver*hourlyRate*2, nil
	}
	//沒有錯誤就傳回計算的周薪(含加班費), 錯誤就傳回nil
	return hoursWorked * hourlyRate, nil
}
