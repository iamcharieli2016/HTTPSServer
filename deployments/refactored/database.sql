-- 创建数据库
CREATE DATABASE IF NOT EXISTS metadata_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE metadata_db;

-- 创建表结构元数据表
CREATE TABLE IF NOT EXISTS table_metadata (
    id INT AUTO_INCREMENT PRIMARY KEY,
    table_schema VARCHAR(255) NOT NULL COMMENT '数据库名称',
    table_name VARCHAR(255) NOT NULL COMMENT '表名',
    table_comment TEXT COMMENT '表注释',
    column_name VARCHAR(255) NOT NULL COMMENT '列名',
    column_type VARCHAR(255) NOT NULL COMMENT '列类型',
    column_comment TEXT COMMENT '列注释',
    db_type VARCHAR(50) NOT NULL DEFAULT 'mysql' COMMENT '数据库类型',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_table_schema (table_schema),
    INDEX idx_table_name (table_name),
    INDEX idx_column_name (column_name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='表结构元数据';

-- 插入示例数据
INSERT INTO table_metadata (table_schema, table_name, table_comment, column_name, column_type, column_comment, db_type) VALUES
('test_db', 'users', '用户表', 'id', 'INT AUTO_INCREMENT', '用户ID', 'mysql'),
('test_db', 'users', '用户表', 'username', 'VARCHAR(50)', '用户名', 'mysql'),
('test_db', 'users', '用户表', 'email', 'VARCHAR(100)', '邮箱地址', 'mysql'),
('test_db', 'users', '用户表', 'created_at', 'TIMESTAMP', '创建时间', 'mysql'),
('test_db', 'users', '用户表', 'updated_at', 'TIMESTAMP', '更新时间', 'mysql'),

('test_db', 'products', '产品表', 'id', 'INT AUTO_INCREMENT', '产品ID', 'mysql'),
('test_db', 'products', '产品表', 'name', 'VARCHAR(255)', '产品名称', 'mysql'),
('test_db', 'products', '产品表', 'price', 'DECIMAL(10,2)', '产品价格', 'mysql'),
('test_db', 'products', '产品表', 'description', 'TEXT', '产品描述', 'mysql'),
('test_db', 'products', '产品表', 'created_at', 'TIMESTAMP', '创建时间', 'mysql'),

('test_db', 'orders', '订单表', 'id', 'INT AUTO_INCREMENT', '订单ID', 'mysql'),
('test_db', 'orders', '订单表', 'user_id', 'INT', '用户ID', 'mysql'),
('test_db', 'orders', '订单表', 'product_id', 'INT', '产品ID', 'mysql'),
('test_db', 'orders', '订单表', 'quantity', 'INT', '数量', 'mysql'),
('test_db', 'orders', '订单表', 'total_amount', 'DECIMAL(10,2)', '总金额', 'mysql'),
('test_db', 'orders', '订单表', 'status', 'ENUM(''pending'',''completed'',''cancelled'')', '订单状态', 'mysql'),
('test_db', 'orders', '订单表', 'created_at', 'TIMESTAMP', '创建时间', 'mysql');

-- 创建用户账户（可选）
CREATE USER IF NOT EXISTS 'api_user'@'localhost' IDENTIFIED BY 'api_password';
GRANT SELECT ON metadata_db.* TO 'api_user'@'localhost';
FLUSH PRIVILEGES; 