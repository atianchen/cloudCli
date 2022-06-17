package channel

/**
 *
 * @author jensen.chen
 * @date 2022/6/9
 */
const MESSAGE_CLOSE = "close"
const MESSAGE_ONTIME = "onTime"

type Message interface {
}

type CommandMessage struct {
	Payload interface{}
	Name    string //消息名
}

func (*CommandMessage) GetPayload() interface{} {
	return nil
}

/**
构建系统关闭Command
*/
func BuildCloseCommand() *CommandMessage {
	return &CommandMessage{Name: MESSAGE_CLOSE}
}

/**
构建定时任务
*/
func BulidOnTimeMessage(payload interface{}) *CommandMessage {
	cmd := &CommandMessage{Name: MESSAGE_CLOSE}
	cmd.Payload = payload
	return cmd
}
