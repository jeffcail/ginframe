package _viper

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
	"strings"
)

// ParseViperConfig 解析配置文件
func ParseViperConfig(path string, config interface{}) {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(config); err != nil {
		panic(err)
	}
}

// LoadCoreConfig 加载nacos配置
func LoadCoreConfig(ip string, port int, cfg string, group string, config interface{}) {
	serverConfigs := []constant.ServerConfig{
		{IpAddr: ip, Port: uint64(port)},
	}
	nacosClient, err := clients.NewConfigClient(vo.NacosClientParam{
		ClientConfig: &constant.ClientConfig{
			NamespaceId:         "",
			TimeoutMs:           5000,
			NotLoadCacheAtStart: true,
		},
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
