package dto

import "35home/common"

// 统一响应结构
type Response struct {
	Code     common.ErrorCode `json:"code"`
	Msg      string           `json:"msg"`
	RespBody interface{}      `json:"respBody,omitempty"`
}

// NewSuccessResponse 创建成功响应
func NewSuccessResponse(respBody interface{}) Response {
	return Response{
		Code:     common.ErrSuccess,
		Msg:      common.ErrorMessages[common.ErrSuccess],
		RespBody: respBody,
	}
}

// NewErrorResponse 创建错误响应
func NewErrorResponse(code common.ErrorCode) Response {
	return Response{
		Code: code,
		Msg:  common.ErrorMessages[code],
	}
}

// NewErrorResponseWithBody 创建带响应体的错误响应
func NewErrorResponseWithBody(code common.ErrorCode, respBody interface{}) Response {
	return Response{
		Code:     code,
		Msg:      common.ErrorMessages[code],
		RespBody: respBody,
	}
}
