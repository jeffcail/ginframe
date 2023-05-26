package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jeffcail/ginframe/common/global"
	"github.com/jeffcail/ginframe/internel/router"
)

// HttpServe init
func HttpServe() {
	r := gin.Default()
	router.ApiRouter(r)
	r.Run(fmt.Sprintf("%s%s", ":", global.Config.Http.BindPort))
}
