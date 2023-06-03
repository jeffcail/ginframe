// 配置文件加载

package config

import (
	_viper "github.com/jeffcail/ginframe/server-common/pkg/viper"
)

// ParseConfig 配置文件加载
func ParseConfig(path string, config interface{}) {
	_viper.ParseViperConfig(path, config)
}

// LoadCoreConfig 加载配置
func LoadCoreConfig(ip string, port int, cfg string, group string, config interface{}) {
	_viper.LoadCoreConfig(ip, port, cfg, group, config)
}
