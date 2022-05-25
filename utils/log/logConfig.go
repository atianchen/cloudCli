package log

const (
	logTmFmt = "2006-01-02 15:04:05.000Z0700"
)

type logConfig interface {
}

// 基础日志输出
type BaseLogOut struct {
	LogToConsole string
	LogToFile    string
}

// 扩展日志输出
type ExtendLogOut struct {
	LogToEs string
}
