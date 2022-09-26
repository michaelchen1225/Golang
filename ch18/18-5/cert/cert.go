package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"log"
	"math/big"
	"net"
	"os"
	"time"
)

const (
	clientCertName = `.\client_cert.pem`
	clientKeyName  = `.\client.key`
	serverCertName = `.\server_cert.pem`
	serverKeyName  = `.\server.key`
	host           = "127.0.0.1"
	hostDNS        = "localhost"
)

func main() {
	if err := generateCert(clientCertName, clientKeyName); err != nil {
		log.Println(err)
	}
	log.Println("產生:", clientCertName, clientKeyName)

	if err := generateCert(serverCertName, serverKeyName); err != nil {
		log.Println(err)
	}
	log.Println("產生:", serverCertName, serverKeyName)
}

//產生憑證的函式
func generateCert(certFile, keyFile string) error {
	//產生一個安全的隨機樹當作序號
	serialNumber, err := rand.Int(rand.Reader, big.NewInt(1000))
	if err != nil {
		return err
	}

	now := time.Now() //取得現在時間

	//產生X.509憑證
	ca := &x509.Certificate{
		//持有人資訊
		Subject: pkix.Name{
			CommonName:    "Company",
			Organization:  []string{"Company, INC."},
			Country:       []string{"US"},
			Province:      []string{""},
			Locality:      []string{"San Francisco"},
			StreetAddress: []string{"Golden Gate Bridge"},
			PostalCode:    []string{"94016"},
		},
		//序號
		SerialNumber: serialNumber,
		//簽章加密法
		SignatureAlgorithm: x509.SHA256WithRSA,
		//生效時間(即現在時間)
		NotBefore: now,
		//有效時間(現在開始2年後)
		NotAfter: now.AddDate(2, 0, 0),
		//公鑰用途(可用來簽署憑證, 要用於數位簽證)
		KeyUsage: x509.KeyUsageCertSign | x509.KeyUsageDigitalSignature,
		//公鑰額外用途(客戶端驗證/伺服器驗證)
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		//憑證可當CA使用
		IPAddresses: []net.IP{net.ParseIP(host)},
		//憑證可使用的網址與網域
		DNSNames: []string{hostDNS},
	}

	//用RSA產生私鑰
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}

	//以憑證/私鑰和其公鑰來簽署該憑證
	//x509.CreateCertificate()的第二參數是待簽署的憑證
	//第三個參數則是CA的憑證; 兩者相同代表是自簽署憑證
	//傳回值DER為憑證內容, 是[]byte切片
	DER, err := x509.CreateCertificate(
		rand.Reader, ca, ca, &privateKey.PublicKey, privateKey)
	if err != nil {
		return err
	}

	//將憑證字串轉成PEM(Privacy Enhanced Mail, Base64編碼)格式
	cert := pem.EncodeToMemory(
		&pem.Block{
			Type:  "CERTIFICATE",
			Bytes: DER,
		})

	//將私鑰轉成PEM格式
	key := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
		})

	//將憑證與私鑰(私鑰只限擁有者存取)存為檔案
	//憑證權限設為0777(可由任何人自由存取)
	if err := os.WriteFile(certFile, cert, 0777); err != nil {
		return err
	}
	//私鑰權限設為0600(只有擁有者能讀寫)
	if err := os.WriteFile(keyFile, key, 0600); err != nil {
		return err
	}

	return nil
}
