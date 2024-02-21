package logger

import (
	"context"
	"log"
	"os"

	"go.uber.org/zap"
)

var defaultLogger Logger

type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})

	Debugf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Errorf(template string, args ...interface{})

	Logger() *zap.Logger
	WithOptionsAddCallerSkip(skip int) Logger
	Sync()

	WithLogger(ctx context.Context) Logger
}

type Options struct {
	FilePath   string
	Level      string
	Format     string
	ProdMode   bool
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

func GetLogger() Logger {
	return defaultLogger
}

func SetDefaultLogger(logger Logger) {
	defaultLogger = logger
}

func LogErrorAndExit(err error, appLogger ...Logger) {
	if len(appLogger) == 1 {
		appLogger[0].Error(err)
	} else {
		log.Println(err)
	}

	os.Exit(1)
}

type Key string

const loggerFieldKey = Key("loggerFieldKey")

func WithValue(ctx context.Context, fields map[string]string) context.Context {
	if len(fields) == 0 {
		return ctx
	}
	var zapFields []zap.Field
	for k, v := range fields {
		zapFields = append(zapFields, zap.String(k, v))
	}
	return context.WithValue(ctx, loggerFieldKey, zapFields)
}

func GetFields(ctx context.Context) []zap.Field {
	zapFields, ok := ctx.Value(loggerFieldKey).([]zap.Field)
	if ok {
		return zapFields
	}
	return nil
}

func GetField(ctx context.Context, key string) *zap.Field {
	loggerFields := GetFields(ctx)
	if loggerFields != nil {
		return nil
	}

	for _, field := range loggerFields {
		if field.Key == key {
			return &field
		}
	}
	return nil
}
