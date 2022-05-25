package log

import (
	"github.com/natefinch/lumberjack"
	"io"
	"os"
)

func (c *ConsoleConfig) Instan() io.Writer {
	return os.Stdout
}

func (f *FileConfig) Instan() io.Writer {
	hook := lumberjack.Logger{
		Filename:   f.LogPath,    // 日志文件路径
		MaxSize:    f.MaxSize,    // megabytes
		MaxBackups: f.MaxBackups, // 最多保留300个备份
		Compress:   f.Compress,   // 是否压缩 disabled by default
	}
	if f.MaxAge > 0 {
		hook.MaxAge = f.MaxAge // days
	}
	return &hook
}
