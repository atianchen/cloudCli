package server

import (
	"cloudCli/node"
)

/**
 * 控制台
 * @author jensen.chen
 * @date 2022/7/25
 */
type ServerConsole struct {
	node.AbstractNode
	cliClientService *CliClientService
}

func (t *ServerConsole) Init() error {
	return t.cliClientService.load()
}

func (t *ServerConsole) Start(context *node.NodeContext) {

}

/**
 * 停止
 */
func (t *ServerConsole) Stop() {

}

/**
 * 获取名称
 */
func (t *ServerConsole) Name() string {
	return "serverConsole"
}

func (t *ServerConsole) HandleMessage(msg interface{}) {

}

func (t *ServerConsole) GetMsgHandler() node.MsgHandler {
	return t
}
