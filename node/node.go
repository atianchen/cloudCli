package node

import (
	"cloudCli/node/extend"
)

/**
 * 任务的相关定义
 * @author jensen.chen
 * @date 2022-05-20
 */

/**
 * 任务
 */
type Node interface {
	Init() error

	/**
	 * 开始
	 */
	Start(context *NodeContext)

	/**
	 * 停止
	 */
	Stop()

	/**
	 * 获取名称
	 */
	Name() string

	/**
	消息监听
	*/
	GetMsgHandler() extend.MsgHandler

	MessageReceive(target Node, channel chan interface{})
}
