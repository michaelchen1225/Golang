package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"os"
)

//加密函式
func encrypt(data, key []byte) (resp []byte, err error) {
	//建立區塊加密物件
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return resp, err
	}
	//使用GCM加密模式
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return resp, err
	}
	//產生一個gcm.NonceSize()長度的[]byte切片
	nonce := make([]byte, gcm.NonceSize())
	//用crypto/rand套件產生一個安全隨機數作為nonce
	if _, err := rand.Read(nonce); err != nil {
		return resp, err
	}
	//加密資料, 並將結果附加到nonce尾端(傳回nonce + 密文)
	resp = gcm.Seal(nonce, nonce, data, nil)
	return resp, nil
}

//解密函式
func decrypt(data, key []byte) (resp []byte, err error) {
	//和加密函式一樣, 建立區塊加密物件並使用GCM加密模式
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return resp, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return resp, err
	}
	//分割nonce加密文
	nonce := data[:gcm.NonceSize()]
	encryptedData := data[gcm.NonceSize():]
	//解密資料(dst傳入nil; 若傳入[]byte切片, 傳回結果就是dst + 解密字串)
	resp, err = gcm.Open(nil, nonce, encryptedData, nil)
	if err != nil {
		return resp, fmt.Errorf("解密錯誤: %v", err)
	}
	return resp, nil
}

func main() {
	data := "My secret text"
	fmt.Printf("原始資料: %s\n", data)

	//產生一個16位元長度隨機金鑰
	key := make([]byte, 16)
	if _, err := rand.Read(key); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//加密
	encrypted, err := encrypt([]byte(data), key)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("加密資料: %x\n", string(encrypted))

	//解密
	decrypted, err := decrypt(encrypted, key)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("解密資料: %s\n", string(decrypted))
}
