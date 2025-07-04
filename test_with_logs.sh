#!/bin/bash

echo "🚀 启动HTTPS服务并显示详细日志"
echo "=================================="

# 启动服务在后台
go run main.go config.go &
SERVER_PID=$!

echo "✅ 服务启动中，PID: $SERVER_PID"
echo "⏳ 等待服务完全启动..."
sleep 3

echo ""
echo "📡 发送第一个API测试请求（查询所有数据）"
echo "============================================"

curl -k -X POST https://localhost:18443/service/D_A_BSPDMETA \
  -H 'Content-Type: application/json' \
  -d '{
    "params": {
        "columnComment": "",
        "columnType": "",
        "columnName": "",
        "tableComment": "",
        "tableName": "",
        "tableSchema": "",
        "dbType": ""
    },
    "serviceId": "D_A_BSPDMETA",
    "showCount": "true",
    "offset": 0,
    "limit": 3,
    "userId": "171635",
    "clientId": "eplat",
    "clientSecret": "eplat20191112144407"
}' | jq '.'

echo ""
echo ""
echo "📡 发送第二个API测试请求（搜索users表）"
echo "========================================="

curl -k -X POST https://localhost:18443/service/D_A_BSPDMETA \
  -H 'Content-Type: application/json' \
  -d '{
    "params": {
        "columnComment": "",
        "columnType": "",
        "columnName": "",
        "tableComment": "",
        "tableName": "users",
        "tableSchema": "",
        "dbType": ""
    },
    "serviceId": "D_A_BSPDMETA",
    "showCount": "true",
    "offset": 0,
    "limit": 10,
    "userId": "171635",
    "clientId": "eplat",
    "clientSecret": "eplat20191112144407"
}' | jq '.'

echo ""
echo ""
echo "📡 发送第三个API测试请求（认证失败测试）"
echo "========================================"

curl -k -X POST https://localhost:18443/service/D_A_BSPDMETA \
  -H 'Content-Type: application/json' \
  -d '{
    "params": {
        "columnComment": "",
        "columnType": "",
        "columnName": "",
        "tableComment": "",
        "tableName": "",
        "tableSchema": "",
        "dbType": ""
    },
    "serviceId": "D_A_BSPDMETA",
    "showCount": "true",
    "offset": 0,
    "limit": 5,
    "userId": "171635",
    "clientId": "wrong_client",
    "clientSecret": "wrong_secret"
}' | jq '.'

echo ""
echo ""
echo "🛑 测试完成，停止服务..."
kill $SERVER_PID
echo "✅ 服务已停止"
echo ""
echo "📝 说明：由于Go服务的日志输出到了后台进程，在这个脚本中看不到。"
echo "💡 要查看详细日志，请运行：go run main.go config.go"
echo "   然后在另一个终端窗口发送curl请求" 