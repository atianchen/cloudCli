package notify

import (
	"cloudCli/domain"
	"cloudCli/gin/dto"
	"cloudCli/gin/dto/notify"
	"cloudCli/repository"
	go_beanutils "github.com/atianchen/go-beanutils"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 *
 * @author jensen.chen
 * @date 2022/7/21
 */
type NotifyHistoryController struct {
	repository *repository.NotifyHistoryRepository
}

func (lc *NotifyHistoryController) Init() {
	lc.repository = &repository.NotifyHistoryRepository{}
}

func (lc *NotifyHistoryController) ListNotifyHistory(c *gin.Context) {

	var param dto.PageRequestDto
	c.BindJSON(&param)
	var items []domain.NotifyHistory
	err := lc.repository.PageQuery(&items, param.Page*param.Limit, param.Limit)
	if err != nil {
		c.JSON(http.StatusOK, dto.BuildErrorMsg(err.Error()))
	} else {
		var dtos []notify.NotifyHistoryDto
		for _, item := range items {
			hisDto := notify.NotifyHistoryDto{}
			go_beanutils.CopyProperties(&hisDto, item)
			dtos = append(dtos, hisDto)
		}
		c.JSON(http.StatusOK, dto.BuildSuccessMsg(dto.PageResponse{Page: param.Page, Limit: param.Limit, Items: &dtos}))
	}
}

func (lc *NotifyHistoryController) NotifyHistoryDetail(c *gin.Context) {
	hisId := c.Query("hisId")
	if len(hisId) > 0 {
		his, err := lc.repository.GetByPrimary(hisId)
		if err == nil {
			notifyDto := notify.NotifyHistoryDto{}
			go_beanutils.CopyProperties(&notifyDto, his)
			c.JSON(http.StatusOK, dto.BuildSuccessMsg(&notifyDto))
		} else {
			c.JSON(http.StatusOK, dto.BuildErrorMsg(err.Error()))
		}
	} else {
		c.JSON(http.StatusOK, dto.BuildErrorMsg("Miss Param"))
	}
}
