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
	paramAction ParamController
}

func (s SysAction) InitAction() {
	s.loginAction = LoginController{}
	s.loginAction.Init()
	s.paramAction = ParamController{}
	s.paramAction.Init()
}
func (s SysAction) AddRouter(g *gin.RouterGroup) {

	sysGroup := g.Group("/sys")
	{
		sysGroup.POST("/login", s.loginAction.Login)
		sysGroup.POST("/currentUser", s.loginAction.CurrentUser)
	}
	paramGroup := g.Group("/sys/param")
	{
		paramGroup.POST("/list", s.paramAction.ListParam)
		paramGroup.POST("/detail", s.paramAction.ParamInfo)
		paramGroup.POST("/update", s.paramAction.UpdateParam)
	}

}
