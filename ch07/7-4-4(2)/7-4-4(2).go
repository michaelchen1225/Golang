package main

import "fmt"

type cat struct {
	name string
}

func main() {
	//建立一個空介面切片, 放入不同型別的值
	i := []interface{}{42, "The book club", true, cat{name: "oreo"}}
	typeExample(i)
}

func typeExample(i []interface{}) {
	for _, x := range i {
		switch v := x.(type) { //對切片每個值做型別 switch
		case int:
			fmt.Printf("%v is int\n", v)
		case string:
			fmt.Printf("%v is a string\n", v)
		case bool:
			fmt.Printf("%v is a bool\n", v)
		default:
			fmt.Printf("%T is unknown type\n", v)
		}
	}
}
