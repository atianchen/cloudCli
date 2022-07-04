package ctx

//import "unsafe"
import (
	"cloudCli/common"
)

/**
 * @author jensen.chen
 * @date 2022-05-20
 * 上下文环境变量
 */
type Context interface {
	common.Extends
	Init(target interface{})
}
