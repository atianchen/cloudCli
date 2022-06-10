package channel

import "cloudCli/common"

/**
 *
 * @author jensen.chen
 * @date 2022/6/9
 */
const MESSAGE_CLOSE = "close"

type PluginMessage interface {
	GetPayload() interface{}
}

type CommandMessage struct {
	common.ModalMap
	Name string //消息名
}

func (*CommandMessage) GetPayload() interface{} {
	return nil
}

func BuildCloseCommand() *CommandMessage {
	return &CommandMessage{Name: MESSAGE_CLOSE}
}
