package profile

import "github.com/gin-gonic/gin"

/**
 * 文件检测的路由
 * @author jensen.chen
 * @date 2022/7/8
 */
type ProfileAction struct {
	docAction DocController
}

func (s ProfileAction) InitAction() {
	s.docAction = DocController{}
	s.docAction.Init()
}
func (s ProfileAction) AddRouter(g *gin.RouterGroup) {

	profileGroup := g.Group("/profile")
	{
		profileGroup.POST("/doc/list", s.docAction.ListDoc)
		profileGroup.POST("/doc/detail", s.docAction.DocDetail)
	}

}
