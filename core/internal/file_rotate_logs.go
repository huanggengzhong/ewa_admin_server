package internal

import (
	"os"
	"path"
	"time"

	"github.com/huanggengzhong/ewa_admin_server/global"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
)

var FileRotateLogs = new(fileRotateLogs)

type fileRotateLogs struct {
}

func (r fileRotateLogs) GetWriteSyncer(level string) (zapcore.WriteSyncer, error) {
	fileWriter, err := rotatelogs.New(
		path.Join(global.EWA_CONFIG.Zap.Director, "%Y-%m-%d", level+".log"),
		rotatelogs.WithClock(rotatelogs.Local),
		rotatelogs.WithMaxAge(time.Duration(global.EWA_CONFIG.Zap.MaxAge)*24*time.Hour),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if global.EWA_CONFIG.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}
