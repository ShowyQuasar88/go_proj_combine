package response

import (
	v1 "backend/api/v1"
	"google.golang.org/protobuf/types/known/anypb"
)

const (
	defaultSuccessMessage = "操作成功"
	defaultFailMessage    = "操作失败"
)

// Success 返回成功响应
func Success() *v1.Response {
	return &v1.Response{
		Code:    v1.ErrorCode_SUCCESS,
		Success: true,
		Message: defaultSuccessMessage,
	}
}

// SuccessWithMsg 返回带消息的成功响应
func SuccessWithMsg(msg string) *v1.Response {
	return &v1.Response{
		Code:    v1.ErrorCode_SUCCESS,
		Success: true,
		Message: msg,
	}
}

// SuccessWithData 返回带数据的成功响应
func SuccessWithData(data *anypb.Any) *v1.Response {
	return &v1.Response{
		Code:    v1.ErrorCode_SUCCESS,
		Success: true,
		Data:    data,
	}
}

// SuccessWithDataAndMsg 返回带数据和消息的成功响应
func SuccessWithDataAndMsg(msg string, data *anypb.Any) *v1.Response {
	return &v1.Response{
		Code:    v1.ErrorCode_SUCCESS,
		Success: true,
		Message: msg,
		Data:    data,
	}
}

// Error 返回错误响应
func Error(code v1.ErrorCode, message string) *v1.Response {
	return &v1.Response{
		Code:    code,
		Success: false,
		Message: message,
	}
}

// ErrorWithData 返回带数据的错误响应
func ErrorWithData(code v1.ErrorCode, message string, data *anypb.Any) *v1.Response {
	return &v1.Response{
		Code:    code,
		Success: false,
		Message: message,
		Data:    data,
	}
}
