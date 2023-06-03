package main

import (
	"context"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/jeffcail/ginframe/server-common/config"
	"github.com/jeffcail/ginframe/server-common/nacosRF"
	"github.com/jeffcail/ginframe/server-common/servers"
	"github.com/jeffcail/ginframe/server-test/global"
	"github.com/jeffcail/ginframe/server-user/pb"
	"github.com/spf13/cast"
	"net/http"
)

var (
	ip    = flag.String("ip", "127.0.0.1", "The nacos of ip address")
	port  = flag.Int("p", 7848, "The nacos of port")
	cfg   = flag.String("c", "server-test.yml", "The nacos of Data ID")
	group = flag.String("g", "gin-frame", "The nacos of Group")
)

func init() {
	flag.Parse()
	loadRemoteConfig(*ip, *port, *cfg, *group, &global.TestConfig)
	nacosRF.InitNacos(global.TestConfig.Nacos)
}

func loadRemoteConfig(ip string, port int, cfg string, group string, configs interface{}) {
	config.LoadCoreConfig(ip, port, cfg, group, configs)
}

func main() {
	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		userServiceClient := servers.GetUserRpcServer("1")
		if userServiceClient == nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 400,
				"msg":  "没有发现user服务",
			})
			return
		}
		// 用户表id 为 1
		info, err := userServiceClient.GetUserInfo(context.Background(), &pb.GetUserInfoRequest{Id: "1"})
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 400,
				"msg":  cast.ToString(err),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": info,
		})
	})

	r.Run(":9999")
}
