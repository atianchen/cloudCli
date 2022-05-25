package log

import (
	"cloudCli/cfg"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"time"
)

type LogInit interface {
	LogWrite(writers ...io.Writer) zapcore.WriteSyncer
	LogFmt() zapcore.Encoder
	LogLevel(logLevel string) zapcore.Level
	Init()
}

type Log struct {
}

var logger *zap.SugaredLogger

/**
 * 日志输出
 */
func (l *Log) LogWrite(writers ...io.Writer) zapcore.WriteSyncer {
	var syncer []zapcore.WriteSyncer
	for _, writer := range writers {
		syncer = append(syncer, zapcore.AddSync(writer))
	}
	return zapcore.NewMultiWriteSyncer(syncer...)
}

/**
 * 日志格式
 */
func (l *Log) LogFmt() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "line",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    cEncodeLevel,
		EncodeTime:     cEncodeTime,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   cEncodeCaller,
		EncodeName:     zapcore.FullNameEncoder,
	}
	var encoder zapcore.Encoder
	encoder = zapcore.NewConsoleEncoder(encoderConfig)
	return encoder
}

/**
 * 日志级别
 */
func (l *Log) LogLevel(logLevel string) zapcore.Level {
	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}
	return level
}

/**
 * 从配置文件中获取信息，初始化日志配置
 */
func (l *Log) Init() {

	// 识别日志配置文件
	logConfig := cfg.GetConfig("cli.log")
	var keys []string
	if logConfig != nil {
		m := logConfig.(map[string]interface{})
		keys = cfg.GetKeys(m)
	}

	var writers []io.Writer
	// 初始化控制台
	consoleCfg := NewConsoleCfg()
	var logout LogOut = consoleCfg
	// 初始化扩展日志输出，可动态扩展
	writers = append(writers, logout.Instan())
	for _, key := range keys {
		var logout LogOut
		switch key {
		case "console":
			csWriter(consoleCfg)
			continue
		case "file":
			logout = fileWriter(logout)
		case "es":
			logout = NewEsConfig()
		default:
		}
		writers = append(writers, logout.Instan())
	}

	core := zapcore.NewCore(
		l.LogFmt(),
		l.LogWrite(writers...),
		l.LogLevel(consoleCfg.LogLevel),
	)
	ZapLogger := zap.New(core)
	ZapLogger = ZapLogger.WithOptions(zap.AddCaller())
	logger = ZapLogger.Sugar()
}

// 自定义格式
func cEncodeLevel(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + "level=" + level.CapitalString() + "]")
}

func cEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + "time=" + t.Format(logTmFmt) + "]")
}

func cEncodeCaller(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + caller.TrimmedPath() + "]")
}

func csWriter(consoleCfg *ConsoleConfig) {
	csConfig := cfg.GetConfig("cli.log.console")
	if csConfig == nil {
		return
	}
	m := csConfig.(map[string]interface{})
	err := mapstructure.Decode(m, &consoleCfg)
	if err != nil {
		fmt.Println("控制台配置解析失败", err)
	}
}

func fileWriter(logout LogOut) LogOut {
	fileConfig := cfg.GetConfig("cli.log.file")
	if fileConfig == nil {
		return nil
	}
	logout = NewFileCfg()
	m := fileConfig.(map[string]interface{})
	err := mapstructure.Decode(m, logout)
	if err != nil {
		fmt.Println("日志文件配置解析失败", err)
		return nil
	}
	return logout
}
