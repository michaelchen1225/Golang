package main

import "fmt"

var level = "pkg" //在套件範圍定義level

func main() {
	fmt.Println("Main start  :", level)

	level := 42 //在main範圍定義level
	if true {
		fmt.Println("Block start :", level)
		funcA()
	}
	fmt.Println("Main end    :", level)
}

func funcA() {
	fmt.Println("funcA start :", level)
}
