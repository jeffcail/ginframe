// 配置文件加载

package config

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
	"strings"
)

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

//LoadCoreConfig
func LoadCoreConfig(ip string, port int, cfg string, group string, config interface{}) {
	serverConfigs := []constant.ServerConfig{
		{IpAddr: ip, Port: uint64(port)},
	}
	nacosClient, err := clients.NewConfigClient(vo.NacosClientParam{
		ClientConfig:  &constant.ClientConfig{TimeoutMs: 5000},
		ServerConfigs: serverConfigs,
	})
	if err != nil {
		panic(err)
	}
	content, err := nacosClient.GetConfig(vo.ConfigParam{
		DataId: cfg,
		Group:  group,
	})
	if err != nil {
		panic(err)
	}
	viper.SetConfigType("yaml")
	err = viper.ReadConfig(strings.NewReader(content))
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(config)
	if err != nil {
		panic(err)
	}
}
