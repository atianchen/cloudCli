package log

type ConsoleConfig struct {
	/*
	  日志级别，默认为INFO
	*/
	LogLevel string `json:"logLevel"`
}

type FileConfig struct {
	/**
	  日志文件路径
	*/
	LogPath string `json:"logPath"`
	/**
	  最大的日志文件大小，MB
	*/
	MaxSize int `json:"maxSize"`
	/**
	  日志备份个数
	*/
	MaxBackups int `json:"maxBackups"`
	/**
	  历史日志的保留天数，默认不设置
	*/
	MaxAge int `json:"maxAge"`
	/**
	  是否压缩
	*/
	Compress bool `json:"compress"`
}

func NewConsoleCfg() *ConsoleConfig {
	entry := &ConsoleConfig{
		LogLevel: "info",
	}
	return entry
}

func NewFileCfg() *FileConfig {

	entry := &FileConfig{
		LogPath:    "./logs/cloudCli.log",
		MaxSize:    20,
		MaxBackups: 10,
		Compress:   false,
	}
	return entry
}
