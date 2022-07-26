package extend

/**
 *
 * @author jensen.chen
 * @date 2022/7/26
 */
/**
消息处理
*/
type MsgHandler interface {
	HandleMessage(msg interface{}, channel chan interface{})
}
