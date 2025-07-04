package auth

import (
	"testing"

	"httpsserver/internal/config"
	"httpsserver/internal/model"
)

func TestAuthenticate(t *testing.T) {
	// 创建测试配置
	cfg := &config.Config{}
	cfg.Auth.ClientID = "test_client"
	cfg.Auth.ClientSecret = "test_secret"

	authSvc := New(cfg)

	tests := []struct {
		name     string
		request  model.ServiceRequest
		expected bool
	}{
		{
			name: "Valid credentials",
			request: model.ServiceRequest{
				ClientID:     "test_client",
				ClientSecret: "test_secret",
				UserID:       "test_user",
			},
			expected: true,
		},
		{
			name: "Invalid client ID",
			request: model.ServiceRequest{
				ClientID:     "wrong_client",
				ClientSecret: "test_secret",
				UserID:       "test_user",
			},
			expected: false,
		},
		{
			name: "Invalid client secret",
			request: model.ServiceRequest{
				ClientID:     "test_client",
				ClientSecret: "wrong_secret",
				UserID:       "test_user",
			},
			expected: false,
		},
		{
			name: "Empty credentials",
			request: model.ServiceRequest{
				ClientID:     "",
				ClientSecret: "",
				UserID:       "test_user",
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := authSvc.Authenticate(tt.request)
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestNew(t *testing.T) {
	cfg := &config.Config{}
	cfg.Auth.ClientID = "test_client"
	cfg.Auth.ClientSecret = "test_secret"

	authSvc := New(cfg)

	if authSvc == nil {
		t.Error("Expected non-nil auth service")
	}

	if authSvc.cfg != cfg {
		t.Error("Expected auth service to store config reference")
	}
}
