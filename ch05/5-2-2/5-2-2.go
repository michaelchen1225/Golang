package main

import (
	"fmt"
	"strings"
)

func main() {
	hdr := []string{"empid", "employee", "address", "hours worked", "hourly rate", "manager"}
	csvHdrCol(hdr)
	hdr2 := []string{"Employee", "Empid", "Hours Worked", "Address", "Manager", "Hourly Rate"}
	csvHdrCol(hdr2)
}

func csvHdrCol(header []string) {
	csvHeadersToColumnIndex := make(map[int]string)
	for i, v := range header {
		//用TrimSpace()把標頭去掉空白和用ToLower()轉成小寫
		//然後比對我們要找的標頭 , 以其索引為鍵放進map
		switch v := strings.ToLower(strings.TrimSpace(v)); v {
		case "employee":
			csvHeadersToColumnIndex[i] = v
		case "hours worked":
			csvHeadersToColumnIndex[i] = v
		case "hourly rate":
			csvHeadersToColumnIndex[i] = v
		}
	}
	fmt.Println(csvHeadersToColumnIndex)
}
