package main

/*
//以header形式提供給cgo的c程式碼
#include <stdio.h>
#include <stdlib.h>

static void myprint(char* s) {
	printf("%s\n", s);
}
*/
import "C" // 匯入cgo
import (
	"fmt"
	"unsafe"
)

func main() {
	s := "Hello C!" //原始字串

	//把字串轉為CString形別
	cString := C.CString(s)
	//結束時釋放CString指向的記憶體空間'
	defer C.free(unsafe.Pointer(cString))
	//呼叫C程式的函式
	C.myprint(cString)

	//將CString轉回成go的[]byte切片
	b := C.GoBytes(unsafe.Pointer(cString), C.int(len(s)))
	fmt.Println(string(b))
}
