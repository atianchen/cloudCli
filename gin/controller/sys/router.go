package sys

import (
	"cloudCli/gin/security"
	"github.com/gin-gonic/gin"
)

/**
 * 用户登录
 * @author jensen.chen
 * @date 2022/7/7
 */
type SysAction struct {
	loginAction  LoginController
	logoutAction LogoutController
	paramAction  ParamController
	configAction ConfigController
}

func (s SysAction) InitAction() {
	s.loginAction = LoginController{}
	s.loginAction.Init()
	s.paramAction = ParamController{}
	s.paramAction.Init()
	s.logoutAction = LogoutController{}
	s.logoutAction.Init()
}
func (s SysAction) AddRouter(g *gin.RouterGroup) {

	sysGroup := g.Group("/sys")
	{
		sysGroup.POST("/login", s.loginAction.Login)
		sysGroup.GET("/logout", s.logoutAction.Logout)
		sysGroup.POST("/currentUser", s.loginAction.CurrentUser)
	}
	paramGroup := g.Group("/sys/param")
	{
		paramGroup.POST("/list", security.JwtAuthInterceptor(), s.paramAction.ListParam)
		paramGroup.POST("/detail", security.JwtAuthInterceptor(), s.paramAction.ParamInfo)
		paramGroup.POST("/update", security.JwtAuthInterceptor(), s.paramAction.UpdateParam)
	}
	configGroup := g.Group("/sys/config")
	{
		configGroup.POST("/list", security.JwtAuthInterceptor(), s.configAction.GetConfig)
	}

}
