package ctx

//import "unsafe"
import (
	"cloudCli/channel"
	"cloudCli/common"
	"reflect"
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

/**
 * 默认环境变量
 */
type NodeContext struct {
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
func (ctx *NodeContext) Init(target interface{}) {
	/**
	通道
	*/
	if target != nil {
		ctx.Channel = channel.CreateChan(reflect.TypeOf(target).Elem().Name())
	}
}

/**
 * 创建默认Context
 */
func CreateNodeContext(target interface{}) *NodeContext {
	ctx := &NodeContext{}
	ctx.Init(target)
	ctx.AttrMap = make(map[string]interface{})
	return ctx
}
