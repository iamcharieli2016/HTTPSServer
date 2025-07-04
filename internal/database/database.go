package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"httpsserver/internal/config"
	"httpsserver/internal/model"
)

// DB 数据库连接实例
type DB struct {
	conn *sql.DB
}

// New 创建新的数据库连接
func New(cfg *config.Config) (*DB, error) {
	// 连接MySQL数据库
	dsn := cfg.GetDSN()
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("数据库连接失败: %w", err)
	}

	// 测试数据库连接
	err = conn.Ping()
	if err != nil {
		return nil, fmt.Errorf("数据库ping失败: %w", err)
	}

	log.Println("数据库连接成功")
	return &DB{conn: conn}, nil
}

// Close 关闭数据库连接
func (db *DB) Close() error {
	return db.conn.Close()
}

// QueryMetadata 查询元数据
func (db *DB) QueryMetadata(req model.ServiceRequest) ([]model.TableMetadata, int64, error) {
	// 构建查询SQL
	query := "SELECT table_schema, table_name, table_comment, column_name, column_type, column_comment, db_type FROM table_metadata WHERE 1=1"
	var args []interface{}

	// 记录搜索条件
	searchConditions := []string{}
	if req.Params.TableSchema != "" {
		searchConditions = append(searchConditions, fmt.Sprintf("TableSchema: %s", req.Params.TableSchema))
	}
	if req.Params.TableName != "" {
		searchConditions = append(searchConditions, fmt.Sprintf("TableName: %s", req.Params.TableName))
	}
	if req.Params.ColumnName != "" {
		searchConditions = append(searchConditions, fmt.Sprintf("ColumnName: %s", req.Params.ColumnName))
	}
	if len(searchConditions) > 0 {
		log.Printf("🔎 搜索条件: %v", searchConditions)
	} else {
		log.Printf("📜 查询所有数据（无搜索条件）")
	}

	// 根据参数添加WHERE条件
	if req.Params.TableSchema != "" {
		query += " AND table_schema LIKE ?"
		args = append(args, "%"+req.Params.TableSchema+"%")
	}
	if req.Params.TableName != "" {
		query += " AND table_name LIKE ?"
		args = append(args, "%"+req.Params.TableName+"%")
	}
	if req.Params.TableComment != "" {
		query += " AND table_comment LIKE ?"
		args = append(args, "%"+req.Params.TableComment+"%")
	}
	if req.Params.ColumnName != "" {
		query += " AND column_name LIKE ?"
		args = append(args, "%"+req.Params.ColumnName+"%")
	}
	if req.Params.ColumnType != "" {
		query += " AND column_type LIKE ?"
		args = append(args, "%"+req.Params.ColumnType+"%")
	}
	if req.Params.ColumnComment != "" {
		query += " AND column_comment LIKE ?"
		args = append(args, "%"+req.Params.ColumnComment+"%")
	}
	if req.Params.DBType != "" {
		query += " AND db_type LIKE ?"
		args = append(args, "%"+req.Params.DBType+"%")
	}

	// 获取总数
	var total int64
	if req.ShowCount == "true" {
		countQuery := "SELECT COUNT(*) FROM (" + query + ") AS count_table"
		err := db.conn.QueryRow(countQuery, args...).Scan(&total)
		if err != nil {
			return nil, 0, fmt.Errorf("查询总数失败: %w", err)
		}
		log.Printf("📊 数据总数: %d", total)
	}

	// 添加分页
	query += " ORDER BY table_schema, table_name, column_name LIMIT ? OFFSET ?"
	args = append(args, req.Limit, req.Offset)

	// 执行查询
	log.Printf("🗃️  执行数据库查询...")
	rows, err := db.conn.Query(query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("数据库查询失败: %w", err)
	}
	defer rows.Close()

	// 处理查询结果
	var results []model.TableMetadata
	for rows.Next() {
		var meta model.TableMetadata
		err := rows.Scan(
			&meta.TableSchema, &meta.TableName, &meta.TableComment,
			&meta.ColumnName, &meta.ColumnType, &meta.ColumnComment,
			&meta.DBType,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("数据扫描失败: %w", err)
		}
		results = append(results, meta)
	}

	log.Printf("✅ 查询成功完成 - 返回记录数: %d", len(results))
	if req.ShowCount == "true" {
		log.Printf("📈 总记录数: %d, 当前页: %d-%d", total, req.Offset, req.Offset+len(results))
	}

	return results, total, nil
}
