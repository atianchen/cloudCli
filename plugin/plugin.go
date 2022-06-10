package plugin

import (
	channel2 "cloudCli/channel"
	"cloudCli/ctx"
)

/**
 * 插件
 */
type Plugin interface {
	Execute(context ctx.Context, params ExecuteParams)
}
type BasePlugin struct {
	Plugin
}

/**
消息分发
*/
func (b *BasePlugin) dispatch(channel chan interface{}) {
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
func (b *BasePlugin) HandleMessage(msg *channel2.Message) *PluginResponse {
	return nil
}

func (b *BasePlugin) Execute(context ctx.Context, params ExecuteParams) {
	var dt = context.(*ctx.DefaultContext)
	if dt.Channel != nil {
		go b.dispatch(dt.Channel)
	}
}
