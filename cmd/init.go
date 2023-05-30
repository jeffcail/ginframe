// init 项目初始化

package cmd

import (
	"flag"
	"fmt"
	"github.com/jeffcail/ginframe/common/global"
	_process "github.com/jeffcail/ginframe/common/process"
	"github.com/jeffcail/ginframe/core/config"
	_rpc "github.com/jeffcail/ginframe/internel/admin/rpc"
	"log"
	"os"
)

var (
	RootDir string
	err     error
	cf      string
)

var (
	ip    = flag.String("ip", "127.0.0.1", "The nacos of ip address")
	port  = flag.Int("p", 7848, "The nacos of port")
	cfg   = flag.String("c", "gin-frame.yml", "The nacos of Data ID")
	group = flag.String("g", "ginframe", "The nacos of Group")
)

func init() {

	RootDir, err = os.Getwd()
	if err != nil {
		panic(err)
	}

	appConf := fmt.Sprintf("%s%s", RootDir, "/application.yml")
	global.NewApplicationConfig()
	config.ParseConfig(appConf, global.AppConfig)

	if global.AppConfig.IsEnableGOMAXPROCS {
		_process.GroRuntimeMaxCpu()
	}

	switch global.AppConfig.ConfigRemote {
	case false:
		cf = fmt.Sprintf("%s%s", RootDir, "/config.yml")
		global.NewGoAppConfig()
		config.ParseConfig(cf, &global.Config)
		break
	default:
		flag.Parse()
		loadRemoteConfig(*ip, *port, *cfg, *group, &global.Config)
		break
	}
}

func loadRemoteConfig(ip string, port int, cfg string, group string, configs interface{}) {
	config.LoadCoreConfig(ip, port, cfg, group, configs)
}

// Init 项目初始化
func Init() {
	errs := make(chan error)
	go func() {
		InitDb()
		InitLog()
		InitEs()
		err = HttpServe()
		if err != nil {
			errs <- err
		}
	}()

	go func() {
		err = WebsocketServer()
		if err != nil {
			errs <- err
		}
	}()

	go func() {
		err = _rpc.RpcServer()
		if err != nil {
			errs <- err
		}
	}()

	select {
	case err = <-errs:
		log.Fatalf("Run server err: %v", err)
	}
}
