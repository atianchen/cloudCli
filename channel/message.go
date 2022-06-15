package channel

import "cloudCli/common"

/**
 *
 * @author jensen.chen
 * @date 2022/6/9
 */
const MESSAGE_CLOSE = "close"
const MESSAGE_ONTIME = "onTime"

type Message struct {
	common.ModalMap
	Payload interface{}
}

type CommandMessage struct {
	Message
	Name string //消息名
}

func (*Message) GetPayload() interface{} {
	return nil
}

func BuildCloseCommand() *CommandMessage {
	return &CommandMessage{Name: MESSAGE_CLOSE}
}
