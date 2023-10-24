package initialize

import (
	"github.com/huanggengzhong/ewa_admin_server/global"
	"github.com/huanggengzhong/ewa_admin_server/initialize/internal"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GormMysql() {
	m := global.EWA_CONFIG.MySQL
	gorm.Open(mysql.New(mysql.Config{
		DSN:                       m.Dsn(),
		DefaultStringSize:         191,   //string 类型字段的默认长度
		SkipInitializeWithVersion: false, //根据版本自动配置
	}), internal.Gorm.Config(m.Prefix, m.Singular))
}
