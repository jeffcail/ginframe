package nacosRF

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"go.uber.org/zap"
	"log"
)

const (
	clusters  = "GIN_FRAME"
	groupName = "GIN_FRAME_GROUP"
)

type WarningConfig struct {
	RoomId int
	Url    string
}

type NacosConfig struct {
	Addr      string
	Port      int
	LogPath   string
	CachePath string
}

type nacosRF struct {
	client naming_client.INamingClient
}

var NacosInstance *nacosRF

// InitNacos init Nacos
func InitNacos(config NacosConfig) *nacosRF {
	//return &nacosRF{}
	d := &nacosRF{}
	clientConfig := constant.ClientConfig{
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              config.LogPath,
		CacheDir:            config.CachePath,
		LogLevel:            "debug",
	}

	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      config.Addr,
			ContextPath: "/nacos",
			Port:        uint64(config.Port),
			Scheme:      "http",
		},
	}

	var err error
	d.client, err = clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		log.Fatal("nacos失败了,err:%v", zap.Error(err))
	}

	NacosInstance = d
	return d
}
