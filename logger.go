package log

import "os"

func Trace(logFields Fields, args ...interface{}) {
    defaultLogger.Trace(logFields, args...)
}

func Info(logFields Fields, args ...interface{}) {
    defaultLogger.Info(logFields, args...)
}

func Error(logFields Fields, args ...interface{}) {
    defaultLogger.Error(logFields, args...)
}

func Warn(logFields Fields, args ...interface{}) {
    defaultLogger.Warn(logFields, args...)
}
func Debug(logFields Fields, args ...interface{}) {
    defaultLogger.Debug(logFields, args...)
}

func SetDefaultLogger(logObj *LoggerObj) {
    defaultLogger = logObj
}

// IsExists 文件是否存在
func IsExists(path string) (os.FileInfo, bool) {
    f, err := os.Stat(path)
    return f, err == nil || os.IsExist(err)
}