package sysAction

import (
	"github.com/gin-gonic/gin"
)

/**
 *
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

	upload := g.Group("/sys")
	{
		upload.POST("/login", s.loginAction.Login)
	}

}
