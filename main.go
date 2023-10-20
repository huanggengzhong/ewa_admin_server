package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/huanggengzhong/ewa_admin_server/core"
	"github.com/huanggengzhong/ewa_admin_server/global"
)

const AppMode = "debug" //运行环境，主要有三种：debug、test、release

func main() {
	gin.SetMode(AppMode)
	// TODO：1.配置初始化
	global.EWA_VIPER = core.InitiallizeViper()
	//配置使用案例:
	fmt.Println("打印配置", global.EWA_CONFIG.App)
	// TODO：2.日志
	// TODO：3.数据库连接
	// TODO：4.其他初始化
	// TODO：5.启动服务

}
