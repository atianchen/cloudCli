package node

import (
	channel2 "cloudCli/channel"
	"reflect"
)

/**
 * 任务的相关定义
 * @author jensen.chen
 * @date 2022-05-20
 */
/**
消息处理
*/
type MsgHandler interface {
	HandleMessage(msg interface{})
}

/**
 * 任务
 */
type Node interface {
	Init()

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
	GetMsgHandler() MsgHandler

	MessageReceive(target Node, channel chan interface{})
}

type AbstractNode struct {
	Transpot chan interface{}
}

func (t *AbstractNode) Name() string {
	return reflect.TypeOf(t).Elem().Name()
}

/**
消息接收
*/
func (t *AbstractNode) MessageReceive(target Node, channel chan interface{}) {
	t.Transpot = channel
L:
	for {
		select {
		case msg := <-channel:
			switch msg.(type) {
			case *channel2.CommandMessage:
				{
					if msg.(*channel2.CommandMessage).Name == "close" {
						break L
					} else {
						handler := target.GetMsgHandler()
						if handler != nil {
							handler.HandleMessage(msg)
						}
					}
				}
			default:
				{
					handler := target.GetMsgHandler()
					if handler != nil {
						handler.HandleMessage(msg)
					}
				}
			}
		default:

		}
	}
}
