package initialize

import (
	"fmt"
	"github.com/huanggengzhong/ewa_admin_server/global"
	"github.com/huanggengzhong/ewa_admin_server/initialize/internal"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GormMysql() *gorm.DB {
	m := global.EWA_CONFIG.MySQL
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       m.Dsn(),
		DefaultStringSize:         191,   //string 类型字段的默认长度
		SkipInitializeWithVersion: false, //根据版本自动配置
	}), internal.Gorm.Config(m.Prefix, m.Singular))

	if err != nil {
		return nil
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		fmt.Println("====3-gorm====: gorm link mysql success")
		return db
	}
}
