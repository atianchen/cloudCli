package log

import (
	"io"
)

/**
 * 扩展日志输出，这里以es为例
 */
func (e *EsConfig) Instan() io.Writer {
	return nil
}
