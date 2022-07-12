package channel

/**
 *
 * @author jensen.chen
 * @date 2022/6/9
 */
const MESSAGE_CLOSE = "close"
const MESSAGE_ONTIME = "onTime"
const MESSGE_RESPONSE = "response"

type Message interface {
}

/**
Web请求的Message
*/
type RequestMessage struct {
	Payload interface{}
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
	cmd := &CommandMessage{Name: MESSAGE_ONTIME}
	cmd.Payload = payload
	return cmd
}
