#!/bin/bash

# HTTPS服务器 Linux启动脚本
# 使用说明：chmod +x start-linux.sh && ./start-linux.sh

echo "🚀 启动HTTPS服务器..."
echo "========================"

# 检查可执行文件是否存在
if [ ! -f "./httpsserver-linux" ]; then
    echo "❌ 错误：找不到可执行文件 httpsserver-linux"
    exit 1
fi

# 检查SSL证书是否存在
if [ ! -f "./server.crt" ] || [ ! -f "./server.key" ]; then
    echo "⚠️  SSL证书不存在，正在生成自签名证书..."
    openssl req -x509 -newkey rsa:4096 -keyout server.key -out server.crt -days 365 -nodes -subj "/CN=localhost"
    echo "✅ SSL证书生成完成"
fi

# 给可执行文件添加执行权限
chmod +x ./httpsserver-linux

# 创建日志目录
mkdir -p logs

# 加载配置文件
if [ -f "config.env" ]; then
    echo "📋 加载配置文件..."
    export $(cat config.env | grep -v '^#' | grep -v '^$' | xargs)
    echo "✅ 配置文件已加载"
else
    echo "⚠️  配置文件config.env不存在，使用默认配置"
fi

# 启动服务
echo "🔥 服务启动中..."
./httpsserver-linux > logs/server.log 2>&1 &

# 获取进程ID
PID=$!
echo $PID > server.pid

echo "✅ 服务已启动"
echo "📋 进程ID: $PID"
echo "🌐 访问地址: https://localhost:18443"
echo "📝 日志文件: logs/server.log"
echo ""
echo "💡 查看日志: tail -f logs/server.log"
echo "🛑 停止服务: kill $PID 或者 ./stop-linux.sh" 