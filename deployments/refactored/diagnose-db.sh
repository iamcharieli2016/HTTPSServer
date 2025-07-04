#!/bin/bash

# 数据库连接诊断脚本
# 使用说明：chmod +x diagnose-db.sh && ./diagnose-db.sh

echo "🔍 数据库连接问题诊断"
echo "====================="

# 1. 检查MySQL服务状态
echo "1. 检查MySQL服务状态..."
if systemctl is-active --quiet mysql; then
    echo "✅ MySQL服务正在运行"
elif systemctl is-active --quiet mysqld; then
    echo "✅ MySQL服务正在运行 (mysqld)"
else
    echo "❌ MySQL服务未运行"
    echo "📋 启动MySQL服务："
    echo "   sudo systemctl start mysql"
    echo "   # 或者"
    echo "   sudo systemctl start mysqld"
    echo ""
fi

# 2. 检查MySQL端口
echo "2. 检查MySQL端口监听..."
netstat -tlnp | grep :3306
if [ $? -eq 0 ]; then
    echo "✅ MySQL端口3306正在监听"
else
    echo "❌ MySQL端口3306未监听"
    echo "📋 可能的原因："
    echo "   - MySQL服务未启动"
    echo "   - MySQL配置了其他端口"
    echo "   - 防火墙阻止了端口"
    echo ""
fi

# 3. 检查本地连接
echo "3. 测试本地MySQL连接..."
if command -v mysql &> /dev/null; then
    echo "📋 测试MySQL连接（请输入密码）："
    mysql -u root -p -e "SELECT 1;" 2>/dev/null
    if [ $? -eq 0 ]; then
        echo "✅ MySQL连接成功"
    else
        echo "❌ MySQL连接失败"
        echo "📋 可能的原因："
        echo "   - 密码错误"
        echo "   - 用户权限问题"
        echo "   - MySQL服务异常"
        echo ""
    fi
else
    echo "⚠️  MySQL客户端未安装"
    echo "📋 安装MySQL客户端："
    echo "   sudo apt install mysql-client  # Ubuntu/Debian"
    echo "   sudo yum install mysql         # CentOS/RHEL"
    echo ""
fi

# 4. 检查数据库是否存在
echo "4. 检查数据库是否存在..."
if command -v mysql &> /dev/null; then
    echo "📋 检查metadata_db数据库（请输入密码）："
    mysql -u root -p -e "USE metadata_db; SELECT COUNT(*) as table_count FROM table_metadata;" 2>/dev/null
    if [ $? -eq 0 ]; then
        echo "✅ 数据库metadata_db存在且有数据"
    else
        echo "❌ 数据库metadata_db不存在或无数据"
        echo "📋 初始化数据库："
        echo "   mysql -u root -p < database.sql"
        echo ""
    fi
fi

# 5. 检查配置文件
echo "5. 检查配置文件..."
if [ -f "config.env" ]; then
    echo "✅ 配置文件config.env存在"
    echo "📋 当前数据库配置："
    grep -E "^DB_" config.env | head -5
    echo ""
else
    echo "❌ 配置文件config.env不存在"
    echo "📋 创建配置文件："
    echo "   cp config.env.example config.env"
    echo "   nano config.env"
    echo ""
fi

echo "🚀 解决方案建议："
echo "=================="
echo "1. 启动MySQL服务："
echo "   sudo systemctl start mysql"
echo "   sudo systemctl enable mysql"
echo ""
echo "2. 初始化数据库："
echo "   mysql -u root -p < database.sql"
echo ""
echo "3. 修改配置文件："
echo "   cp config.env.example config.env"
echo "   nano config.env  # 设置正确的数据库密码"
echo ""
echo "4. 重启服务："
echo "   ./stop-linux.sh"
echo "   ./start-linux.sh"
echo ""
echo "5. 查看日志："
echo "   tail -f logs/server.log" 