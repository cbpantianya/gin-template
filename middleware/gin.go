package middleware

import (
	"fmt"

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
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
	}
}

