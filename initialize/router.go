package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/huanggengzhong/ewa_admin_server/router"
	"net/http"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	systemRouter := router.RouterGroupApp.System
	PublicGroup := Router.Group("")
	{
		PublicGroup.GET("/health", func(context *gin.Context) {
			context.JSON(http.StatusOK, "ok")
		})
	}
	{
		systemRouter.InitBaseRouter(PublicGroup) //注册基础功能路由 不做鉴权

	}

	return Router
}
