package services

import (
	"context"
	"github.com/jeffcail/ginframe/server-common/utils/wtime"
	"github.com/jeffcail/ginframe/server-user/grpcservices/impl"
	"github.com/jeffcail/ginframe/server-user/pb"
	"github.com/spf13/cast"
)

type RpcUserService struct {
	pb.UnimplementedRpcUserServiceServer
}

// GetUserInfo 获取管理员账号信息
func (ras *RpcUserService) GetUserInfo(ctx context.Context, r *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error) {
	admin := impl.GetUserInfo(r.Id)
	return &pb.GetUserInfoResponse{
		Id:        cast.ToInt64(r.Id),
		Username:  admin.Username,
		Nickname:  admin.Nickname,
		Phone:     admin.Phone,
		Email:     admin.Email,
		CreatedAt: wtime.WTtime.FormatTime(admin.CreatedAt),
	}, nil
}
