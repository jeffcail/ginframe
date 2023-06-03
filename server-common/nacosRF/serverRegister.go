package nacosRF

import (
	"fmt"
	"github.com/jeffcail/ginframe/server-common/utils/ip"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"log"
)

func (this *nacosRF) Register(serverName string, port int) {
	ok, err := this.client.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          _ip.GetIp(),
		Port:        uint64(port),
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Metadata:    map[string]string{},
		ClusterName: clusters,
		ServiceName: serverName,
		GroupName:   groupName,
		Ephemeral:   true,
	})
	if err != nil {
		log.Fatal(err)
	}
	if !ok {
		log.Fatal(fmt.Sprintf("注册服务【%s】发生错误", serverName))
	}
}
