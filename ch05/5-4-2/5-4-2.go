package main

import (
	"fmt"
)

func main() {
	i := 0
	increment := func() int {
		i++
		return i
	}

	fmt.Println(increment())
	fmt.Println(increment())
	i += 10
	fmt.Println(increment())
}
