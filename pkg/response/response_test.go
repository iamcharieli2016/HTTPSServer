package response

import (
	"reflect"
	"testing"
)

func TestNewSuccessResponse(t *testing.T) {
	data := map[string]string{"key": "value"}
	resp := NewSuccessResponse(data)

	if !resp.Success {
		t.Error("Expected Success to be true")
	}

	if !reflect.DeepEqual(resp.Data, data) {
		t.Error("Expected data to match input")
	}

	if resp.Total != 0 {
		t.Error("Expected Total to be 0 for simple success response")
	}

	if resp.Message != "" {
		t.Error("Expected Message to be empty for success response")
	}
}

func TestNewSuccessResponseWithTotal(t *testing.T) {
	data := []string{"item1", "item2"}
	total := int64(100)
	resp := NewSuccessResponseWithTotal(data, total)

	if !resp.Success {
		t.Error("Expected Success to be true")
	}

	if !reflect.DeepEqual(resp.Data, data) {
		t.Error("Expected data to match input")
	}

	if resp.Total != total {
		t.Errorf("Expected Total to be %d, got %d", total, resp.Total)
	}

	if resp.Message != "" {
		t.Error("Expected Message to be empty for success response")
	}
}

func TestNewErrorResponse(t *testing.T) {
	message := "Test error message"
	resp := NewErrorResponse(message)

	if resp.Success {
		t.Error("Expected Success to be false")
	}

	if resp.Data != nil {
		t.Error("Expected Data to be nil for error response")
	}

	if resp.Total != 0 {
		t.Error("Expected Total to be 0 for error response")
	}

	if resp.Message != message {
		t.Errorf("Expected Message to be '%s', got '%s'", message, resp.Message)
	}
}

func TestServiceResponseStructure(t *testing.T) {
	// 测试结构体字段的JSON标签
	resp := ServiceResponse{
		Success: true,
		Data:    "test",
		Total:   10,
		Message: "test message",
	}

	// 验证结构体可以正常创建和访问
	if resp.Success != true {
		t.Error("Failed to set Success field")
	}

	if resp.Data != "test" {
		t.Error("Failed to set Data field")
	}

	if resp.Total != 10 {
		t.Error("Failed to set Total field")
	}

	if resp.Message != "test message" {
		t.Error("Failed to set Message field")
	}
}
