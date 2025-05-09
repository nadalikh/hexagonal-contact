package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	var err error
	config := zap.NewProductionConfig()
	encoder := zap.NewProductionEncoderConfig()
	encoder.TimeKey = "timestamp"
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig = encoder
	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}
func Info(msg string, fields ...zap.Field) {
	log.Info(msg, fields...)
}
func Debug(msg string, fields ...zap.Field) {
	log.Debug(msg, fields...)
}
func Error(msg string, fields ...zap.Field) {
	log.Error(msg, fields...)
}
