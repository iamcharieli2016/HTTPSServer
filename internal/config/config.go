package config

import (
	"os"
)

// Config 应用配置结构体
type Config struct {
	// 数据库配置
	DB struct {
		Host     string
		Port     string
		User     string
		Password string
		Database string
		Charset  string
	}

	// 服务器配置
	Server struct {
		Port     string
		CertFile string
		KeyFile  string
	}

	// 认证配置
	Auth struct {
		ClientID     string
		ClientSecret string
	}
}

// Load 加载配置
func Load() *Config {
	config := &Config{}

	// 数据库配置（支持环境变量）
	config.DB.Host = getEnv("DB_HOST", "localhost")
	config.DB.Port = getEnv("DB_PORT", "3306")
	config.DB.User = getEnv("DB_USER", "root")
	config.DB.Password = getEnv("DB_PASSWORD", "")
	config.DB.Database = getEnv("DB_DATABASE", "metadata_db")
	config.DB.Charset = getEnv("DB_CHARSET", "utf8mb4")

	// 服务器配置
	config.Server.Port = getEnv("SERVER_PORT", "18443")
	config.Server.CertFile = getEnv("SSL_CERT_FILE", "server.crt")
	config.Server.KeyFile = getEnv("SSL_KEY_FILE", "server.key")

	// 认证配置
	config.Auth.ClientID = getEnv("CLIENT_ID", "eplat")
	config.Auth.ClientSecret = getEnv("CLIENT_SECRET", "eplat2019111214440")

	return config
}

// GetDSN 构建数据库连接字符串
func (c *Config) GetDSN() string {
	return c.DB.User + ":" + c.DB.Password + "@tcp(" + c.DB.Host + ":" + c.DB.Port + ")/" + c.DB.Database + "?charset=" + c.DB.Charset + "&parseTime=True&loc=Local"
}

// getEnv 获取环境变量，如果不存在则返回默认值
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
