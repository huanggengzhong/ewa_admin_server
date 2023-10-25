package system

import (
	"github.com/gin-gonic/gin"
	"github.com/huanggengzhong/ewa_admin_server/model/system"
	"github.com/huanggengzhong/ewa_admin_server/utils"
	"net/http"
)

type BaseRouter struct {
}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) gin.IRoutes {
	baseRouter := Router.Group("base")
	{
		baseRouter.POST("login", func(context *gin.Context) {
			context.JSON(http.StatusOK, "登录成功")
		})
		baseRouter.POST("register", func(context *gin.Context) {
			var form system.Register
			if err := context.ShouldBindJSON(&form); err != nil {
				context.JSON(http.StatusOK, gin.H{
					"error": utils.GetErrorMsg(form, err),
				})
				return
			}
			context.JSON(http.StatusOK, "ok")
		})
	}
	return baseRouter
}
