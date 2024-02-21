package zap

import (
	"backend/internal/core/logger"
	"context"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type zapLogger struct {
	logger *zap.Logger
}

func New(logger *zap.Logger) logger.Logger {
	return &zapLogger{
		logger: logger,
	}
}

// https://nyogjtrc.github.io/posts/2019/09/log-rotate-with-zap-logger/
func NewLogger(option *logger.Options) logger.Logger {
	var ioWriter = &lumberjack.Logger{
		Filename:  option.FilePath,
		MaxSize:   10,
		MaxAge:    30,
		LocalTime: true,
		Compress:  option.Compress,
	}

	if option.MaxSize != 0 {
		ioWriter.MaxSize = option.MaxSize
	}
	if option.MaxBackups != 0 {
		ioWriter.MaxBackups = option.MaxBackups
	}
	if option.MaxAge != 0 {
		ioWriter.MaxAge = option.MaxAge
	}

	writeFile := zapcore.AddSync(ioWriter)
	writeStdout := zapcore.AddSync(os.Stdout)

	if option.ProdMode {
		encoderConfig := zap.NewProductionEncoderConfig()
		encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)

		logLevel := parseLogLevel(option.Level, zap.InfoLevel)
		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			zapcore.NewMultiWriteSyncer(writeFile, writeStdout),
			logLevel,
		)
		zapLog := &zapLogger{
			logger: zap.New(
				core,
				zap.AddCaller(),
				zap.AddCallerSkip(1),
				zap.AddStacktrace(zap.ErrorLevel),
			),
		}
		logger.SetDefaultLogger(zapLog)
		return zapLog
	} else {
		encoderConfig := zap.NewDevelopmentEncoderConfig()
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)

		logLevel := parseLogLevel(option.Level, zap.DebugLevel)
		core := zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			zapcore.NewMultiWriteSyncer(writeFile, writeStdout),
			logLevel,
		)
		zapLog := &zapLogger{
			logger: zap.New(
				core,
				zap.AddCaller(),
				zap.AddCallerSkip(1),
				zap.AddStacktrace(zap.ErrorLevel),
			),
		}
		logger.SetDefaultLogger(zapLog)
		return zapLog
	}
}

func parseLogLevel(textLogLevel string, defaultLogLevel zapcore.Level) zapcore.Level {
	logLevel, err := zapcore.ParseLevel(textLogLevel)
	if err != nil {
		return defaultLogLevel
	}
	return logLevel
}

func (l *zapLogger) WithOptionsAddCallerSkip(skip int) logger.Logger {
	return New(l.logger.WithOptions(zap.AddCallerSkip(skip)))
}

func (l *zapLogger) Logger() *zap.Logger {
	return l.logger
}

func (l *zapLogger) WithLogger(ctx context.Context) logger.Logger {
	zapFields := logger.GetFields(ctx)
	if len(zapFields) == 0 {
		return New(l.logger)
	}

	return New(l.logger.With(zapFields...))
}

func (l *zapLogger) Debug(args ...interface{}) {
	l.logger.Sugar().Debug(args)
}

func (l *zapLogger) Info(args ...interface{}) {
	l.logger.Sugar().Info(args)
}

func (l *zapLogger) Warn(args ...interface{}) {
	l.logger.Sugar().Warn(args)
}

func (l *zapLogger) Error(args ...interface{}) {
	l.logger.Sugar().Error(args)
}

func (l *zapLogger) Debugf(template string, args ...interface{}) {
	l.logger.Sugar().Debugf(template, args...)
}

func (l *zapLogger) Infof(template string, args ...interface{}) {
	l.logger.Sugar().Infof(template, args...)
}

func (l *zapLogger) Warnf(template string, args ...interface{}) {
	l.logger.Sugar().Warnf(template, args...)
}

func (l *zapLogger) Errorf(template string, args ...interface{}) {
	l.logger.Sugar().Errorf(template, args...)
}

func (l *zapLogger) Sync() {
	_ = l.logger.Sync()
}
