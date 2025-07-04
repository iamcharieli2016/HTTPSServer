#!/bin/bash

echo "🚀 启动HTTPS后端接口服务"
echo "=============================="

# 1. 检查Go环境
if ! command -v go &> /dev/null; then
    echo "❌ 未找到Go环境，请先安装Go 1.21+"
    exit 1
fi

echo "✅ Go环境检查通过"

# 2. 安装依赖
echo "📦 安装项目依赖..."
go mod tidy

# 3. 生成SSL证书（如果不存在）
if [ ! -f "server.crt" ] || [ ! -f "server.key" ]; then
    echo "🔐 生成SSL证书..."
    if command -v openssl &> /dev/null; then
        openssl req -x509 -newkey rsa:4096 -keyout server.key -out server.crt -days 365 -nodes -subj "/CN=localhost"
        echo "✅ SSL证书生成成功"
    else
        echo "❌ 未找到openssl命令，请手动生成SSL证书："
        echo "   openssl req -x509 -newkey rsa:4096 -keyout server.key -out server.crt -days 365 -nodes -subj \"/CN=localhost\""
        exit 1
    fi
else
    echo "✅ SSL证书已存在"
fi

# 4. 检查MySQL连接
echo "🗄️  检查MySQL连接..."
echo "请确保MySQL服务正在运行，并且已经执行了database.sql初始化脚本"
echo "如果还没有初始化数据库，请运行："
echo "   mysql -u root -p < database.sql"
echo ""

# 5. 启动服务
echo "🌟 启动HTTPS服务..."
echo "服务将在 https://localhost:8443 启动"
echo "测试命令："
echo "curl -k -X POST https://localhost:8443/service/D_A_BSPDMETA \\"
echo "  -H 'Content-Type: application/json' \\"
echo "  -d '{\"params\":{\"columnComment\":\"\",\"columnType\":\"\",\"columnName\":\"\",\"tableComment\":\"\",\"tableName\":\"\",\"tableSchema\":\"\",\"dbType\":\"\"},\"serviceId\":\"D_A_BSPDMETA\",\"showCount\":\"true\",\"offset\":0,\"limit\":10,\"userId\":\"171635\",\"clientId\":\"eplat\",\"clientSecret\":\"eplat20191112144407\"}'"
echo ""
echo "按 Ctrl+C 停止服务"
echo "=============================="

go run main.go 