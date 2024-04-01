package pkg

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhaoxuat/bazel/pkg/util"
)

type service struct {
	engine *gin.Engine
}

func NewService() *service {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": util.Pong,
		})
	})
	return &service{
		engine: r,
	}
}

func (s *service) Run() (err error) {
	return s.engine.Run()
}
