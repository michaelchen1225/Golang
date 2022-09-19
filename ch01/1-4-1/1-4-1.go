package main

import "fmt"

func main() {
	var total float64 = 2 * 13 //主餐 (2人, 每份 $13)
	fmt.Println("+ 主餐:", total)

	total = total + (4 * 2.25) //飲料(每杯 $2.25, 4杯)
	fmt.Println("+ 飲料:", total)

	total = total - 5 //折抵$5
	fmt.Println("- 折抵:", total)

	tip := total * 0.1 //小費10%
	total = total + tip
	fmt.Println("+ 小費:", total)

	split := total / 2 //分攤額
	fmt.Println("分攤額:", split)

	visitCount := 24 //來店次數(之前來過24次)
	visitCount = visitCount + 1
	remainder := visitCount % 5 //用餘數算符檢查除以5餘數是否為0(是5的倍數), 傳回布林值
	if remainder == 0 {
		fmt.Println("您已獲得來店滿 5 次折價券!")
	}
}
