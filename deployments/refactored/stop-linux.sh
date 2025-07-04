#!/bin/bash

# HTTPS服务器 Linux停止脚本
# 使用说明：chmod +x stop-linux.sh && ./stop-linux.sh

echo "🛑 停止HTTPS服务器..."
echo "====================="

# 检查PID文件是否存在
if [ -f "server.pid" ]; then
    PID=$(cat server.pid)
    echo "📋 找到进程ID: $PID"
    
    # 检查进程是否存在
    if ps -p $PID > /dev/null 2>&1; then
        echo "🔥 正在停止服务..."
        kill $PID
        sleep 2
        
        # 检查是否成功停止
        if ps -p $PID > /dev/null 2>&1; then
            echo "⚠️  正常停止失败，强制停止..."
            kill -9 $PID
        fi
        
        echo "✅ 服务已停止"
        rm -f server.pid
    else
        echo "⚠️  进程不存在，可能已经停止"
        rm -f server.pid
    fi
else
    echo "⚠️  PID文件不存在，尝试查找进程..."
    
    # 查找并停止所有相关进程
    PIDS=$(pgrep -f "httpsserver-linux")
    if [ -n "$PIDS" ]; then
        echo "🔥 找到进程: $PIDS"
        echo "🛑 正在停止所有相关进程..."
        pkill -f "httpsserver-linux"
        echo "✅ 所有进程已停止"
    else
        echo "ℹ️  没有找到运行中的服务进程"
    fi
fi

echo "🏁 停止操作完成" 