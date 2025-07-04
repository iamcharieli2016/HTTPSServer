# 多阶段构建 - 构建阶段
FROM golang:1.21-alpine AS builder

# 设置工作目录
WORKDIR /app

# 安装必要的工具
RUN apk add --no-cache git ca-certificates tzdata

# 复制go.mod和go.sum文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用程序
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o httpsserver ./cmd/server

# 运行阶段
FROM alpine:latest

# 安装ca-certificates用于HTTPS请求
RUN apk --no-cache add ca-certificates

# 创建非root用户
RUN addgroup -g 1001 -S httpsserver && \
    adduser -u 1001 -S httpsserver -G httpsserver

# 设置工作目录
WORKDIR /app

# 从构建阶段复制二进制文件
COPY --from=builder /app/httpsserver .

# 复制配置文件
COPY configs/config.env.example ./config.env.example

# 复制数据库脚本
COPY sql/database.sql ./database.sql

# 复制证书生成脚本
COPY scripts/create-ssl-cert.sh ./create-ssl-cert.sh
RUN chmod +x ./create-ssl-cert.sh

# 创建证书目录
RUN mkdir -p /app/certs

# 更改文件所有者
RUN chown -R httpsserver:httpsserver /app

# 切换到非root用户
USER httpsserver

# 暴露端口
EXPOSE 18443

# 健康检查
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
    CMD curl -k https://localhost:18443/health || exit 1

# 启动命令
CMD ["./httpsserver"] 