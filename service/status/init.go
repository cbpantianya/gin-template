package status

import (
	"gin-template/v2/server"
)

type Status struct{}

var instance *Status

func init() {
	instance = &Status{}
	err := server.Register(instance)
	if err != nil {
		panic(err)
	}
}

func (p *Status) ServiceInfo() *server.ServiceInfo {
	return &server.ServiceInfo{
		ID:   "Status",
		Inst: instance,
	}
}

func (p *Status) ServiceInit(s *server.Server) error {
	return nil
}

func (p *Status) ServiceStart(s *server.Server) error {
	// 服务状态监测
	s.Engine.GET("/v1/status", Check(s))
	return nil
}

func (p *Status) ServiceStop() error {
	return nil
}
