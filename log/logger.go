package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type Logger struct {
	logLevel zap.AtomicLevel
	logger   *zap.SugaredLogger
}

func NewLogger(path string, level string) *Logger {
	return newFileLogger(path, level)
}

func (l *Logger) Debug(format string, values ...interface{}) {
	l.logger.Debugf(format, values...)
}

func (l *Logger) Info(format string, values ...interface{}) {
	l.logger.Infof(format, values...)
}

func (l *Logger) Warn(format string, values ...interface{}) {
	l.logger.Warnf(format, values...)
}

func (l *Logger) Error(format string, values ...interface{}) {
	l.logger.Errorf(format, values...)
}

func (l *Logger) Fatal(format string, values ...interface{}) {
	l.logger.Fatalf(format, values...)
}

func (l *Logger) setLevel(level string) {
	switch level {
	case "DEBUG":
		l.logLevel.SetLevel(zapcore.DebugLevel)
	case "INFO":
		l.logLevel.SetLevel(zapcore.InfoLevel)
	case "WARN":
		l.logLevel.SetLevel(zapcore.WarnLevel)
	case "ERROR":
		l.logLevel.SetLevel(zapcore.ErrorLevel)
	case "FATAL":
		l.logLevel.SetLevel(zapcore.FatalLevel)
	default:
		l.logLevel.SetLevel(zapcore.ErrorLevel)
	}
}

func newFileLogger(path string, level string) *Logger {
	var logLevel = zap.NewAtomicLevel()
	writer := zapcore.AddSync(&lumberjack.Logger{
		Filename:  path,
		MaxSize:   1024, //MB
		LocalTime: true,
		Compress:  true,
	})

	core := zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()), writer, logLevel)
	l := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(2))
	logger := &Logger{logger: l.Sugar(), logLevel: logLevel}
	logger.setLevel(level)
	return logger
}

func newStdoutLogger(level string) *Logger {
	var logLevel = zap.NewAtomicLevel()
	writer := zapcore.AddSync(os.Stdout)
	core := zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()), writer, logLevel)
	l := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(2))
	logger := &Logger{logger: l.Sugar(), logLevel: logLevel}
	logger.setLevel(level)
	return logger
}
