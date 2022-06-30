package log

import (
	"cloudCli/cfg"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
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
	logConfig, _ := cfg.GetConfig("cli.log")
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

// GinLogger 接收gin框架默认的日志，

func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		cost := time.Since(start)
		logger.Infow(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
		c.Next()
	}
}

// GinRecovery
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			var err interface{} = recover()
			if err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}

func csWriter(consoleCfg *ConsoleConfig) {
	csConfig, _ := cfg.GetConfig("cli.log.console")
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
	fileConfig, _ := cfg.GetConfig("cli.log.file")
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
