# Troubleshooting Guide

This guide helps you resolve common issues with the HTTPS Server.

## 🔗 Database Connection Issues

### Error: `dial tcp [::1]:3306: connect: connection refused`

**Symptoms:** Server cannot connect to MySQL database.

**Causes:**
- MySQL service is not running
- Wrong database configuration
- Configuration file not loaded

**Solutions:**

1. **Check MySQL service**
   ```bash
   # Check if MySQL is running
   sudo systemctl status mysql
   # or
   sudo systemctl status mysqld
   
   # Start MySQL if not running
   sudo systemctl start mysql
   ```

2. **Verify configuration**
   ```bash
   # Check if config.env exists
   ls -la config.env
   
   # Verify configuration content
   cat config.env
   ```

3. **Test database connection**
   ```bash
   mysql -h localhost -P 3306 -u root -p
   ```

4. **Check if database and table exist**
   ```bash
   mysql -u root -p -e "USE metadata_db; SHOW TABLES;"
   ```

## 🔐 SSL Certificate Issues

### Error: `SSL certificate file not exist`

**Solutions:**

1. **Generate self-signed certificate (development)**
   ```bash
   openssl req -x509 -newkey rsa:4096 -keyout server.key -out server.crt -days 365 -nodes -subj "/CN=localhost"
   ```

2. **Use production certificate**
   ```bash
   # Place your certificate files
   cp /path/to/your/cert.pem server.crt
   cp /path/to/your/key.pem server.key
   ```

## 🔑 Authentication Issues

### Error: `客户端认证失败`

**Causes:**
- Wrong CLIENT_ID or CLIENT_SECRET
- Configuration not loaded

**Solutions:**

1. **Check configuration**
   ```bash
   grep -E "CLIENT_" config.env
   ```

2. **Verify request credentials**
   ```json
   {
     "clientId": "eplat",
     "clientSecret": "eplat2019111214440"
   }
   ```

## 🚀 Port Issues

### Error: `bind: address already in use`

**Solutions:**

1. **Find process using the port**
   ```bash
   lsof -i :18443
   netstat -tulpn | grep 18443
   ```

2. **Kill the process**
   ```bash
   sudo kill -9 <PID>
   ```

3. **Change port**
   ```bash
   # Edit config.env
   SERVER_PORT=18444
   ```

## 📊 Performance Issues

### High memory usage

**Solutions:**

1. **Enable log rotation**
   ```bash
   # Create logrotate config
   sudo nano /etc/logrotate.d/httpsserver
   ```

2. **Monitor connections**
   ```bash
   # Check MySQL connections
   mysql -u root -p -e "SHOW PROCESSLIST;"
   ```

## 🔍 Debugging

### Enable debug logging

1. **Set environment variable**
   ```bash
   export GIN_MODE=debug
   ```

2. **Check application logs**
   ```bash
   tail -f logs/server.log
   ```

3. **Check system logs**
   ```bash
   sudo journalctl -u httpsserver -f
   ```

## 📋 Common Commands

### Check service status
```bash
# Systemd service
sudo systemctl status httpsserver

# Manual process
ps aux | grep httpsserver
```

### View logs
```bash
# Application logs
tail -f logs/server.log

# System logs
sudo journalctl -u httpsserver --since "1 hour ago"
```

### Test API
```bash
# Basic connectivity test
curl -k https://localhost:18443/

# Full API test
curl -k -X POST https://localhost:18443/service/D_A_BSPDMETA \
  -H 'Content-Type: application/json' \
  -d '{"serviceId":"D_A_BSPDMETA","clientId":"eplat","clientSecret":"eplat2019111214440","userId":"test","params":{},"offset":0,"limit":1}'
```

## 🆘 Getting Help

If you're still experiencing issues:

1. **Check logs first**
   ```bash
   tail -100 logs/server.log
   ```

2. **Gather system information**
   ```bash
   uname -a
   go version
   mysql --version
   ```

3. **Create an issue** with:
   - Error message
   - Log output
   - System information
   - Steps to reproduce 