package server

import (
	"cloudCli/domain"
	"cloudCli/node"
	"cloudCli/node/extend"
)

/**
 * 控制台
 * @author jensen.chen
 * @date 2022/7/25
 */
type ServerConsole struct {
	node.AbstractNode
	deployNodeService *DeployNodeService
	nodes             []domain.DeployNode
}

func (t *ServerConsole) Init() error {
	t.deployNodeService = &DeployNodeService{}
	t.nodes = []domain.DeployNode{}
	return t.deployNodeService.LoadAll(&t.nodes)
}

func (t *ServerConsole) Start(context *node.NodeContext) {
	/**
	如果配置了Leader属性，则启用Server端功能
	*/
	//leader, err := cfg.GetConfig("cli.server.leader")

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

func (t *ServerConsole) HandleMessage(msg interface{}, channel chan interface{}) {

}

func (t *ServerConsole) GetMsgHandler() extend.MsgHandler {
	return t
}
