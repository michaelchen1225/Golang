package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 5
	b := &a

	v := reflect.ValueOf(b).Elem()          //取得b指向的指標
	fmt.Printf("%v %T\n", v.Int(), v.Int()) //轉成整數, 用fmt查看形別和值

	v.SetInt(10) //修改b指向的值
	fmt.Printf("%v %T\n", v.Int(), v.Int())
	fmt.Printf("%v %T\n", v.Interface(), v.Interface())
	fmt.Printf("%v %T\n", *b, *b)
}
