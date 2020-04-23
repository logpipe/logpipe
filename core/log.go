package core

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var levelMap = make(map[string]int)

func init() {
	levelMap["DEBUG"] = 1
	levelMap["INFO"] = 2
	levelMap["WARN"] = 3
	levelMap["ERROR"] = 4
}

type Logger struct {
	logger *zap.SugaredLogger
}

var logLevel = zap.NewAtomicLevel()

func NewLogger(path string, level string) *Logger {

	setLevel(level)
	writer := zapcore.AddSync(&lumberjack.Logger{
		Filename:  path,
		MaxSize:   1024, //MB
		LocalTime: true,
		Compress:  true,
	})

	core := zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()), writer, logLevel)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	return &Logger{logger: logger.Sugar()}
}

func (l *Logger) Debug(format string, values ...interface{}) {
	l.logger.Debugf(format, values)
}

func (l *Logger) Info(format string, values ...interface{}) {
	l.logger.Infof(format, values)
}

func (l *Logger) Warn(format string, values ...interface{}) {
	l.logger.Warnf(format, values)
}

func (l *Logger) Error(format string, values ...interface{}) {
	l.logger.Errorf(format, values)
}

func setLevel(level string) {
	switch level {
	case "DEBUG":
		logLevel.SetLevel(zapcore.DebugLevel)
	case "INFO":
		logLevel.SetLevel(zapcore.InfoLevel)
	case "WARN":
		logLevel.SetLevel(zapcore.WarnLevel)
	case "ERROR":
		logLevel.SetLevel(zapcore.ErrorLevel)
	case "FATAL":
		logLevel.SetLevel(zapcore.FatalLevel)
	default:
		logLevel.SetLevel(zapcore.ErrorLevel)
	}
}
