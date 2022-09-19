package main

import "fmt"

var level = "pkg" //套件範圍變數

func main() {
	fmt.Println("Main start  :", level) //main()層級
	if true {
		fmt.Println("Block start :", level) //main底下的if層級
		funcA()
	}
}

func funcA() {
	fmt.Println("funcA start :", level) //funcA()函式層級
}
