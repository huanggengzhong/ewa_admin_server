package initialize

import (
	"github.com/huanggengzhong/ewa_admin_server/global"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	switch global.EWA_CONFIG.App.DbType {
	case "mysql":
		return GormMysql()
		//其它数据库有需要添加
	default:
		return GormMysql()
	}

}
