package _uber

import (
	"github.com/robfig/cron"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

// InitLogger 初始化logger日志
func InitLogger(filepath string) *zap.Logger {
	encoder := getEncoder()
	writeSyncer := getLogWriter(filepath)
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	consoleDebug := zapcore.Lock(os.Stdout)
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	p := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.DebugLevel
	})
	var allCode []zapcore.Core
	allCode = append(allCode, core)
	allCode = append(allCode, zapcore.NewCore(consoleEncoder, consoleDebug, p))
	c := zapcore.NewTee(allCode...)
	return zap.New(c, zap.AddCaller())
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(filepath string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:  filepath,
		MaxSize:   10240,
		MaxAge:    7,
		Compress:  true,
		LocalTime: true,
	}

	c := cron.New()
	c.AddFunc("0 0 0 1/1 * ?", func() {
		lumberJackLogger.Rotate()
	})
	c.Start()
	return zapcore.AddSync(lumberJackLogger)
}
