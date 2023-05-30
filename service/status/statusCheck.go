package status

import (
	"context"
	"gin-template/v2/server"
	"gin-template/v2/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type StatusCheckRespone struct {
	MySQL bool `json:"mysql"`
	Redis bool `json:"redis"`
}

func StatusCheck(s *server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var responese StatusCheckRespone
		// 使用自动生成空模板，检查MySQL连接
		err := s.MySQL.AutoMigrate()
		if err != nil {
			responese.MySQL = false
			s.Logger.Error().Msgf("MySQL Connection Error: %v", err)
		}else {
			responese.MySQL = true
		}

		// 测试Redis连接
		err = s.Redis.Set(context.Background(), "connect_test", "1", time.Second).Err()
		if err != nil {
			responese.Redis = false
			s.Logger.Error().Msgf("Redis Connection Error: %v", err)
		}else {
			responese.Redis = true
		}

		ctx.JSON(utils.SuccessResp(responese))

	}
}
