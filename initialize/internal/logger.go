package internal

import (
	"fmt"
	"github.com/huanggengzhong/ewa_admin_server/global"
	"gorm.io/gorm/logger"
)

type writer struct {
	logger.Writer
}

func NewWriter(w logger.Writer) *writer {
	return &writer{
		w,
	}
}

func (w *writer) Printf(message string, data ...interface{}) {
	var logZap bool
	switch global.EWA_CONFIG.App.DbType {
	case "mysql":
		logZap = global.EWA_CONFIG.MySQL.LogZap
	}
	if logZap {
		global.EWA_LOG.Info(fmt.Sprintf(message+"\n", data))
	} else {
		w.Writer.Printf(message, data...)
	}
}
