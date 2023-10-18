package core

import (
	"fmt"
	"github.com/spf13/viper"
)

// 优先级: 命令行 > 环境变量 > 默认值
func InitViper() {
	fmt.Println("haha")
	//vip :=
	viper.New()
}
