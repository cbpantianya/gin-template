package utils

import "github.com/gin-gonic/gin"

// SuccessResp 请求成功返回格式化后的JSON
func SuccessResp(data interface{}) (httpCode int, h gin.H) {
	return 200, gin.H{
		"error": 0,
		"msg":   "success",
		"data":  data,
	}
}

// ErrorResp 请求失败返回格式化后的JSON 请严格使用错误码
func ErrorResp(code int, msg interface{}, data interface{}) (int, gin.H) {
	return 200, gin.H{
		"error": code,
		"msg":   msg,
		"data":  data,
	}
}