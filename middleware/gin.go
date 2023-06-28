package middleware

import (
	"fmt"
	"gin-template/v2/utils"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func GinLoggerMiddleware(log *zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		// 写入日志
		log.Info().Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Str("ip", c.ClientIP()).
			Str("code", fmt.Sprintf("%d", c.Writer.Status())).
			Msg("")
	}
}

func GinCORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		if IfInOrigin(origin) {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Writer.Header().Set("Vary", "Origin")
		}

		c.Writer.Header().Set("Access-Control-Allow-Headers", GenerateHeaders())
		c.Writer.Header().Set("Access-Control-Allow-Methods", GenerateMethods())
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
	}
}

func IfInOrigin(origin string) bool {
	if origin == "" {
		return false
	}
	for _, value := range utils.GConfig.HTTPServer.Origin {
		if origin == value || "*" == value {
			return true
		}
	}
	return false
}

func GenerateMethods() string {
	str := ""
	for index, value := range utils.GConfig.HTTPServer.Methods {
		if index == 0 {
			str = str + value
		} else {
			str = str + ", " + value
		}
	}
	return str
}

func GenerateHeaders() string {
	str := ""
	for index, value := range utils.GConfig.HTTPServer.Headers {
		if index == 0 {
			str = str + value
		} else {
			str = str + ", " + value
		}
	}
	return str
}
