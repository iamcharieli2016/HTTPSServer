package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"httpsserver/internal/auth"
	"httpsserver/internal/config"
	"httpsserver/internal/database"
	"httpsserver/internal/handler"
	"httpsserver/internal/utils"
)

func main() {
	// 加载配置
	cfg := config.Load()

	// 初始化数据库连接
	db, err := database.New(cfg)
	if err != nil {
		log.Fatal("数据库初始化失败:", err)
	}
	defer db.Close()

	// 初始化认证服务
	authSvc := auth.New(cfg)

	// 初始化处理器
	h := handler.New(db, authSvc)

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
	r.POST("/service/:serviceId", h.HandleServiceRequest)
	r.GET("/health", h.HandleHealthCheck)

	// 生成自签名证书（仅用于开发测试）
	utils.GenerateSelfSignedCert(cfg.Server.CertFile, cfg.Server.KeyFile)

	// 启动HTTPS服务器
	log.Printf("HTTPS服务器启动在端口%s...", cfg.Server.Port)
	log.Fatal(r.RunTLS(":"+cfg.Server.Port, cfg.Server.CertFile, cfg.Server.KeyFile))
}
