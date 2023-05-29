package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jeffcail/ginframe/common/global"
	"github.com/jeffcail/ginframe/internel/admin/router"
)

// HttpServe init
func HttpServe() error {
	r := gin.Default()
	router.ApiRouter(r)
	return r.Run(fmt.Sprintf("%s%s", ":", global.Config.Http.BindPort))
}
