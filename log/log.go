package log

import (
	"sync"
)

var levelMap = make(map[string]int)
var appLogger *Logger
var once = sync.Once{}

func init() {
	levelMap["DEBUG"] = 1
	levelMap["INFO"] = 2
	levelMap["WARN"] = 3
	levelMap["ERROR"] = 4

	appLogger = newStdoutLogger("DEBUG")
}

func InitAppLogger(path string, level string) {
	once.Do(func() {
		appLogger = NewLogger(path, level)
	})
}

func Debug(format string, values ...interface{}) {
	appLogger.Debug(format, values...)
}

func Info(format string, values ...interface{}) {
	appLogger.Info(format, values...)
}

func Warn(format string, values ...interface{}) {
	appLogger.Warn(format, values...)
}

func Error(format string, values ...interface{}) {
	appLogger.Error(format, values...)
}

func Fatal(format string, values ...interface{}) {
	appLogger.Fatal(format, values...)
}
