package ping

import (
	"gin-template/v2/server"
	"gin-template/v2/utils"
	"github.com/gin-gonic/gin"
)

type ping struct {}

var instance *ping

func init() {
	instance = &ping{}
	server.Register(instance)
}

func (p *ping)ServiceInfo() *server.ServiceInfo {
	return &server.ServiceInfo{
        ID: "PING",
		Inst: instance,
    }
}

func (p* ping)ServiceInit(s *server.Server) error {
	return nil
}

func (p* ping)ServiceStart(s *server.Server) error {
	s.Engine.GET("/ping",func (c *gin.Context)  {
		c.JSON(utils.SuccessResp("pong"))
	})
	return nil
}

func (p* ping)ServiceStop() error {
	return nil
}

