package main

import (
	"bytes"
	"log"
	"testing"
)

func Test_Main(t *testing.T) {
	for i := 0; i < 10000; i++ {
		var s bytes.Buffer //建立一個Buffer結構(符合io.Writer介面)
		log.SetOutput(&s)  //將log輸出結果寫道s
		log.SetFlags(0)    //設定log輸出時不帶其他資訊
		main()

		//只要log輸出內容不是字串5050\n就算失敗
		if s.String() != "5050\n" {
			t.Fail()
		}
	}
}
