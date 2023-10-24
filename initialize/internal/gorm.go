package internal

import "gorm.io/gorm"

var Gorm = new(_gorm)

type _gorm struct {
}

// Config gorm 自定义配置
func (g *_gorm) Config(prefix string, singular bool) *gorm.Config {
	//将传入的字符串前缀和单复数形式参数应用到 GORM 的命名策略中，并禁用迁移过程中的外键约束，返回最终生成的 GORM 配置信息。

	config := &gorm.Config{
		//命名策略
	}
	return config
}
