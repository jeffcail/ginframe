// init 项目初始化

package cmd

import (
	"fmt"
	"github.com/jeffcail/gin-app/common/global"
	"github.com/jeffcail/gin-app/core/config"
	"os"
)

var (
	ROOTDIR string
	err     error
	cf      string
)

func init() {

	ROOTDIR, err = os.Getwd()
	if err != nil {
		panic(err)
	}

	cf = fmt.Sprintf("%s%s", ROOTDIR, "\\gin-app.yaml")

	global.NewGoAppConfig()

	config.ParseConfig(cf, global.Config)
	fmt.Println(global.Config.GinAppName)
	fmt.Println(global.Config.Http.BindPort)
}

// Init 项目初始化
func Init() {

}
