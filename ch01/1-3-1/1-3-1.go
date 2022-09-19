package main

import (
	"fmt"
)

func main() {
	offset := 5 //注意:這裡不能直接寫offset = 5
	fmt.Println(offset)

	offset = 10 //將offset從5改成10 (只有在更改變數值才可以只用等號)
	fmt.Println(offset)
}
