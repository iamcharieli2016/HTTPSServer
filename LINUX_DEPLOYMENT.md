# Linux服务器部署说明

## 🎉 编译完成

您的Go项目已经成功编译为Linux可执行文件，并创建了完整的部署包。

## 📦 部署包内容

### 部署包文件
- **httpsserver-linux-deploy.tar.gz** (7.1MB) - 完整部署包

### 部署包包含以下文件
```
deploy/
├── httpsserver-linux       # Linux可执行文件 (13MB)
├── database.sql            # 数据库初始化脚本
├── config.env.example      # 配置文件示例
├── start-linux.sh          # Linux启动脚本
├── stop-linux.sh           # Linux停止脚本
├── create-ssl-cert.sh      # SSL证书生成脚本
├── httpsserver.service     # systemd服务文件
└── DEPLOY.md               # 详细部署指南
```

## 🚀 快速部署步骤

### 1. 上传部署包到Linux服务器
```bash
# 上传部署包
scp httpsserver-linux-deploy.tar.gz user@your-server:/tmp/

# 登录服务器
ssh user@your-server

# 解压部署包
cd /opt
sudo tar -xzf /tmp/httpsserver-linux-deploy.tar.gz
sudo mv deploy httpsserver
sudo chown -R $(whoami):$(whoami) httpsserver
cd httpsserver
```

### 2. 安装MySQL（如果未安装）
```bash
# Ubuntu/Debian
sudo apt update && sudo apt install mysql-server

# CentOS/RHEL
sudo yum install mysql-server
```

### 3. 初始化数据库
```bash
# 启动MySQL服务
sudo systemctl start mysql

# 初始化数据库
mysql -u root -p < database.sql
```

### 4. 配置和启动服务
```bash
# 复制配置文件
cp config.env.example config.env

# 编辑配置文件（修改数据库密码等）
nano config.env

# 启动服务
./start-linux.sh
```

### 5. 测试API
```bash
curl -k -X POST https://localhost:18443/service/D_A_BSPDMETA \
  -H 'Content-Type: application/json' \
  -d '{
    "params": {"tableName": ""},
    "serviceId": "D_A_BSPDMETA",
    "showCount": "true",
    "offset": 0,
    "limit": 10,
    "userId": "171635",
    "clientId": "eplat",
    "clientSecret": "eplat20191112144407"
  }'
```

## 🔧 管理命令

```bash
# 启动服务
./start-linux.sh

# 停止服务
./stop-linux.sh

# 查看日志
tail -f logs/server.log

# 生成SSL证书
./create-ssl-cert.sh

# systemd服务管理
sudo systemctl start httpsserver
sudo systemctl stop httpsserver
sudo systemctl status httpsserver
```

## 📖 详细说明

查看 `deploy/DEPLOY.md` 文件获取完整的部署指南，包含：
- 系统要求
- 详细安装步骤
- 生产环境配置
- 故障排除
- 维护命令

## 🎯 快速检查清单

- [ ] 上传部署包到Linux服务器
- [ ] 解压到 `/opt/httpsserver`
- [ ] 安装并配置MySQL
- [ ] 执行数据库初始化脚本
- [ ] 修改配置文件
- [ ] 启动服务
- [ ] 测试API接口
- [ ] 配置防火墙（开放18443端口）
- [ ] 设置systemd服务（生产环境）

## 🚨 注意事项

1. **端口配置**：服务运行在18443端口，确保防火墙允许此端口
2. **SSL证书**：默认生成自签名证书，生产环境建议使用CA颁发的证书
3. **数据库密码**：记得在config.env中设置正确的MySQL密码
4. **权限设置**：确保可执行文件有执行权限
5. **日志监控**：服务日志保存在logs/server.log中

## 🔗 相关链接

- API文档：查看README.md
- 详细部署指南：deploy/DEPLOY.md
- 服务端口：https://localhost:18443 