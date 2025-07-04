# Makefile for HTTPS Server Project

# 应用程序名称
APP_NAME := httpsserver
VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo "v1.0.0")
COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
DATE := $(shell date '+%Y-%m-%d %H:%M:%S')

# 构建目录
BUILD_DIR := build
CMD_DIR := cmd/server

# Go参数
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

# 构建标志
LDFLAGS := -ldflags "-X main.version=$(VERSION) -X main.commit=$(COMMIT) -X main.date=$(DATE)"

# 默认目标
.PHONY: all
all: build

# 构建应用程序
.PHONY: build
build:
	@echo "🔨 构建应用程序..."
	@mkdir -p $(BUILD_DIR)
	@CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build $(LDFLAGS) -o $(BUILD_DIR)/$(APP_NAME) ./$(CMD_DIR)
	@echo "✅ 构建完成: $(BUILD_DIR)/$(APP_NAME)"

# 构建Linux版本
.PHONY: build-linux
build-linux:
	@echo "🔨 构建Linux版本..."
	@mkdir -p $(BUILD_DIR)
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(APP_NAME)-linux ./$(CMD_DIR)
	@echo "✅ Linux版本构建完成: $(BUILD_DIR)/$(APP_NAME)-linux"

# 构建Windows版本
.PHONY: build-windows
build-windows:
	@echo "🔨 构建Windows版本..."
	@mkdir -p $(BUILD_DIR)
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(APP_NAME)-windows.exe ./$(CMD_DIR)
	@echo "✅ Windows版本构建完成: $(BUILD_DIR)/$(APP_NAME)-windows.exe"

# 构建macOS版本
.PHONY: build-darwin
build-darwin:
	@echo "🔨 构建macOS版本..."
	@mkdir -p $(BUILD_DIR)
	@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(APP_NAME)-darwin ./$(CMD_DIR)
	@echo "✅ macOS版本构建完成: $(BUILD_DIR)/$(APP_NAME)-darwin"

# 构建所有平台版本
.PHONY: build-all
build-all: build-linux build-windows build-darwin
	@echo "🎉 所有平台构建完成"

# 运行应用程序
.PHONY: run
run:
	@echo "🚀 运行应用程序..."
	@go run ./$(CMD_DIR)

# 运行测试
.PHONY: test
test:
	@echo "🧪 运行测试..."
	@go test -v ./...

# 运行测试覆盖率
.PHONY: test-coverage
test-coverage:
	@echo "🧪 运行测试覆盖率..."
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "📊 测试覆盖率报告生成: coverage.html"

# 代码格式化
.PHONY: fmt
fmt:
	@echo "🎨 格式化代码..."
	@go fmt ./...

# 代码检查
.PHONY: lint
lint:
	@echo "🔍 代码检查..."
	@golangci-lint run

# 清理构建文件
.PHONY: clean
clean:
	@echo "🧹 清理构建文件..."
	@rm -rf $(BUILD_DIR)
	@rm -f coverage.out coverage.html
	@echo "✅ 清理完成"

# 安装依赖
.PHONY: deps
deps:
	@echo "📦 安装依赖..."
	@go mod tidy
	@go mod download

# 创建SSL证书
.PHONY: cert
cert:
	@echo "🔐 生成SSL证书..."
	@openssl req -x509 -newkey rsa:4096 -keyout server.key -out server.crt -days 365 -nodes -subj "/CN=localhost"
	@echo "✅ SSL证书生成完成"

# 创建部署包
.PHONY: package
package: build-linux
	@echo "📦 创建部署包..."
	@mkdir -p deployments/linux
	@cp $(BUILD_DIR)/$(APP_NAME)-linux deployments/linux/$(APP_NAME)
	@cp configs/config.env.example deployments/linux/
	@cp sql/database.sql deployments/linux/
	@cp scripts/*.sh deployments/linux/
	@cp deployments/systemd/$(APP_NAME).service deployments/linux/
	@tar -czf deployments/$(APP_NAME)-$(VERSION)-linux.tar.gz -C deployments/linux .
	@echo "✅ 部署包创建完成: deployments/$(APP_NAME)-$(VERSION)-linux.tar.gz"

# 显示帮助信息
.PHONY: help
help:
	@echo "可用的命令:"
	@echo "  build         - 构建应用程序"
	@echo "  build-linux   - 构建Linux版本"
	@echo "  build-windows - 构建Windows版本"
	@echo "  build-darwin  - 构建macOS版本"
	@echo "  build-all     - 构建所有平台版本"
	@echo "  run           - 运行应用程序"
	@echo "  test          - 运行测试"
	@echo "  test-coverage - 运行测试覆盖率"
	@echo "  fmt           - 格式化代码"
	@echo "  lint          - 代码检查"
	@echo "  clean         - 清理构建文件"
	@echo "  deps          - 安装依赖"
	@echo "  cert          - 生成SSL证书"
	@echo "  package       - 创建部署包"
	@echo "  help          - 显示帮助信息" 