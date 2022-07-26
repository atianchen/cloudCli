package node

import (
	channel2 "cloudCli/channel"
	"reflect"
	"time"
)

/**
 *
 * @author jensen.chen
 * @date 2022/7/26
 */
type AbstractNode struct {
}

func (t *AbstractNode) Name() string {
	return reflect.TypeOf(t).Elem().Name()
}

/**
消息接收
*/
func (t *AbstractNode) MessageReceive(target Node, channel chan interface{}) {
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
							handler.HandleMessage(msg, channel)
						}
					}
				}
			default:
				{
					handler := target.GetMsgHandler()
					if handler != nil {
						handler.HandleMessage(msg, channel)
					}
				}
			}
		default:

		}
		time.Sleep(5 * time.Second)
	}
}
