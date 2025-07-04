package response

// ServiceResponse 服务响应结构体
type ServiceResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Total   int64       `json:"total,omitempty"`
	Message string      `json:"message,omitempty"`
}

// NewSuccessResponse 创建成功响应
func NewSuccessResponse(data interface{}) ServiceResponse {
	return ServiceResponse{
		Success: true,
		Data:    data,
	}
}

// NewSuccessResponseWithTotal 创建带总数的成功响应
func NewSuccessResponseWithTotal(data interface{}, total int64) ServiceResponse {
	return ServiceResponse{
		Success: true,
		Data:    data,
		Total:   total,
	}
}

// NewErrorResponse 创建错误响应
func NewErrorResponse(message string) ServiceResponse {
	return ServiceResponse{
		Success: false,
		Message: message,
	}
}
