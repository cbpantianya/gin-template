package utils

import "github.com/BurntSushi/toml"

type Config struct {
	// HTTP配置
	HTTPServer struct {
		Host             string   // 主机地址
		Port             int      // 端口
	}
	// 数据库配置
	MySQL struct {
		Host     string
		Port     int
		User     string
		Password string
		Database string
	}
	// Redis配置
	Redis struct {
		Host     string
		Port     int
		Password string
		Database int
	}
}

var GConfig Config

func CFInit() {
	// 解析配置文件
	_, err := toml.DecodeFile("./config/config.toml", &GConfig)
	if err != nil {
		// 此时还未初始化日志服务，所以直接输出到控制台
		panic("Failed to parse config file!")
	}
}