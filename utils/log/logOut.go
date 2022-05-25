package log

import (
	"io"
)

type LogOut interface {
	Instan() io.Writer
}

/**
 * 为了调用简便，这里重新对logger封装
 */
func Debug(s ...interface{}) {
	logger.Debug(s)
}
func Info(s ...interface{}) {
	logger.Info(s...)
}
func Warn(s ...interface{}) {
	logger.Warn(s)
}
func Error(s ...interface{}) {
	logger.Error(s)
}
func Panic(s ...interface{}) {
	logger.Panic(s)
}

/**
 * 格式化输出
 */
func Debugf(s string, args ...interface{}) {
	logger.Debugf(s, args...)
}
func Infof(s string, args ...interface{}) {
	logger.Infof(s, args...)
}
func Warnf(s string, args ...interface{}) {
	logger.Warnf(s, args...)
}
func Errorf(s string, args ...interface{}) {
	logger.Errorf(s, args...)
}
func Panicf(s string, args ...interface{}) {
	logger.Panicf(s, args...)
}

/**
 * 输出:  string: {key: value} 的形式
 */
func DebugKv(s string, keysAndValues ...interface{}) {
	logger.Debugw(s, keysAndValues...)
}
func InfoKv(s string, keysAndValues ...interface{}) {
	logger.Infow(s, keysAndValues...)
}
func WarnKv(s string, keysAndValues ...interface{}) {
	logger.Warnw(s, keysAndValues...)
}
func ErrorKv(s string, keysAndValues ...interface{}) {
	logger.Errorw(s, keysAndValues...)
}
func PanicKv(s string, keysAndValues ...interface{}) {
	logger.Panicw(s, keysAndValues...)
}
