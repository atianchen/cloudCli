package profile

import (
	"cloudCli/gin/security"
	"github.com/gin-gonic/gin"
)

/**
 * 文件检测的路由
 * @author jensen.chen
 * @date 2022/7/8
 */
type ProfileAction struct {
	docAction    DocController
	docHisAction DocHisController
}

func (s ProfileAction) InitAction() {
	s.docAction = DocController{}
	s.docAction.Init()
	s.docHisAction = DocHisController{}
	s.docHisAction.Init()
}

func (s ProfileAction) AddRouter(g *gin.RouterGroup) {

	profileGroup := g.Group("/profile")
	{
		profileGroup.POST("/doc/list", security.JwtAuthInterceptor(), s.docAction.ListDoc)
		profileGroup.POST("/doc/detail", security.JwtAuthInterceptor(), s.docAction.DocDetail)
		profileGroup.POST("/doc/delete", security.JwtAuthInterceptor(), s.docAction.DeleteDoc)
		profileGroup.POST("/doc/reset", security.JwtAuthInterceptor(), s.docAction.Reset)
	}
	inspectGroup := g.Group("/inspect")
	{
		inspectGroup.POST("/his/list", security.JwtAuthInterceptor(), s.docHisAction.ListDocHis)
		inspectGroup.POST("/his/detail", security.JwtAuthInterceptor(), s.docHisAction.DocHisDetail)
		inspectGroup.POST("/his/handle", security.JwtAuthInterceptor(), s.docHisAction.HandleDocHis)
	}

}
