package main

import "fmt"

const GlobalLimit = 100                  //單筆資料筆上限
const MaxCacheSize int = 2 * GlobalLimit //快取最大容量 (單筆上限 x 2)

const (
	CacheKeyBook = "book_" //書本id的前綴字
	CacheKeyCD   = "cd_"   //CD id 的前綴字
)

var cache map[string]string //快去集合

func cacheGet(key string) string { //從快取取出某個鍵的值
	return cache[key]
}

func cacheSet(key, val string) { //將某個鍵和值寫入快取
	if len(cache)+1 >= MaxCacheSize { //如果快取大小已達極限就跳出函式
		return
	}
	cache[key] = val
}

func SetBook(isbn string, name string) { //加入書本資料
	cacheSet(CacheKeyBook+isbn, name) //加入書本前綴字後呼叫cacheSet()
}

func GetBook(isbn string) string { //讀取書本資料
	return cacheGet(CacheKeyBook + isbn) //加上書本前綴字後呼叫cacheGet()
}

func SetCD(sku string, title string) { //加入CD資料
	cacheSet(CacheKeyCD+sku, title) //加上CD前綴字後呼叫cacheSet()
}

func GetCD(sku string) string { //讀取CD資料
	return cacheGet(CacheKeyCD + sku) //加上CD前綴字後呼叫cacheGet()
}

func main() {
	cache = make(map[string]string) //初始化快取
	//在快取寫入資料
	SetBook("1234-5678", "Get Ready To Go")
	SetCD("1234-5678", "Get Ready To Go Audio Book")
	//讀取和印出快取資料
	fmt.Println("Book :", GetBook("1234-5678"))
	fmt.Println("CD   :", GetCD("1234-5678"))
}
