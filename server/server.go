package server

import (
	"context"
	"fmt"
	"gin-template/v2/middleware"
	"gin-template/v2/utils"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var S Server

type Server struct {
	Engine *gin.Engine
	Logger zerolog.Logger
	MySQL  *gorm.DB
	Redis  *redis.Client
}

func Init() {
	S.Logger = utils.LogInit()
	utils.CFInit()
	gin.SetMode(gin.ReleaseMode)
	S.Engine = gin.New()
	S.Engine.Use(middleware.GinLoggerMiddleware(&S.Logger))
	S.Engine.Use(gin.Recovery())
	S.Engine.Use(middleware.GinCORS())
	// 连接数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", utils.GConfig.MySQL.User, utils.GConfig.MySQL.Password, utils.GConfig.MySQL.Host, utils.GConfig.MySQL.Port, utils.GConfig.MySQL.Database)
	var err error
	S.MySQL, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err!= nil {
        S.Logger.Fatal().Err(err).Msg("Failed to connect to mysql")
    }
	// 连接Redis
	S.Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", utils.GConfig.Redis.Host, utils.GConfig.Redis.Port),
		Password: utils.GConfig.Redis.Password,
		DB:       utils.GConfig.Redis.Database,
	})

	// 测试Redis连接
	err = S.Redis.Ping(context.Background()).Err()
	if err != nil {
		S.Logger.Fatal().Err(err).Msg("Failed to connect to redis")
	}


	RegisterService2Server()
}

func Start() {
	S.Logger.Info().Msg(fmt.Sprintf("http://%s:%d", utils.GConfig.HTTPServer.Host, utils.GConfig.HTTPServer.Port))
	err := S.Engine.Run(fmt.Sprintf("%s:%d", utils.GConfig.HTTPServer.Host, utils.GConfig.HTTPServer.Port))
	if err != nil {
		S.Logger.Fatal().Err(err).Msg(fmt.Sprintf("http://%s:%d", utils.GConfig.HTTPServer.Host, utils.GConfig.HTTPServer.Port))
	}
}
