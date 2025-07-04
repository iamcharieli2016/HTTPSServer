package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// 请求结构体
type ServiceRequest struct {
	Params struct {
		ColumnComment string `json:"columnComment"`
		ColumnType    string `json:"columnType"`
		ColumnName    string `json:"columnName"`
		TableComment  string `json:"tableComment"`
		TableName     string `json:"tableName"`
		TableSchema   string `json:"tableSchema"`
		DBType        string `json:"dbType"`
	} `json:"params"`
	ServiceID    string `json:"serviceId"`
	ShowCount    string `json:"showCount"`
	Offset       int    `json:"offset"`
	Limit        int    `json:"limit"`
	UserID       string `json:"userId"`
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

// 响应结构体
type ServiceResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Total   int64       `json:"total,omitempty"`
	Message string      `json:"message,omitempty"`
}

// 数据库表结构信息
type TableMetadata struct {
	TableSchema   string `json:"tableSchema"`
	TableName     string `json:"tableName"`
	TableComment  string `json:"tableComment"`
	ColumnName    string `json:"columnName"`
	ColumnType    string `json:"columnType"`
	ColumnComment string `json:"columnComment"`
	DBType        string `json:"dbType"`
}

var db *sql.DB
var config *Config

func main() {
	// 加载配置
	config = getConfig()

	// 初始化数据库连接
	initDB()
	defer db.Close()

	// 设置Gin为发布模式
	gin.SetMode(gin.ReleaseMode)

	// 创建Gin引擎
	r := gin.Default()

	// 添加请求日志记录中间件
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[%s] %s \"%s %s %s\" %d %s \"%s\" \"%s\" %d\n",
			param.TimeStamp.Format("2006/01/02 - 15:04:05"),
			param.ClientIP,
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
			param.BodySize,
		)
	}))

	// 添加CORS中间件
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// 定义路由
	r.POST("/service/:serviceId", handleServiceRequest)

	// 生成自签名证书（仅用于开发测试）
	generateSelfSignedCert()

	// 启动HTTPS服务器
	log.Printf("HTTPS服务器启动在端口%s...", config.Server.Port)
	log.Fatal(r.RunTLS(":"+config.Server.Port, config.Server.CertFile, config.Server.KeyFile))
}

func initDB() {
	var err error
	// 连接MySQL数据库
	dsn := config.GetDSN()
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}

	// 测试数据库连接
	err = db.Ping()
	if err != nil {
		log.Fatal("数据库ping失败:", err)
	}

	log.Println("数据库连接成功")
}

func handleServiceRequest(c *gin.Context) {
	serviceId := c.Param("serviceId")

	// 记录请求详情
	log.Printf("🔵 接收到请求 - ServiceID: %s, ClientIP: %s", serviceId, c.ClientIP())

	var req ServiceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("❌ 请求参数解析失败: %v", err)
		errorResponse := ServiceResponse{
			Success: false,
			Message: "请求参数格式错误: " + err.Error(),
		}
		// 打印错误返回内容
		if responseJSON, jsonErr := json.Marshal(errorResponse); jsonErr == nil {
			log.Printf("📄 错误返回内容: %s", string(responseJSON))
		}
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	// 记录请求参数
	log.Printf("📋 请求参数 - UserID: %s, ClientID: %s, Offset: %d, Limit: %d",
		req.UserID, req.ClientID, req.Offset, req.Limit)

	// 验证客户端信息
	if req.ClientID != config.Auth.ClientID || req.ClientSecret != config.Auth.ClientSecret {
		log.Printf("🚫 客户端认证失败 - ClientID: %s", req.ClientID)
		authErrorResponse := ServiceResponse{
			Success: false,
			Message: "客户端认证失败",
		}
		// 打印认证失败返回内容
		if responseJSON, jsonErr := json.Marshal(authErrorResponse); jsonErr == nil {
			log.Printf("📄 认证失败返回内容: %s", string(responseJSON))
		}
		c.JSON(http.StatusUnauthorized, authErrorResponse)
		return
	}

	log.Printf("✅ 客户端认证成功 - UserID: %s", req.UserID)

	// 根据serviceId处理不同的服务
	switch serviceId {
	case "D_A_BSPDMETA":
		handleMetadataQuery(c, req)
	default:
		notFoundResponse := ServiceResponse{
			Success: false,
			Message: "未找到指定的服务: " + serviceId,
		}
		log.Printf("❌ 服务未找到 - ServiceID: %s", serviceId)
		// 打印服务未找到返回内容
		if responseJSON, jsonErr := json.Marshal(notFoundResponse); jsonErr == nil {
			log.Printf("📄 服务未找到返回内容: %s", string(responseJSON))
		}
		c.JSON(http.StatusNotFound, notFoundResponse)
	}
}

