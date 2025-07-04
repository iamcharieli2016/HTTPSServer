package utils

import (
	"os"
	"testing"
)

func TestFileExists(t *testing.T) {
	// 测试存在的文件
	// 创建临时文件
	tmpFile := "test_file.tmp"
	file, err := os.Create(tmpFile)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	file.Close()
	defer os.Remove(tmpFile)

	if !FileExists(tmpFile) {
		t.Error("Expected FileExists to return true for existing file")
	}

	// 测试不存在的文件
	if FileExists("non_existent_file.tmp") {
		t.Error("Expected FileExists to return false for non-existent file")
	}
}

func TestFileExistsOld(t *testing.T) {
	// 测试存在的文件
	// 创建临时文件
	tmpFile := "test_file_old.tmp"
	file, err := os.Create(tmpFile)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	file.Close()
	defer os.Remove(tmpFile)

	if !FileExistsOld(tmpFile) {
		t.Error("Expected FileExistsOld to return true for existing file")
	}

	// 测试不存在的文件
	if FileExistsOld("non_existent_file_old.tmp") {
		t.Error("Expected FileExistsOld to return false for non-existent file")
	}
}

func TestGenerateSelfSignedCert(t *testing.T) {
	// 测试证书已存在的情况
	certFile := "test_cert.crt"
	keyFile := "test_key.key"

	// 创建临时证书文件
	cert, err := os.Create(certFile)
	if err != nil {
		t.Fatalf("Failed to create test cert file: %v", err)
	}
	cert.Close()
	defer os.Remove(certFile)

	key, err := os.Create(keyFile)
	if err != nil {
		t.Fatalf("Failed to create test key file: %v", err)
	}
	key.Close()
	defer os.Remove(keyFile)

	// 当证书文件已存在时，函数应该直接返回，不会panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("GenerateSelfSignedCert panicked when certificates exist: %v", r)
		}
	}()

	GenerateSelfSignedCert(certFile, keyFile)
}

// 为了测试不存在证书文件的情况，我们需要模拟log.Fatal
// 这里提供一个简单的测试框架示例
func TestGenerateSelfSignedCertMissingFiles(t *testing.T) {
	// 这个测试会导致log.Fatal，在实际环境中应该使用依赖注入
	// 或者将log.Fatal替换为返回错误的方式
	t.Skip("Skipping test that would cause log.Fatal")
}
