package dto

// 统一响应结构
type Response struct {
	Code      string      `json:"code"`
	Msg       string      `json:"msg"`
	RespBody  interface{} `json:"respBody,omitempty"`
}
