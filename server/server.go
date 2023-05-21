package server

import (
	"gin-template/v2/middleware"
	"gin-template/v2/utils"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

var S Server

type Server struct {
	Engine *gin.Engine
	Logger zerolog.Logger
}

func Init() {
	S.Logger = utils.LogInit()
	gin.SetMode(gin.ReleaseMode)
	S.Engine = gin.New()
	S.Engine.Use(middleware.GinLoggerMiddleware(&S.Logger))
	S.Engine.Use(gin.Recovery())
	S.Engine.Use(middleware.GinCORS())
	RegisterService2Server()
}

func Start() {
	S.Logger.Info().Msg("http://0.0.0.0:8080")
	S.Engine.Run("0.0.0.0:8080")
}