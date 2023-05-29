package _rpc

import (
	"fmt"
	"github.com/jeffcail/ginframe/common/global"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
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

	fmt.Println(os.Stdout, "%s", "start connection")
	return http.Serve(l, nil)
}
