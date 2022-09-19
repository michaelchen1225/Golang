package main

import "fmt"

type id string //自訂型別

func getIDs() (id, id, id) {
	//用自訂型別建立變數
	var id1 id
	var id2 id = "1234-5678"
	var id3 id
	id3 = "1234-5678"
	return id1, id2, id3
}

func main() {
	id1, id2, id3 := getIDs()
	//自訂型別互相比對
	fmt.Println("id1 == id2        :", id1 == id2)
	fmt.Println("id2 == id3        :", id2 == id3)
	//轉成字串後跟字串比對
	fmt.Println("id2 == \"1234-5678\":", string(id2) == "1234-5678")
}
