package logger

import (
	"github.com/jeffcail/ginframe/common/global"
	_uber "github.com/jeffcail/ginframe/pkg/uber"
	"go.uber.org/zap"
)

var GinLogger *zap.Logger

func InitLogger() {
	GinLogger = _uber.InitLogger(global.Config.Http.LogPath)
}