func handleMetadataQuery(c *gin.Context, req ServiceRequest) {
	log.Printf("🔍 开始处理元数据查询请求")

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
		err := db.QueryRow(countQuery, args...).Scan(&total)
		if err != nil {
			log.Printf("❌ 查询总数失败: %v", err)
			countErrorResponse := ServiceResponse{
				Success: false,
				Message: "查询总数失败: " + err.Error(),
			}
			// 打印查询总数失败返回内容
			if responseJSON, jsonErr := json.Marshal(countErrorResponse); jsonErr == nil {
				log.Printf("📄 查询总数失败返回内容: %s", string(responseJSON))
			}
			c.JSON(http.StatusInternalServerError, countErrorResponse)
			return
		}
		log.Printf("📊 数据总数: %d", total)
	}

	// 添加分页
	query += " ORDER BY table_schema, table_name, column_name LIMIT ? OFFSET ?"
	args = append(args, req.Limit, req.Offset)

	// 执行查询
	log.Printf("🗃️  执行数据库查询...")
	rows, err := db.Query(query, args...)
	if err != nil {
		log.Printf("❌ 数据库查询失败: %v", err)
		queryErrorResponse := ServiceResponse{
			Success: false,
			Message: "数据库查询失败: " + err.Error(),
		}
		// 打印数据库查询失败返回内容
		if responseJSON, jsonErr := json.Marshal(queryErrorResponse); jsonErr == nil {
			log.Printf("📄 数据库查询失败返回内容: %s", string(responseJSON))
		}
		c.JSON(http.StatusInternalServerError, queryErrorResponse)
		return
	}
	defer rows.Close()

	// 处理查询结果
	var results []TableMetadata
	for rows.Next() {
		var meta TableMetadata
		err := rows.Scan(
			&meta.TableSchema, &meta.TableName, &meta.TableComment,
			&meta.ColumnName, &meta.ColumnType, &meta.ColumnComment,
			&meta.DBType,
		)
		if err != nil {
			log.Printf("❌ 数据扫描失败: %v", err)
			scanErrorResponse := ServiceResponse{
				Success: false,
				Message: "数据扫描失败: " + err.Error(),
			}
			// 打印数据扫描失败返回内容
			if responseJSON, jsonErr := json.Marshal(scanErrorResponse); jsonErr == nil {
				log.Printf("📄 数据扫描失败返回内容: %s", string(responseJSON))
			}
			c.JSON(http.StatusInternalServerError, scanErrorResponse)
			return
		}
		results = append(results, meta)
	}

	// 构建响应
	response := ServiceResponse{
		Success: true,
		Data:    results,
	}

	if req.ShowCount == "true" {
		response.Total = total
	}

	log.Printf("✅ 查询成功完成 - 返回记录数: %d", len(results))
	if req.ShowCount == "true" {
		log.Printf("📈 总记录数: %d, 当前页: %d-%d", total, req.Offset, req.Offset+len(results))
	}

	// 打印详细的返回数据内容
	responseJSON, err := json.Marshal(response)
	if err != nil {
		log.Printf("❌ JSON序列化失败: %v", err)
	} else {
		log.Printf("📄 返回数据内容:")
		log.Printf("%s", string(responseJSON))
	}

	c.JSON(http.StatusOK, response)
}

func generateSelfSignedCert() {
	// 检查证书文件是否存在
	if fileExists(config.Server.CertFile) && fileExists(config.Server.KeyFile) {
		return
	}

	log.Println("生成自签名SSL证书...")

	// 这里简化处理，实际生产环境应该使用真实证书
	// 您可以使用openssl命令生成：
	// openssl req -x509 -newkey rsa:4096 -keyout server.key -out server.crt -days 365 -nodes -subj "/CN=localhost"

	log.Printf("请手动生成SSL证书文件:")
	log.Printf("openssl req -x509 -newkey rsa:4096 -keyout %s -out %s -days 365 -nodes -subj \"/CN=localhost\"", config.Server.KeyFile, config.Server.CertFile)
	log.Fatal("SSL证书文件不存在，请先生成证书文件")
}

func fileExists(filename string) bool {
	_, err := http.Dir(".").Open(filename)
	return err == nil
}
