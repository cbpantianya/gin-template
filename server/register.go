package server

import (
	"errors"
	"gin-template/v2/utils"
)

type ServiceInfo struct {
	ID   string
	Inst Instance
}

type Instance interface {
	ServiceInfo() *ServiceInfo // 服务信息
	ServiceInit(*Server) error  // 服务初始化
	ServiceStart(*Server) error // 服务开始
	ServiceStop() error        // 服务停止
}

var Services = make(map[string]*ServiceInfo)

func Register(inst Instance) error {

	var se = inst.ServiceInfo()

	// 检查服务信息
	if err := ServiceInfoCheck(se); err != nil {
		return err // 日志
	}

	// 添加前缀，避免重复的Key
	id := utils.AddRandomPrefix(se.ID)
	Services[id] = se

	return nil

}

func RegisterService2Server() {

	for _, ins := range Services {
		if err := ins.Inst.ServiceInit(&S); err != nil {
			S.Logger.Panic().Msg("Failed to init service")
		}
		if err := ins.Inst.ServiceStart(&S); err != nil {
			S.Logger.Panic().Msg("Failed to start service")
		}
		S.Logger.Info().Msg("Successfully start service:" + ins.ID)
	}
}

func ServiceInfoCheck(ins *ServiceInfo) error {
	if ins.ID == "" {
		return errors.New("Failed to register service: ID is empty")
	} else if ins.Inst == nil {
		return errors.New("Failed to register service: Instance is nil")
	}

	return nil
}
