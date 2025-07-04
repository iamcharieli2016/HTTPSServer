#!/bin/bash

# SSL证书生成脚本
# 使用说明：chmod +x create-ssl-cert.sh && ./create-ssl-cert.sh

echo "🔐 生成SSL证书..."
echo "=================="

# 检查OpenSSL是否存在
if ! command -v openssl &> /dev/null; then
    echo "❌ 错误：OpenSSL未安装"
    echo "请先安装OpenSSL："
    echo "  Ubuntu/Debian: sudo apt install openssl"
    echo "  CentOS/RHEL: sudo yum install openssl"
    exit 1
fi

# 生成自签名证书
echo "🔥 正在生成自签名SSL证书..."
openssl req -x509 -newkey rsa:4096 -keyout server.key -out server.crt -days 365 -nodes -subj "/CN=localhost"

if [ $? -eq 0 ]; then
    echo "✅ SSL证书生成成功"
    echo "📋 证书文件: server.crt"
    echo "🔑 私钥文件: server.key"
    echo "⏰ 有效期: 365天"
    echo ""
    echo "⚠️  注意：这是自签名证书，仅用于测试"
    echo "💡 生产环境建议使用Let's Encrypt等权威CA颁发的证书"
else
    echo "❌ SSL证书生成失败"
    exit 1
fi 