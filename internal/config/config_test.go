package config

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	// 测试默认配置
	cfg := Load()

	if cfg.DB.Host != "localhost" {
		t.Errorf("Expected DB host to be 'localhost', got '%s'", cfg.DB.Host)
	}

	if cfg.DB.Port != "3306" {
		t.Errorf("Expected DB port to be '3306', got '%s'", cfg.DB.Port)
	}

	if cfg.Server.Port != "18443" {
		t.Errorf("Expected server port to be '18443', got '%s'", cfg.Server.Port)
	}

	if cfg.Auth.ClientID != "eplat" {
		t.Errorf("Expected client ID to be 'eplat', got '%s'", cfg.Auth.ClientID)
	}
}

func TestLoadWithEnvironmentVariables(t *testing.T) {
	// 设置环境变量
	os.Setenv("DB_HOST", "testhost")
	os.Setenv("DB_PORT", "3307")
	os.Setenv("SERVER_PORT", "8080")
	os.Setenv("CLIENT_ID", "testclient")

	// 清理环境变量
	defer func() {
		os.Unsetenv("DB_HOST")
		os.Unsetenv("DB_PORT")
		os.Unsetenv("SERVER_PORT")
		os.Unsetenv("CLIENT_ID")
	}()

	cfg := Load()

	if cfg.DB.Host != "testhost" {
		t.Errorf("Expected DB host to be 'testhost', got '%s'", cfg.DB.Host)
	}

	if cfg.DB.Port != "3307" {
		t.Errorf("Expected DB port to be '3307', got '%s'", cfg.DB.Port)
	}

	if cfg.Server.Port != "8080" {
		t.Errorf("Expected server port to be '8080', got '%s'", cfg.Server.Port)
	}

	if cfg.Auth.ClientID != "testclient" {
		t.Errorf("Expected client ID to be 'testclient', got '%s'", cfg.Auth.ClientID)
	}
}

func TestGetDSN(t *testing.T) {
	cfg := &Config{}
	cfg.DB.User = "testuser"
	cfg.DB.Password = "testpass"
	cfg.DB.Host = "testhost"
	cfg.DB.Port = "3306"
	cfg.DB.Database = "testdb"
	cfg.DB.Charset = "utf8mb4"

	expected := "testuser:testpass@tcp(testhost:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	actual := cfg.GetDSN()

	if actual != expected {
		t.Errorf("Expected DSN '%s', got '%s'", expected, actual)
	}
}

func TestGetEnv(t *testing.T) {
	// 测试不存在的环境变量
	result := getEnv("NON_EXISTENT_VAR", "default_value")
	if result != "default_value" {
		t.Errorf("Expected 'default_value', got '%s'", result)
	}

	// 测试存在的环境变量
	os.Setenv("TEST_VAR", "test_value")
	defer os.Unsetenv("TEST_VAR")

	result = getEnv("TEST_VAR", "default_value")
	if result != "test_value" {
		t.Errorf("Expected 'test_value', got '%s'", result)
	}
}
