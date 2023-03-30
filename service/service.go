package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const Addr = "localhost:8089"

type Service struct {
	// 服务名称
}

func (s *Service) Start() {
	r := gin.Default()
	r.GET("/user/list", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "ok", "data": "user list"})
	})
	_ = r.Run(Addr)

}
