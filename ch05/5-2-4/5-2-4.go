package main

import (
	"fmt"
	"strings"
)

func main() {
	hdr := []string{"empid", "employee", "address", "hours worked", "hourly rate", "manager"}
	result := csvHdrCol(hdr) //接收回傳值
	fmt.Println("Result:")
	fmt.Println(result)
	fmt.Println()

	hdr2 := []string{"employee", "empid", "hours worked", "address", "manager", "hourly rate"}
	result2 := csvHdrCol(hdr2)
	fmt.Println("Result2:")
	fmt.Println(result2)
	fmt.Println()
}

func csvHdrCol(hdr []string) (csvIdxToCol map[int]string) { //定義傳回值的名稱和型別
	csvIdxToCol = make(map[int]string) //初始化傳回變數
	for i, v := range hdr {
		switch v := strings.ToLower(strings.TrimSpace(v)); v {
		case "employee":
			csvIdxToCol[i] = v
		case "hours worked":
			csvIdxToCol[i] = v
		case "hourly rate":
			csvIdxToCol[i] = v
		}
	}
	return //用naked return 傳回 csvIdxToCol
}
