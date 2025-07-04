package auth

import (
	"log"

	"httpsserver/internal/config"
	"httpsserver/internal/model"
)

// Service 认证服务
type Service struct {
	cfg *config.Config
}

// New 创建新的认证服务
func New(cfg *config.Config) *Service {
	return &Service{
		cfg: cfg,
	}
}

// Authenticate 验证客户端认证信息
func (s *Service) Authenticate(req model.ServiceRequest) bool {
	// 记录认证尝试
	log.Printf("🔐 认证请求 - ClientID: %s, UserID: %s", req.ClientID, req.UserID)

	// 验证客户端信息
	if req.ClientID != s.cfg.Auth.ClientID || req.ClientSecret != s.cfg.Auth.ClientSecret {
		log.Printf("🚫 客户端认证失败 - 期望ClientID: %s, 实际ClientID: %s", s.cfg.Auth.ClientID, req.ClientID)
		return false
	}

	log.Printf("✅ 客户端认证成功 - UserID: %s", req.UserID)
	return true
}
