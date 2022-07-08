package profile

import (
	"cloudCli/domain"
	"cloudCli/gin/dto"
	"cloudCli/repository"
	beanutils "github.com/atianchen/go-beanutils"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 *
 * @author jensen.chen
 * @date 2022/7/8
 */
type DocController struct {
	repository *repository.DocRepository
}

func (lc *DocController) Init() {
	lc.repository = &repository.DocRepository{}
}

/**
 * 列出所有文件
 */
func (lc *DocController) ListDoc(c *gin.Context) {

	var param dto.PageRequestDto
	c.BindJSON(&param)
	var docs []*domain.DocInfo
	err := lc.repository.PageQuery(&docs, param.Page*param.Limit, param.Limit, param.Keyword)
	if err != nil {
		c.JSON(http.StatusOK, dto.BuildErrorMsg(err.Error()))
	} else {
		var items []*DocDto
		for _, doc := range docs {
			docDto := DocDto{}
			beanutils.CopyProperties(&docDto, doc)
			items = append(items, &docDto)
		}
		c.JSON(http.StatusOK, dto.BuildSuccessMsg(dto.PageResponse{Page: param.Page, Limit: param.Limit, Items: &items}))
	}
}

/**
文档详情
*/
func (lc *DocController) DocDetail(c *gin.Context) {
	docId := c.Query("docId")
	if len(docId) > 0 {
		doc, err := lc.repository.GetByPrimary(docId)
		if err == nil {
			docDto := DocDto{}
			beanutils.CopyProperties(&docDto, doc)
			c.JSON(http.StatusOK, dto.BuildSuccessMsg(&docDto))
		} else {
			c.JSON(http.StatusOK, dto.BuildErrorMsg(err.Error()))
		}
	} else {
		c.JSON(http.StatusOK, dto.BuildErrorMsg("Miss Param"))
	}
}

/**
删除文档
*/
func (lc *DocController) DeleteDoc(c *gin.Context) {
	docId := c.Query("docId")
	if len(docId) > 0 {
		if err := lc.repository.RemoveByPrimary(docId); err != nil {
			c.JSON(http.StatusOK, dto.BuildEmptySuccessMsg())
		} else {
			c.JSON(http.StatusOK, dto.BuildErrorMsg(err.Error()))
		}
	} else {
		c.JSON(http.StatusOK, dto.BuildErrorMsg("Miss Param"))
	}
}
