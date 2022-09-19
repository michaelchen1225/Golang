package main //宣告套件(package),所有的Go語言檔案都必須以套件宣告起頭
//如果想直接執行此套件,就必須將套件命名為main

import ( //接者,以下程式碼會匯入各種套件
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

var helloList = []string{ //這裡宣告了一個全域變數
	"Hello, world",
	"Καλημέρα κόσμε",
	"こんにちは世界",
	" ایند مالس",
	"Привет, мир",
}

func main() { //這裡宣告的是一個函式,所謂函式就是一段程式碼,呼叫函式時就會執行這段程式
	rand.Seed(time.Now().UnixNano()) //不過go語言的main()函式是Go程式碼的進入點,當你執行go程式時,它會自動呼叫main()
	index := rand.Intn(len(helloList))
	msg, err := hello(index)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)
}

func hello(index int) (string, error) {
	if index < 0 || index > len(helloList)-1 {
		return "", errors.New("out of range: " + strconv.Itoa(index))
	}
	return helloList[index], nil //沒有錯誤就回傳訊息
}
