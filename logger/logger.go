package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	var err error

	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	log, err = config.Build(zap.AddCallerSkip(1))

	if err != nil {
		panic(err)
	}

}

func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

func Warn(message string, fields ...zap.Field) {
	log.Warn(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	log.Error(message, fields...)
}

func Panic(message string, fields ...zap.Field) {
	log.Panic(message, fields...)
}

func Fatal(message string, fields ...zap.Field) {
	log.Fatal(message, fields...)
}
