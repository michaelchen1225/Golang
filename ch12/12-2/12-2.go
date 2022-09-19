package main

import (
	"flag"
	"fmt"
)

func main() {
	//定義一個旗標 -value, 接收整數, 預設值為 -1
	v := flag.Int("value", -1, "Needs a value for the flag.")
	flag.Parse()
	fmt.Println(*v)
}
