package main

import (
	"fmt"
	"time"
)

func main() {
	debug, loglevel, starUpTime := false, "info", time.Now()
	fmt.Println(debug, loglevel, starUpTime)
}
