package global

import (
	"github.com/huanggengzhong/ewa_admin_server/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	EWA_CONFIG config.Configuration
	EWA_VIPER  *viper.Viper
	EWA_LOG    *zap.Logger
)
