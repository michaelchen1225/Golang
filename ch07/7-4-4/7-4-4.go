package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s interface{} = "42"
	fmt.Println(strconv.Atoi(s))
}
