package main //定義套件名稱為main

import (
	"fmt" //匯入fmt套件
)

var foo string = "bar" //宣告套件範圍變數
//變數宣告: var  變數名稱  變數型別 = 值
func main() {
	var baz string = "qux" //宣告函示範圍變數
	fmt.Println(foo, baz)  //印出變數
}
