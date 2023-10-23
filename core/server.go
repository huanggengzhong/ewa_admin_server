package core

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/huanggengzhong/ewa_admin_server/global"
	"net/http"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	address := fmt.Sprintf(":%d", global.EWA_CONFIG.App.Port)
	s := initServer(address, r)
	// 保证文本顺序输出
	time.Sleep(10 * time.Millisecond)
	fmt.Println("运行端口:", address)
	s.ListenAndServe()

}

func initServer(address string, router *gin.Engine) server {
	// 使用endless库创建一个HTTP服务器，其中address是服务器的监听地址（如:8080）
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 20 * time.Second
	s.WriteTimeout = 20 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}
