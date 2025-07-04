# Linux服务器部署指南

本指南将帮助您在Linux服务器上部署HTTPS服务器。

## 系统要求

- Linux服务器（Ubuntu 18.04+ / CentOS 7+ / RHEL 7+）
- MySQL 5.7+ 或 MariaDB 10.3+
- OpenSSL
- 网络端口18443可用

## 部署步骤

### 1. 上传文件

将整个`deploy`目录上传到Linux服务器，例如：

```bash
# 使用scp上传（在本地执行）
scp -r deploy/ user@your-server:/opt/httpsserver/

# 或者使用rsync
rsync -avz deploy/ user@your-server:/opt/httpsserver/
```

### 2. 安装MySQL

#### Ubuntu/Debian:
```bash
sudo apt update
sudo apt install mysql-server mysql-client
```

#### CentOS/RHEL:
```bash
sudo yum install mysql-server mysql
# 或者使用dnf（CentOS 8+）
sudo dnf install mysql-server mysql
```

#### 启动MySQL服务:
```bash
sudo systemctl enable mysql
sudo systemctl start mysql
```

### 3. 初始化数据库

```bash
cd /opt/httpsserver
mysql -u root -p < database.sql
```

### 4. 配置环境变量

复制配置文件并修改：

```bash
cp config.env.example config.env
# 编辑配置文件
nano config.env
```

修改以下配置：
```env
# 数据库配置
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_mysql_password
DB_DATABASE=metadata_db
DB_CHARSET=utf8mb4

# 服务器配置
SERVER_PORT=18443
SSL_CERT_FILE=server.crt
SSL_KEY_FILE=server.key

# 认证配置
CLIENT_ID=eplat
CLIENT_SECRET=eplat20191112144407
```

### 5. 设置权限

```bash
# 给脚本添加执行权限
chmod +x httpsserver-linux
chmod +x start-linux.sh
chmod +x stop-linux.sh

# 创建日志目录
mkdir -p logs
```

### 6. 启动服务

#### 方式1：使用启动脚本（推荐测试）

```bash
./start-linux.sh
```

查看日志：
```bash
tail -f logs/server.log
```

停止服务：
```bash
./stop-linux.sh
```

#### 方式2：使用systemd服务（推荐生产环境）

```bash
# 创建专用用户
sudo useradd -r -s /bin/false httpsserver

# 设置目录权限
sudo chown -R httpsserver:httpsserver /opt/httpsserver

# 复制服务文件
sudo cp httpsserver.service /etc/systemd/system/

# 重载systemd配置
sudo systemctl daemon-reload

# 启动服务
sudo systemctl enable httpsserver
sudo systemctl start httpsserver

# 查看状态
sudo systemctl status httpsserver

# 查看日志
sudo journalctl -u httpsserver -f
```

### 7. 防火墙配置

#### Ubuntu (UFW):
```bash
sudo ufw allow 18443/tcp
sudo ufw reload
```

#### CentOS/RHEL (firewalld):
```bash
sudo firewall-cmd --add-port=18443/tcp --permanent
sudo firewall-cmd --reload
```

### 8. 测试API

```bash
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
    "limit": 10,
    "userId": "171635",
    "clientId": "eplat",
    "clientSecret": "eplat20191112144407"
}'
```

## 生产环境建议

### 1. 使用真实SSL证书

```bash
# 使用Let's Encrypt免费证书
sudo certbot certonly --standalone -d your-domain.com
```

然后在配置文件中指定证书路径：
```env
SSL_CERT_FILE=/etc/letsencrypt/live/your-domain.com/fullchain.pem
SSL_KEY_FILE=/etc/letsencrypt/live/your-domain.com/privkey.pem
```

### 2. 配置反向代理

使用Nginx作为反向代理：

```nginx
server {
    listen 80;
    server_name your-domain.com;
    return 301 https://$server_name$request_uri;
}

server {
    listen 443 ssl http2;
    server_name your-domain.com;
    
    ssl_certificate /etc/letsencrypt/live/your-domain.com/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/your-domain.com/privkey.pem;
    
    location / {
        proxy_pass https://localhost:18443;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```

### 3. 日志轮转

创建日志轮转配置：
```bash
sudo nano /etc/logrotate.d/httpsserver
```

内容：
```
/opt/httpsserver/logs/*.log {
    daily
    rotate 30
    compress
    delaycompress
    missingok
    notifempty
    create 644 httpsserver httpsserver
    postrotate
        systemctl reload httpsserver
    endscript
}
```

### 4. 监控和告警

使用systemd的监控功能：
```bash
# 查看服务状态
sudo systemctl status httpsserver

# 查看最近的日志
sudo journalctl -u httpsserver --since "1 hour ago"

# 实时查看日志
sudo journalctl -u httpsserver -f
```

## 常见问题

### 1. 端口被占用
```bash
# 查看端口占用
sudo netstat -tulpn | grep 18443
sudo lsof -i :18443

# 停止占用端口的进程
sudo kill -9 PID
```

### 2. 权限问题
```bash
# 检查文件权限
ls -la /opt/httpsserver/
sudo chown -R httpsserver:httpsserver /opt/httpsserver/
```

### 3. 数据库连接失败
```bash
# 测试数据库连接
mysql -u root -p metadata_db -e "SELECT COUNT(*) FROM table_metadata;"
```

### 4. SSL证书问题
```bash
# 重新生成自签名证书
openssl req -x509 -newkey rsa:4096 -keyout server.key -out server.crt -days 365 -nodes -subj "/CN=localhost"
```

## 维护命令

```bash
# 查看服务状态
sudo systemctl status httpsserver

# 重启服务
sudo systemctl restart httpsserver

# 查看日志
sudo journalctl -u httpsserver -f

# 检查配置
./httpsserver-linux --help

# 备份数据库
mysqldump -u root -p metadata_db > backup_$(date +%Y%m%d_%H%M%S).sql
``` 