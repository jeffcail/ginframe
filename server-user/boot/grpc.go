package boot

import (
	"fmt"
	"github.com/jeffcail/ginframe/server-common/const"
	"github.com/jeffcail/ginframe/server-user/global"
	"github.com/jeffcail/ginframe/server-user/grpcservices/services"
	"github.com/jeffcail/ginframe/server-user/pb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

// RunGrpcServer 运行grpc 服务
func RunGrpcServer() error {
	s := grpc.NewServer()
	// 注册心跳
	healthServer := health.NewServer()
	healthServer.SetServingStatus(_const.HEALTHCHECK_SERVICE, healthpb.HealthCheckResponse_SERVING)
	healthpb.RegisterHealthServer(s, healthServer)

	// 注册服务
	pb.RegisterRpcUserServiceServer(s, new(services.RpcUserService))

	reflection.Register(s)
	fmt.Println(global.Config.Grpc.Port)
	listen, err := net.Listen("tcp", ":"+global.Config.Grpc.Port)
	if err != nil {
		log.Fatalf("Router: Start grpc failed, err: %v", zap.Error(err))
	}
	return s.Serve(listen)
}
