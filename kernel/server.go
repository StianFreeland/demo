package kernel

import (
	"demo/services/zlog"
	"go.uber.org/zap"
	"net/http"
)

func handleServer() {
	defer wg.Done()
	zlog.Warn(moduleName, "handle server begin")
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		zlog.Error(moduleName, "handle server", zap.Error(err))
		return
	}
	zlog.Warn(moduleName, "handle server end")
}
