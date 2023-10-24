package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/huanggengzhong/ewa_admin_server/core"
	"github.com/huanggengzhong/ewa_admin_server/global"
	"go.uber.org/zap"
)

const AppMode = "debug" //运行环境，主要有三种：debug、test、release

func main() {
	gin.SetMode(AppMode)
	// TODO：1.配置初始化
	global.EWA_VIPER = core.InitializeViper()
	//配置使用案例:
	fmt.Println("打印配置", global.EWA_CONFIG.App)
	// TODO：2.日志
	global.EWA_LOG = core.InitializeZap()
	zap.ReplaceGlobals(global.EWA_LOG)
	global.EWA_LOG.Info("服务启动成功!", zap.String("测试键", "内容哈哈哈哈"))
	// TODO：3.数据库连接
	// TODO：4.其他初始化
	// TODO：5.启动服务
	core.RunServer()

}
