# 🔧 数据库连接问题解决方案

## ❌ 问题分析

根据您的错误信息 `dial tcp [::1]:3306: connect: connection refused`，问题原因是：

1. 程序试图连接本地 `[::1]:3306`（IPv6 localhost）
2. 但您的配置是连接远程数据库 `10.220.69.40:3307`
3. 配置文件 `config.env` 没有被正确加载

## ✅ 解决方案

### 1. 重新部署更新的程序包

**新的部署包：`httpsserver-linux-deploy-v2.tar.gz`** 已经修复了配置加载问题。

### 2. 您的配置文件内容

确保您的 `config.env` 文件内容如下：

```env
# 数据库配置
DB_HOST=10.220.69.40
DB_PORT=3307
DB_USER=dcuser
DB_PASSWORD=DataCanvas!23
DB_DATABASE=metadata_db
DB_CHARSET=utf8mb4

# 服务器配置
SERVER_PORT=18443
SSL_CERT_FILE=server.crt
SSL_KEY_FILE=server.key

# 认证配置
CLIENT_ID=eplat
CLIENT_SECRET=eplat2019111214440
```

### 3. 在Linux服务器上的操作步骤

```bash
# 停止当前服务
./stop-linux.sh

# 删除旧文件
rm -f httpsserver-linux

# 上传新的部署包
# scp httpsserver-linux-deploy-v2.tar.gz user@your-server:/tmp/

# 解压新的部署包
tar -xzf /tmp/httpsserver-linux-deploy-v2.tar.gz --strip-components=1

# 确保配置文件正确
cat > config.env << 'EOF'
# 数据库配置
DB_HOST=10.220.69.40
DB_PORT=3307
DB_USER=dcuser
DB_PASSWORD=DataCanvas!23
DB_DATABASE=metadata_db
DB_CHARSET=utf8mb4

# 服务器配置
SERVER_PORT=18443
SSL_CERT_FILE=server.crt
SSL_KEY_FILE=server.key

# 认证配置
CLIENT_ID=eplat
CLIENT_SECRET=eplat2019111214440
EOF

# 给新文件添加执行权限
chmod +x httpsserver-linux
chmod +x *.sh

# 启动服务
./start-linux.sh
```

### 4. 验证修复

启动后，您应该会看到类似这样的输出：

```
🚀 启动HTTPS服务器...
========================
✅ 可执行文件检查通过
📋 加载配置文件...
✅ 配置文件已加载
🔐 SSL证书检查...
✅ SSL证书存在
🔥 服务启动中...
✅ 服务已启动
📋 进程ID: 12345
🌐 访问地址: https://localhost:18443
```

### 5. 检查连接

```bash
# 查看日志
tail -f logs/server.log

# 应该看到类似这样的成功日志：
# 2025/07/03 16:45:00 数据库连接成功
# 2025/07/03 16:45:00 HTTPS服务器启动在端口18443...
```

### 6. 测试API（使用正确的客户端密钥）

```bash
curl -k -X POST https://localhost:18443/service/D_A_BSPDMETA \
  -H 'Content-Type: application/json' \
  -d '{
    "params": {
        "tableName": ""
    },
    "serviceId": "D_A_BSPDMETA",
    "showCount": "true",
    "offset": 0,
    "limit": 10,
    "userId": "171635",
    "clientId": "eplat",
    "clientSecret": "eplat2019111214440"
  }'
```

## 🔍 故障排除

### 如果仍然连接失败：

1. **检查数据库服务器是否可达**：
   ```bash
   telnet 10.220.69.40 3307
   ```

2. **检查防火墙**：
   ```bash
   # 确保3307端口开放
   sudo firewall-cmd --list-ports
   ```

3. **测试数据库连接**：
   ```bash
   mysql -h 10.220.69.40 -P 3307 -u dcuser -p
   ```

4. **检查配置是否正确加载**：
   ```bash
   # 启动时添加调试信息
   echo "数据库配置: $DB_HOST:$DB_PORT" >> logs/debug.log
   ```

### 如果数据库需要初始化：

```bash
# 连接到远程数据库并初始化
mysql -h 10.220.69.40 -P 3307 -u dcuser -p < database.sql
```

## 📋 更新内容

新版本修复了以下问题：
- ✅ 配置文件加载问题
- ✅ 客户端密钥匹配
- ✅ 环境变量正确传递
- ✅ 启动脚本增强

## 🎯 关键改进

1. **配置文件自动加载**：启动脚本现在会自动加载 `config.env` 文件
2. **客户端密钥匹配**：默认值已更新为 `eplat2019111214440`
3. **更好的错误处理**：启动脚本会显示配置加载状态

请使用新的 `httpsserver-linux-deploy-v2.tar.gz` 部署包！ 