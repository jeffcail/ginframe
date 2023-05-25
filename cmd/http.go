package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jeffcail/ginframe/common/global"
)

// HttpServe init
func HttpServe() {
	r := gin.Default()

	r.Run(fmt.Sprintf("%s%s", ":", global.Config.Http.BindPort))
}
