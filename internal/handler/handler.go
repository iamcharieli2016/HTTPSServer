package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"httpsserver/internal/auth"
	"httpsserver/internal/database"
	"httpsserver/internal/model"
	"httpsserver/pkg/response"
)

// Handler HTTP处理器
type Handler struct {
	db      *database.DB
	authSvc *auth.Service
}

// New 创建新的处理器
func New(db *database.DB, authSvc *auth.Service) *Handler {
	return &Handler{
		db:      db,
		authSvc: authSvc,
	}
}

// HandleServiceRequest 处理服务请求
func (h *Handler) HandleServiceRequest(c *gin.Context) {
	serviceId := c.Param("serviceId")

	// 记录请求详情
	log.Printf("🔵 接收到请求 - ServiceID: %s, ClientIP: %s", serviceId, c.ClientIP())

	var req model.ServiceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("❌ 请求参数解析失败: %v", err)
		errorResponse := response.NewErrorResponse("请求参数格式错误: " + err.Error())

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
	if !h.authSvc.Authenticate(req) {
		authErrorResponse := response.NewErrorResponse("客户端认证失败")

		// 打印认证失败返回内容
		if responseJSON, jsonErr := json.Marshal(authErrorResponse); jsonErr == nil {
			log.Printf("📄 认证失败返回内容: %s", string(responseJSON))
		}

		c.JSON(http.StatusUnauthorized, authErrorResponse)
		return
	}

	// 根据serviceId处理不同的服务
	switch serviceId {
	case "D_A_BSPDMETA":
		h.HandleMetadataQuery(c, req)
	default:
		notFoundResponse := response.NewErrorResponse("未找到指定的服务: " + serviceId)
		log.Printf("❌ 服务未找到 - ServiceID: %s", serviceId)

		// 打印服务未找到返回内容
		if responseJSON, jsonErr := json.Marshal(notFoundResponse); jsonErr == nil {
			log.Printf("📄 服务未找到返回内容: %s", string(responseJSON))
		}

		c.JSON(http.StatusNotFound, notFoundResponse)
	}
}

// HandleMetadataQuery 处理元数据查询
func (h *Handler) HandleMetadataQuery(c *gin.Context, req model.ServiceRequest) {
	log.Printf("🔍 开始处理元数据查询请求")

	// 查询数据库
	results, total, err := h.db.QueryMetadata(req)
	if err != nil {
		log.Printf("❌ 数据库查询失败: %v", err)
		errorResponse := response.NewErrorResponse("数据库查询失败: " + err.Error())

		// 打印数据库查询失败返回内容
		if responseJSON, jsonErr := json.Marshal(errorResponse); jsonErr == nil {
			log.Printf("📄 数据库查询失败返回内容: %s", string(responseJSON))
		}

		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	// 构建响应
	var resp response.ServiceResponse
	if req.ShowCount == "true" {
		resp = response.NewSuccessResponseWithTotal(results, total)
	} else {
		resp = response.NewSuccessResponse(results)
	}

	// 打印详细的返回数据内容
	responseJSON, err := json.Marshal(resp)
	if err != nil {
		log.Printf("❌ JSON序列化失败: %v", err)
	} else {
		log.Printf("📄 返回数据内容:")
		log.Printf("%s", string(responseJSON))
	}

	c.JSON(http.StatusOK, resp)
}

// HandleHealthCheck 健康检查端点
func (h *Handler) HandleHealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"message": "HTTPS Server is running",
	})
}
