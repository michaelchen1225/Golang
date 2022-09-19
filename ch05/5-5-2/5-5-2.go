package main

import "fmt"

type calc func(int, int) string //自訂函式型別

func main() {
	calculator(add, 5, 6) //把其他函式當成引數
}

//add函式會符合自訂的 calc 型別
func add(i, j int) string {
	result := i + j
	return fmt.Sprintf("%d + %d = %d", i, j, result)
}

//接收自訂函式型別參數f
//效果等同於 f fun(int, int) string
func calculator(f calc, i, j int) {
	fmt.Println(f(i, j))
}
