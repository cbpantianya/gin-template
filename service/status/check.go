package status

import (
	"context"
	"gin-template/v2/server"
	"gin-template/v2/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type CheckResponse struct {
	MySQL bool `json:"mysql"`
	Redis bool `json:"redis"`
}

func Check(s *server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response CheckResponse
		// 使用自动生成空模板，检查MySQL连接
		err := s.MySQL.AutoMigrate()
		if err != nil {
			response.MySQL = false
			s.Logger.Error().Msgf("MySQL Connection Error: %v", err)
		} else {
			response.MySQL = true
		}

		// 测试Redis连接
		err = s.Redis.Set(context.Background(), "connect_test", "1", time.Second).Err()
		if err != nil {
			response.Redis = false
			s.Logger.Error().Msgf("Redis Connection Error: %v", err)
		} else {
			response.Redis = true
		}

		ctx.JSON(utils.SuccessResp(response))

	}
}
