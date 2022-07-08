package sys

import (
	"github.com/gin-gonic/gin"
)

/**
 * 用户登录
 * @author jensen.chen
 * @date 2022/7/7
 */
type SysAction struct {
	loginAction LoginController
}

func (s SysAction) InitAction() {
	s.loginAction = LoginController{}
	s.loginAction.Init()
}
func (s SysAction) AddRouter(g *gin.RouterGroup) {

	sysGroup := g.Group("/sys")
	{
		sysGroup.POST("/login", s.loginAction.Login)
	}

}
