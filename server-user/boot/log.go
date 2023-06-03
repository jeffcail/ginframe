package boot

import (
	_uber "github.com/jeffcail/ginframe/server-common/pkg/uber"
	"github.com/jeffcail/ginframe/server-user/ulogger"
)

// InitLog init log
func InitLog(filepath string) {
	logger := _uber.InitLogger(filepath)
	ulogger.SetLogger(logger)

	ulogger.UserLogger.Info("🚀🚀🚀🚀🚀🚀 logger success...🚀🚀🚀🚀🚀🚀")
}
