package _rpc

import (
	"github.com/jeffcail/ginframe/server-user/global"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

// RpcServer rpc server
func RpcServer() error {

	rpc.Register(new(Arith))

	rpc.HandleHTTP()
	addr := global.Config.Rpc.Listener
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("fatal error: ", err)
	}
	return http.Serve(l, nil)
}
