package node

import (
	"cloudCli/channel"
	"cloudCli/common"
)

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
 *
 * @author jensen.chen
 * @date 2022/7/4
 */

/**
初始化环境变量
*/
func (ctx *NodeContext) Init(target interface{}) {
	/**
	通道
	*/
	if target != nil {
		ctx.Channel = channel.CreateChan(target.(Node).Name())
	}
}

/**
 * 创建默认Context
 */
func CreateNodeContext(target Node) *NodeContext {
	ctx := &NodeContext{}
	ctx.Init(target)
	ctx.AttrMap = make(map[string]interface{})
	return ctx
}
