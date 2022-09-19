package main

import (
	"fmt"
	"time"
)

func main() {
	switch dayBorn := time.Sunday; { //只有起始賦值敘述
	case dayBorn == time.Sunday || dayBorn == time.Saturday:
		fmt.Println("生日為周末")
	default:
		fmt.Println("生日非周末")
	}
}
