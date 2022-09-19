package main

import (
	"fmt"
	"time"
)

func getConfig() (bool, string, time.Time) {
	return false, "info", time.Now()
}

func main() {
	Debug, loglevel, starUpTime := getConfig()
	fmt.Println(Debug, loglevel, starUpTime)
}
