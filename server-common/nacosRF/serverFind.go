package nacosRF

import (
	"github.com/nacos-group/nacos-sdk-go/model"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

type ServerInfo struct {
	Ip         string
	Port       uint64
	ServerName string
}

// FindOneInstance 获取某个服务连接信息
func (this *nacosRF) FindOneInstance(serverName string) (*ServerInfo, error) {
	instance, err := this.client.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		Clusters:    []string{clusters},
		ServiceName: serverName,
		GroupName:   groupName,
	})
	if err != nil {
		return nil, err
	}
	return &ServerInfo{
		Ip:         instance.Ip,
		Port:       instance.Port,
		ServerName: instance.ServiceName,
	}, nil
}

// FindAllInstance 获取全部服务连接信息
func (this *nacosRF) FindAllInstance(serverName string) ([]*ServerInfo, error) {
	ins, err := this.client.SelectInstances(vo.SelectInstancesParam{
		Clusters:    []string{clusters},
		ServiceName: serverName,
		GroupName:   groupName,
		HealthyOnly: true,
	})
	if err != nil {
		return nil, err
	}
	infos := make([]*ServerInfo, 0)
	for _, info := range ins {
		infos = append(infos, &ServerInfo{
			Ip:         info.Ip,
			Port:       info.Port,
			ServerName: info.ServiceName,
		})
	}
	return infos, nil
}

// ListenInstance 监听某个服务状态信息
func (this *nacosRF) ListenInstance(serverName string, callback func(server *ServerInfo)) {
	this.client.Subscribe(&vo.SubscribeParam{
		ServiceName: serverName,
		Clusters:    []string{clusters},
		GroupName:   groupName,
		SubscribeCallback: func(services []model.SubscribeService, err error) {
			for _, service := range services {
				if service.Enable && service.Healthy {
					callback(&ServerInfo{
						Ip:         service.Ip,
						Port:       service.Port,
						ServerName: service.ServiceName,
					})
				}
			}
		},
	})
}
