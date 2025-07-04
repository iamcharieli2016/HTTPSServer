package utils

import (
	"log"
	"net/http"
	"os"
)

// GenerateSelfSignedCert 生成自签名证书
func GenerateSelfSignedCert(certFile, keyFile string) {
	// 检查证书文件是否存在
	if FileExists(certFile) && FileExists(keyFile) {
		return
	}

	log.Println("生成自签名SSL证书...")

	// 这里简化处理，实际生产环境应该使用真实证书
	// 您可以使用openssl命令生成：
	// openssl req -x509 -newkey rsa:4096 -keyout server.key -out server.crt -days 365 -nodes -subj "/CN=localhost"

	log.Printf("请手动生成SSL证书文件:")
	log.Printf("openssl req -x509 -newkey rsa:4096 -keyout %s -out %s -days 365 -nodes -subj \"/CN=localhost\"", keyFile, certFile)
	log.Fatal("SSL证书文件不存在，请先生成证书文件")
}

// FileExists 检查文件是否存在
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

// FileExistsOld 检查文件是否存在（旧版本方法）
func FileExistsOld(filename string) bool {
	_, err := http.Dir(".").Open(filename)
	return err == nil
}
