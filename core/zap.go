package core

import (
	"fmt"
	"github.com/huanggengzhong/ewa_admin_server/core/internal"
	"github.com/huanggengzhong/ewa_admin_server/global"
	"github.com/huanggengzhong/ewa_admin_server/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// (logger *zap.Logger)
func InitializeZap() (logger *zap.Logger) {
	isExist, _ := utils.PathExists(global.EWA_CONFIG.Zap.Director)
	if !isExist {
		os.Mkdir(global.EWA_CONFIG.Zap.Director, os.ModePerm) //ModePerm覆盖所有Unix权限位（用于通过&获取类型位
	}
	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))
	if global.EWA_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}

	fmt.Println("---zap log 初始化成功---")
	return logger
	//官方example
	//core := zapcore.NewTee(
	//	zapcore.NewCore(kafkaEncoder, topicErrors, highPriority),
	//	zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
	//	zapcore.NewCore(kafkaEncoder, topicDebugging, lowPriority),
	//	zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority),
	//)
	//
	//// From a zapcore.Core, it's easy to construct a Logger.
	//logger := zap.New(core)
}
