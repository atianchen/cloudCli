package plugin

import (
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
			case CommandMessage:
				{
					if msg.(*CommandMessage).Name == "close" {
						break L
					} else {
						b.HandleMessage(msg.(*PluginMessage))
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
func (b *BasePlugin) HandleMessage(msg *PluginMessage) *PluginResponse {
	return nil
}

func (b *BasePlugin) Execute(context ctx.Context, params ExecuteParams) {
	var dt = context.(*ctx.DefaultContext)
	if dt.Channel != nil {
		go b.dispatch(dt.Channel)
	}
}
