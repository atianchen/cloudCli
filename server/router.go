package server

import (
	"cloudCli/cfg"
	"cloudCli/gin/security"
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
	pingController       PingController
	ticketController     TicketController
	nodeManageController NodeManageController
}

func (s *ServerAction) InitAction() {
	s.pingController = PingController{}
	s.pingController.Init()
	s.ticketController = TicketController{}
	s.ticketController.Init()
	s.nodeManageController = NodeManageController{}
	s.nodeManageController.Init()
}

func (s *ServerAction) AddRouter(g *gin.RouterGroup) {
	leader, err := cfg.GetConfig("cli.server.leader")
	if err == nil && leader.(bool) {
		log.Debug("Enable Leader Server Action")
		serverGroup := g.Group("/node")
		{
			serverGroup.POST("/registe", ServerAuthInterceptor(), s.pingController.RegisteNode) //注册节点
			serverGroup.POST("/ping", ServerAuthInterceptor(), s.pingController.NodePing)       //节点Ping
			serverGroup.POST("/verify", s.ticketController.VerifyTicket)                        //验证票据
			serverGroup.GET("/redirect", s.ticketController.RedirectNode)                       //重定向
			serverGroup.POST("/list", security.JwtAuthInterceptor(), s.nodeManageController.ListNode)
		}
	}

}
