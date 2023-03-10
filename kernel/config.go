package kernel

import (
	"demo/services/zlog"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var port int64

func loadConfig() {
	zlog.Warn(moduleName, "load config begin")
	port = viper.GetInt64("kernel.port")
	if port == 0 {
		zlog.Fatal(moduleName, "load config", zap.Int64("port", port))
	}
	zlog.Warn(moduleName, "load config", zap.Int64("port", port))
	zlog.Warn(moduleName, "load config end")
}
