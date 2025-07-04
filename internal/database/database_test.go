package database

import (
	"testing"

	"httpsserver/internal/config"
	"httpsserver/internal/model"
)

func TestNew(t *testing.T) {
	cfg := &config.Config{}
	cfg.DB.Host = "localhost"
	cfg.DB.Port = "3306"
	cfg.DB.User = "test"
	cfg.DB.Password = "test"
	cfg.DB.Database = "test"
	cfg.DB.Charset = "utf8mb4"

	// 这个测试会失败因为没有实际的数据库连接，但我们可以验证返回值
	db, err := New(cfg)

	// 由于没有实际的数据库，这应该会返回错误
	if err == nil {
		t.Error("Expected error when connecting to non-existent database")
	}

	if db != nil {
		t.Error("Expected nil database when connection fails")
	}
}

func TestServiceRequestStructure(t *testing.T) {
	// 测试ServiceRequest结构体的创建和字段访问
	req := model.ServiceRequest{
		UserID:       "test_user",
		ClientID:     "test_client",
		ClientSecret: "test_secret",
		ServiceID:    "test_service",
		ShowCount:    "true",
		Offset:       0,
		Limit:        10,
	}

	// 设置参数
	req.Params.TableSchema = "test_schema"
	req.Params.TableName = "test_table"
	req.Params.TableComment = "test table comment"
	req.Params.ColumnName = "test_column"
	req.Params.ColumnType = "varchar(255)"
	req.Params.ColumnComment = "test column comment"
	req.Params.DBType = "mysql"

	// 验证字段设置
	if req.UserID != "test_user" {
		t.Error("Failed to set UserID")
	}

	if req.ClientID != "test_client" {
		t.Error("Failed to set ClientID")
	}

	if req.ClientSecret != "test_secret" {
		t.Error("Failed to set ClientSecret")
	}

	if req.ServiceID != "test_service" {
		t.Error("Failed to set ServiceID")
	}

	if req.ShowCount != "true" {
		t.Error("Failed to set ShowCount")
	}

	if req.Offset != 0 {
		t.Error("Failed to set Offset")
	}

	if req.Limit != 10 {
		t.Error("Failed to set Limit")
	}

	if req.Params.TableSchema != "test_schema" {
		t.Error("Failed to set TableSchema")
	}

	if req.Params.TableName != "test_table" {
		t.Error("Failed to set TableName")
	}

	if req.Params.TableComment != "test table comment" {
		t.Error("Failed to set TableComment")
	}

	if req.Params.ColumnName != "test_column" {
		t.Error("Failed to set ColumnName")
	}

	if req.Params.ColumnType != "varchar(255)" {
		t.Error("Failed to set ColumnType")
	}

	if req.Params.ColumnComment != "test column comment" {
		t.Error("Failed to set ColumnComment")
	}

	if req.Params.DBType != "mysql" {
		t.Error("Failed to set DBType")
	}
}

func TestTableMetadataStructure(t *testing.T) {
	// 测试TableMetadata结构体的创建和字段访问
	meta := model.TableMetadata{
		TableSchema:   "test_schema",
		TableName:     "test_table",
		TableComment:  "test table comment",
		ColumnName:    "test_column",
		ColumnType:    "varchar(255)",
		ColumnComment: "test column comment",
		DBType:        "mysql",
	}

	// 验证字段设置
	if meta.TableSchema != "test_schema" {
		t.Error("Failed to set TableSchema")
	}

	if meta.TableName != "test_table" {
		t.Error("Failed to set TableName")
	}

	if meta.TableComment != "test table comment" {
		t.Error("Failed to set TableComment")
	}

	if meta.ColumnName != "test_column" {
		t.Error("Failed to set ColumnName")
	}

	if meta.ColumnType != "varchar(255)" {
		t.Error("Failed to set ColumnType")
	}

	if meta.ColumnComment != "test column comment" {
		t.Error("Failed to set ColumnComment")
	}

	if meta.DBType != "mysql" {
		t.Error("Failed to set DBType")
	}
}
