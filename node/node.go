package node

import (
	channel2 "cloudCli/channel"
	"cloudCli/ctx"
	"reflect"
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
	Init()

	/**
	 * 开始
	 */
	Start(context ctx.Context)

	/**
	 * 停止任
	 */
	Stop()

	/**
	 * 获取名称
	 */
	Name() string
}

type AbstractNode struct {
}

func (t *AbstractNode) Name() string {
	return reflect.TypeOf(t).Elem().Name()
}

/**
消息分发
*/
func (b *AbstractNode) dispatch(channel chan interface{}) {
L:
	for {
		select {
		case msg := <-channel:
			switch msg.(type) {
			case channel2.CommandMessage:
				{
					if msg.(*channel2.CommandMessage).Name == "close" {
						break L
					} else {
						b.HandleMessage(msg.(*channel2.Message))
					}
				}
			}
		default:
		}
	}
}

/**
处理消息
*/
func (b *AbstractNode) HandleMessage(msg *channel2.Message) *AsyncResponse {
	return nil
}
