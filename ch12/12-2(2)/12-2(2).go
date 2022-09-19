package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	n := flag.String("name", "", "your first name")
	i := flag.Int("age", -1, "your age")
	b := flag.Bool("married", false, "are you married?")
	flag.Parse()

	if *n == "" { //若名子旗標值為空字串, 但表使用者沒有加上該旗標, 會未給值
		fmt.Println("Name is required.")
		flag.PrintDefaults() //印出所有旗標的預設值
		os.Exit(1)           //結束程式
	}
	fmt.Println("Name: ", *n)
	fmt.Println("Age: ", *i)
	fmt.Println("Married: ", *b)
}
