package main

import "fmt"

func main() {
	fname := "Joe"
	gpa := 3.75
	hasJob := true
	age := 24
	hourlyWage := 45.53
	//印出字串與浮點數
	fmt.Printf("%s 的 GPA: %f\n", fname, gpa)
	//印出布林值
	fmt.Printf("有工作: %t\n", hasJob)
	//印出整數與一般值
	fmt.Printf("年齡: %d, 時薪: %v\n", age, hourlyWage)
}
