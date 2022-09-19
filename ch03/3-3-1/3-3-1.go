package main

import (
	"fmt"
	"runtime"
)

func main() {
	var list []int8 // 換成 var list []int8 試試
	for i := 0; i < 10000000; i++ {
		list = append(list, 100)
	}
	//印出記憶體用量
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("TotalAlloc (Heap) = %v MiB\n", m.TotalAlloc/1024/1024)
}
