package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

func SetupLogger() *zap.Logger {
	atomLv := zap.NewAtomicLevel()
	atomLv.SetLevel(zapcore.InfoLevel)

	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
	jsonEncoder := zapcore.NewJSONEncoder(cfg)

	consoleWriter := zapcore.AddSync(os.Stdout)

	fileWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./logs/workflow.log",
		MaxSize:    100,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
	})

	consoleCore := zapcore.NewCore(
		jsonEncoder,
		consoleWriter,
		atomLv,
	)

	fileCore := zapcore.NewCore(
		jsonEncoder,
		fileWriter,
		atomLv,
	)

	loggerCore := zapcore.NewTee(consoleCore, fileCore)

	return zap.New(loggerCore)
}
