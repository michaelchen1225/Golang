package main

import "fmt"

type cat struct {
	name string
}

func main() {
	i := 99                //整數形別
	b := false             //布林值
	str := "test"          //字串
	c := cat{name: "oreo"} //cat結構型別
	printDetails(i, b, str, c)
}

func printDetails(data ...interface{}) { //接收不定數量的空介面參數
	for _, i := range data {
		fmt.Printf("%v, %T\n", i, i) //印出值和型別
	}
}
