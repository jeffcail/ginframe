package global

import "github.com/jeffcail/ginframe/server-common/nacosRF"

var TestConfig *GlobalConfig

type GlobalConfig struct {
	GinAppDebug string              `json:"gin_app_debug"`
	Http        Http                `json:"http"`
	Grpc        Grpc                `json:"grpc"`
	Nacos       nacosRF.NacosConfig `json:"nacos"`
}

// Http http服务配置
type Http struct {
	BindPort string `json:"bind_port"`
	LogPath  string `json:"log_path"`
}

// Grpc Grpc配置
type Grpc struct {
	Port string `json:"port"`
}
