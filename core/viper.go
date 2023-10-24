package core

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/huanggengzhong/ewa_admin_server/core/internal"
	"github.com/huanggengzhong/ewa_admin_server/global"
	"github.com/spf13/viper"
	"os"
)

// 优先级: 项目目录文件>命令行 > bash环境变量 > gin默认值
func InitializeViper(path ...string) *viper.Viper {

	var config string

	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "开始StringVar绑定config")
		flag.Parse()
		//fmt.Println("正在获取命令行c的值:", config)

		if config == "" {
			/*
						判断 internal.ConfigEnv 常量存储的环境变量是否为空
						比如我们启动项目的时候，先执行：
				export EWA_CONFIG="config.yaml"
						再启动项目go run main.go
						这时候 os.Getenv(internal.ConfigEnv) 得到的就是 config.yaml
						当然，也可以通过 os.Setenv(internal.ConfigEnv, "config.yaml") 在初始化之前设置
			*/
			configEnv := os.Getenv(internal.ConfigEnv)
			if configEnv == "" {
				//fmt.Println("使用gin环境配置")
				//todo,使用gin环境配置
				switch gin.Mode() {
				case gin.DebugMode:
					config = internal.ConfigDebugFile
				case gin.TestMode:
					config = internal.ConfigTestFile
				case gin.ReleaseMode:
					config = internal.ConfigReleaseFile
				}
			} else {
				//使用bash环境设置的值
				config = configEnv
			}
			//fmt.Println("configEnv:", configEnv)
		} else {
			//使用命令行路径,命令行正确自动赋值到config
		}
	} else {
		//使用目录yaml文件
		config = path[0]
		fmt.Println()
	}
	vip := viper.New()
	//fmt.Println("config文件路径是:", config)
	vip.SetConfigFile(config)
	vip.SetConfigType("yaml")
	err := vip.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("读取配置文件失败,失败信息:%s\n", err))
	}
	//fmt.Printf("读取值:%v", vip.Get("app"))
	vip.WatchConfig()
	vip.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置文件变化", e.Name)
		if err = vip.Unmarshal(&global.EWA_CONFIG); err != nil {
			fmt.Println("配置文件更换序列化失败", err)
		}
	})
	//Unmarshal作用绑定在EWA_CONFIG上,使global.EWA_CONFIG.App可以获取想要的信息
	if err = vip.Unmarshal(&global.EWA_CONFIG); err != nil {
		fmt.Println("配置文件初始化化失败", err)
	}
	fmt.Printf("----viper配置初始化成功----")
	return vip
}
