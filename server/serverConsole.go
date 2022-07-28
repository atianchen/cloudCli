package server

import (
	channel2 "cloudCli/channel"
	"cloudCli/domain"
	"cloudCli/node"
	"cloudCli/node/extend"
	"cloudCli/utils/log"
	"time"
)

/**
 * 控制台
 * @author jensen.chen
 * @date 2022/7/25
 */
const EXPIRE_PERIOD = 10 * time.Minute

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
	switch msg.(type) {
	case *channel2.CommandMessage:
		{
			switch msg.(*channel2.CommandMessage).Name {
			case channel2.MESSAGE_ONTIME:
				{
					t.clearExpireNode()
				}
			}
		}
	}
}

/**
清除过期的注册信息
*/
func (t *ServerConsole) clearExpireNode() {
	if err := t.deployNodeService.RemoveExpireNode(int(time.Now().Add(-1 * EXPIRE_PERIOD).Unix())); err != nil {
		log.Error("Clear Expired Node Error ", err.Error())
	} else {
		log.Error("Clear Expired Node ")
	}

}

func (t *ServerConsole) GetMsgHandler() extend.MsgHandler {
	return t
}
