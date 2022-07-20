package profile

import (
	"cloudCli/domain"
	"cloudCli/gin/dto"
	"cloudCli/repository"
	"cloudCli/utils/timeUtils"
	go_beanutils "github.com/atianchen/go-beanutils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

/**
 * 查看变更历史
 * @author jensen.chen
 * @date 2022/7/19
 */
type DocHisController struct {
	repository    *repository.DocHisRepository
	docRepository *repository.DocRepository
}

func (lc *DocHisController) Init() {
	lc.repository = &repository.DocHisRepository{}
	lc.docRepository = &repository.DocRepository{}
}

/**
 * 列出所有文件
 */
func (lc *DocHisController) ListDocHis(c *gin.Context) {
	var param dto.DocHisPageRequestDto
	c.BindJSON(&param)
	var docs []domain.DocHistory
	err := lc.repository.PageQuery(&docs, param.Page*param.Limit, param.Limit, param.Status, param.Keyword)
	if err != nil {
		c.JSON(http.StatusOK, dto.BuildErrorMsg(err.Error()))
	} else {
		var items []DocHisDto
		for _, doc := range docs {
			docDto := DocHisDto{}
			go_beanutils.CopyProperties(&docDto, doc)
			items = append(items, docDto)
		}
		c.JSON(http.StatusOK, dto.BuildSuccessMsg(dto.PageResponse{Page: param.Page, Limit: param.Limit, Items: &items}))
	}
}

/**
文档详情
*/
func (lc *DocHisController) DocHisDetail(c *gin.Context) {
	hisId := c.Query("hisId")
	if len(hisId) > 0 {
		his, err := lc.repository.GetByPrimary(hisId)
		if err == nil {
			docHisDto := DocHisDetailDto{}
			go_beanutils.CopyProperties(&docHisDto, his)
			c.JSON(http.StatusOK, dto.BuildSuccessMsg(&docHisDto))
		} else {
			c.JSON(http.StatusOK, dto.BuildErrorMsg(err.Error()))
		}
	} else {
		c.JSON(http.StatusOK, dto.BuildErrorMsg("Miss Param"))
	}
}

/**
处理结果
*/
func (lc *DocHisController) HandleDocHis(c *gin.Context) {
	var param dto.DocHandleDto
	c.BindJSON(&param)
	time := timeUtils.TimeConfig{time.Now()}
	his, _ := lc.repository.GetByPrimary(param.Id)
	param.HandleTime = time.Unix()
	err := lc.repository.UpdateHandleResult(&param)
	switch param.HandleResult {
	case domain.DOCHIS_RESULT_RESERVE:
		{
			//接受现有的变更,修改DOC的数据
			err = lc.docRepository.UpdateContent(his.DocId, his.Hash, his.Content)
		}
	default:
		{

		}
	}
	if err == nil {
		c.JSON(http.StatusOK, dto.BuildEmptySuccessMsg())
	} else {
		c.JSON(http.StatusOK, dto.BuildErrorMsg(err.Error()))
	}
}
