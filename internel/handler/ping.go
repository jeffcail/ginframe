package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jeffcail/ginframe/internel/input"
	out2 "github.com/jeffcail/ginframe/internel/out"
	"net/http"
)

// Ping test router
func Ping(c *gin.Context) {
	var param input.PingInput
	err := c.Bind(&param)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 1000,
			"msg":    err,
			"data":   nil,
		})
	}

	out := &out2.PingOutput{
		Name:    "test",
		Content: "pong...",
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 2000,
		"msg":    "success",
		"data":   out,
	})
}
