// init 项目初始化

package cmd

import (
	"flag"
	"fmt"
	"github.com/jeffcail/ginframe/common/global"
	"github.com/jeffcail/ginframe/core/config"
	"log"
	"os"
)

var (
	RootDir string
	err     error
	cf      string
)

func init() {

	RootDir, err = os.Getwd()
	if err != nil {
		panic(err)
	}

	appConf := fmt.Sprintf("%s%s", RootDir, "/application.yml")
	global.NewApplicationConfig()
	config.ParseConfig(appConf, &global.AppConfig)

	switch global.AppConfig.ConfigRemote {
	case false:
		cf = fmt.Sprintf("%s%s", RootDir, "/config.yml")
		global.NewGoAppConfig()
		config.ParseConfig(cf, &global.Config)
		break
	default:
		loadRemoteConfig()
		break
	}

}

var (
	ip    = flag.String("ip", "ip", "The nacos of ip address")
	port  = flag.Int("p", 0, "The nacos of port")
	cfg   = flag.String("c", "default", "The nacos of Data ID")
	group = flag.String("g", "default", "The nacos of Group")
)

func loadRemoteConfig() {
	flag.Parse()
	config.LoadCoreConfig(*ip, *port, *cfg, *group, global.Config)
}

// Init 项目初始化
func Init() {
	errs := make(chan error)
	go func() {
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

	select {
	case err = <-errs:
		log.Fatalf("Run server err: %v")
	}
}
