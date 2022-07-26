package server

import (
	"cloudCli/cfg"
	"cloudCli/utils/log"
	"github.com/gin-gonic/gin"
)

/**
 *
 * @author jensen.chen
 * @date 2022/7/26
 */
/**
 * 文件检测的路由
 * @author jensen.chen
 * @date 2022/7/8
 */
type ServerAction struct {
	pingController   PingController
	ticketController TicketController
}

func (s *ServerAction) InitAction() {
	s.pingController = PingController{}
	s.pingController.Init()
	s.ticketController = TicketController{}
	s.ticketController.Init()
}

func (s *ServerAction) AddRouter(g *gin.RouterGroup) {
	leader, err := cfg.GetConfig("cli.server.leader")
	if err == nil && leader.(bool) {
		log.Debug("Enable Leader Server Action")
		serverGroup := g.Group("/node")
		{
			serverGroup.POST("/registe", s.pingController.RegisteNode) //注册节点
			serverGroup.POST("/ping", s.pingController.NodePing)       //节点Ping
		}
	}

}
