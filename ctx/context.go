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
	Init()
}

/**
 * 默认环境变量
 */
type DefaultContext struct {
	common.BaseObj
	common.ModalMap
	/*
		消息通道
	*/
	Channel chan interface{}
}

/**
初始化环境变量
*/
func (ctx *DefaultContext) Init() {
	/**
	通道
	*/
	ctx.Channel = make(chan interface{})
}

/**
 * 创建默认Context
 */
func CreateContext() Context {
	ctx := &DefaultContext{}
	ctx.Init()
	ctx.AttrMap = make(map[string]interface{})
	return ctx //(Context)(unsafe.Pointer(ctx))
}
