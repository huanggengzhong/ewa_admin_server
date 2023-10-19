package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitializeViper(path ...string) *viper.Viper {
	v := viper.New()

	// 设置默认配置
	v.SetDefault("port2", 8080)
	v.SetDefault("debug", false)

	// 从传入的配置文件路径加载配置
	for _, p := range path {
		fmt.Println("p值:", p)
		v.SetConfigFile(p)
		if err := v.ReadInConfig(); err != nil {
			fmt.Printf("无法加载配置文件 %s: %s\n", p, err)
		}
	}

	return v
}

func main() {
	v := InitializeViper("./test/config_test.json", "./test/config_test2.json")

	// 从 Viper 实例中获取配置值
	port1 := v.GetInt("port1")
	port2 := v.GetInt("port2")

	test := v.GetString("test1")
	debug := v.GetBool("debug")

	fmt.Printf("端口号1：%d\n", port1)
	fmt.Printf("端口号2：%d\n", port2)

	fmt.Printf("test值：%v\n", test)

	fmt.Printf("调试模式：%v\n", debug)
}
