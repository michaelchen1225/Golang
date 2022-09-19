package main

import (
	"errors"
	"fmt"
)

func main() {
	a()
	fmt.Println("這一行現在會印出了")
}

func a() {
	b("good-bye")
	fmt.Println("返回a()")
}

func b(mesg string) {
	defer func() { //用defer來確保匿名函式在panic發生後執行
		//若有panic發生 , 用recover()救回程式
		if r := recover(); r != nil {
			fmt.Println("b()發生錯誤:", r) //印出error
		}
	}()
	if mesg == "good-bye" {
		panic(errors.New("出事情了!!")) //引發panic
	}
	fmt.Print(mesg)
}
