package notify

import (
	"cloudCli/gin/security"
	"github.com/gin-gonic/gin"
)

/**
 *
 * @author jensen.chen
 * @date 2022/7/21
 */
type NofityAction struct {
	historyController NotifyHistoryController
}

func (s *NofityAction) InitAction() {
	s.historyController = NotifyHistoryController{}
	s.historyController.Init()
}

func (s *NofityAction) AddRouter(g *gin.RouterGroup) {

	notifyGroup := g.Group("/notify")
	{
		notifyGroup.POST("/history/list", security.JwtAuthInterceptor(), s.historyController.ListNotifyHistory)
		notifyGroup.POST("/history/detail", security.JwtAuthInterceptor(), s.historyController.NotifyHistoryDetail)
	}

}
