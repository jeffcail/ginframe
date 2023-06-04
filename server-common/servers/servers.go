package servers

import (
	"context"
	"fmt"
	_const "github.com/jeffcail/ginframe/server-common/const"
	"github.com/jeffcail/ginframe/server-common/nacosRF"
	"github.com/jeffcail/ginframe/server-user/pb"
	"google.golang.org/grpc"
	"sync"
)

type Servers struct {
	conn *grpc.ClientConn
}

var (
	userMap sync.Map
	rpcUser sync.Once
)

func init() {
	userMap = sync.Map{}
}

// GetUserRpcServer 发现用户服
func GetUserRpcServer(index string) pb.RpcUserServiceClient {
	serverName := fmt.Sprintf(_const.UserServerName + index)
	v, has := userMap.Load(serverName)
	if !has {
		info, err := nacosRF.NacosInstance.FindOneInstance(serverName)
		if err != nil {
			return nil
		}

		//conn, err := grpc.Dial(fmt.Sprintf("%s:%d", info.Ip, info.Port), grpc.WithInsecure())
		conn, err := grpc.DialContext(context.Background(), fmt.Sprintf("%s:%d", info.Ip, info.Port),
			grpc.WithInsecure(), grpc.WithDisableHealthCheck())
		if err != nil {
			return nil
		}
		p := pb.NewRpcUserServiceClient(conn)
		userMap.Store(serverName, p)
		go func() {
			rpcUser.Do(func() {
				nacosRF.NacosInstance.ListenInstance(serverName, func(server *nacosRF.ServerInfo) {
					info, err := nacosRF.NacosInstance.FindOneInstance(serverName)
					if err != nil {
						return
					}
					conn, err := grpc.DialContext(context.Background(), fmt.Sprintf("%s:%s", info.Ip, info.Port),
						grpc.WithInsecure(), grpc.WithDisableHealthCheck())
					if err != nil {
						return
					}
					p = pb.NewRpcUserServiceClient(conn)
					userMap.Store(serverName, p)
				})
			})
		}()
		return p
	}
	return v.(pb.RpcUserServiceClient)
}
