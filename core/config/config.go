// 配置文件加载

package config

import "github.com/spf13/viper"

// ParseConfig 配置文件加载
func ParseConfig(path string, config interface{}) {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(config); err != nil {
		panic(err)
	}
}
