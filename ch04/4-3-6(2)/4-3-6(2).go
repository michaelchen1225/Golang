package main

import "fmt"

func fillArray(arr [10]int) [10]int { //參數和回傳值都是10個元素的int陣列
	for i := 0; i < len(arr); i++ {
		arr[i] = i + 1 //陣列元素會是1~10
	}
	return arr
}

func opArray(arr [10]int) [10]int {
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * arr[i] //元素平方
	}
	return arr
}

func main() {
	var arr [10]int
	arr = fillArray(arr)
	arr = opArray(arr)
	fmt.Println(arr)
}
