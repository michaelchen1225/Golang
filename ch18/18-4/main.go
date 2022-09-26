package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"fmt"
	"os"
)

func main() {
	data := []byte("My secret document")

	//產生公鑰與私鑰(使用crypto/rand產生亂數)
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//用私鑰產生數位簽章
	signedData := ed25519.Sign(privateKey, data)
	fmt.Printf("數位簽章:\n%x\n", signedData)

	//用公鑰, 資料和數位簽章來驗證簽章是否有效
	verified := ed25519.Verify(publicKey, data, signedData)
	fmt.Println("驗證:", verified)
}
