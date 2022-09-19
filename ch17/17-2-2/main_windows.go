//檔名_windows.go,只會在Windows系統編譯
package main

import "fmt"

func main() {
	fmt.Println("Hello Windows!")
	fmt.Println(greetings()) //呼叫來自main套件,位於其他檔案的函式
}
