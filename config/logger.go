package config

import (
	"context"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm/logger"
)

func NewLogger() *zap.Logger {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "timestamp"

	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:       false,
		DisableCaller:     true,
		DisableStacktrace: true,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig:     encoderConfig,
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stderr"},
	}
	logger, err := config.Build()
	if err != nil {
		panic(err)
	}
	return logger
}

func NewDatabaseLogger(logger *zap.Logger) *DatabaseLogger {
	return &DatabaseLogger{logger: logger}
}

type DatabaseLogger struct {
	logger *zap.Logger
}

func (g *DatabaseLogger) LogMode(logger.LogLevel) logger.Interface {
	return g
}
func (g *DatabaseLogger) Info(_ context.Context, msg string, data ...interface{}) {
	g.logger.Sugar().Infof(msg, data)
}
func (g *DatabaseLogger) Warn(_ context.Context, msg string, data ...interface{}) {
	g.logger.Sugar().Warnf(msg, data)
}
func (g *DatabaseLogger) Error(_ context.Context, msg string, data ...interface{}) {
	g.logger.Sugar().Errorf(msg, data)
}
func (g *DatabaseLogger) Trace(_ context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin).String()
	sql, rows := fc()
	fields := []zapcore.Field{
		zap.Int64("rows", rows),
		zap.String("elapsed", elapsed),
	}
	log := g.logger.With(fields...)
	if log.Sugar().Infof(sql); err != nil {
		log.Sugar().Errorf(err.Error())
	}
}
