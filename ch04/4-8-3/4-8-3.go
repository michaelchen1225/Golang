package main

import (
	"errors"
	"fmt"
)

func doubler(v interface{}) (string, error) {
	switch t := v.(type) {
	case string: //型別為字串時
		return t + t, nil
	case bool: //型別為布林值時
		if t {
			return "truetrue", nil
		}
		return "falsefalse", nil
	case float32, float64: //型別為浮點數時
		if f, ok := t.(float64); ok {
			return fmt.Sprint(f * 2), nil
		}
		return fmt.Sprint(t.(float32) * 2), nil
	case int: //其餘整數型別檢查
		return fmt.Sprint(t * 2), nil
	case int8:
		return fmt.Sprint(t * 2), nil
	case int16:
		return fmt.Sprint(t * 2), nil
	case int32:
		return fmt.Sprint(t * 2), nil
	case int64:
		return fmt.Sprint(t * 2), nil
	case uint:
		return fmt.Sprint(t * 2), nil
	case uint8:
		return fmt.Sprint(t * 2), nil
	case uint16:
		return fmt.Sprint(t * 2), nil
	case uint32:
		return fmt.Sprint(t * 2), nil
	case uint64:
		return fmt.Sprint(t * 2), nil
	default:
		return "", errors.New("傳入了未支援的值")
	}
}

func main() {
	res, _ := doubler(-5)
	fmt.Println("-5  :", res)
	res, _ = doubler(5)
	fmt.Println("5   :", res)
	res, _ = doubler("yum")
	fmt.Println("yum :", res)
	res, _ = doubler(true)
	fmt.Println("true:", res)
	res, _ = doubler(float32(3.14))
	fmt.Println("3.14:", res)
}
