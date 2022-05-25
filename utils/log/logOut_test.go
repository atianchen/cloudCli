package log

import (
	"os"
	"testing"
)

/**
 * 字符串类型日志输出
 * 日志默认info级别，debug日志不会输出
 */
func TestInfo(t *testing.T) {
	var logger LogInit = &Log{}
	logger.Init()
	Debug("log out debug string")
	Info("log out info string")
}

/**
 * 格式化日志输出
 */
func TestInfof(t *testing.T) {
	var logger LogInit = &Log{}
	logger.Init()
	hostname, _ := os.Hostname()
	Infof("log out: %s", hostname)
}

/**
 * key和value类型输出，支持多个kv
 * log out: {"hostname": "LAPTOP-CH2HG70S"}
 */
func TestInfoKv(t *testing.T) {
	var logger LogInit = &Log{}
	logger.Init()
	hostname, _ := os.Hostname()
	InfoKv("log out: ", "hostname", hostname)
}
