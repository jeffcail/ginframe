package ulogger

import "go.uber.org/zap"

var UserLogger *zap.Logger

// SetLogger set logger
func SetLogger(_logger *zap.Logger) {
	UserLogger = _logger
}
