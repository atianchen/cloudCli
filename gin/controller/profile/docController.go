package profile

import (
	"cloudCli/channel"
	"cloudCli/domain"
	"cloudCli/gin/dto"
	"cloudCli/node/profile"
	"cloudCli/repository"
	"cloudCli/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

/**
 * 文档增删改查
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
	var docs []domain.DocInfo
	err := lc.repository.PageQuery(&docs, param.Page*param.Limit, param.Limit, param.Keyword)
	if err != nil {
		c.JSON(http.StatusOK, dto.BuildErrorMsg(err.Error()))
	} else {
		var items []DocListDto
		for _, doc := range docs {
			docDto := DocListDto{}
			utils.CopyProperties(&docDto, doc)
			items = append(items, docDto)
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
			utils.CopyProperties(&docDto, doc)
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

func (lc *DocController) Reset(c *gin.Context) {
	nodeChan, _ := channel.GetChan(profile.PROFILE_NODE_NAME)
	if nodeChan != nil {
		nodeChan <- profile.BuildRestCommand()
		select {
		case res := <-nodeChan:
			{
				response := res.(*channel.AsyncResponse)
				if response.Err == nil {
					c.JSON(http.StatusOK, dto.BuildEmptySuccessMsg())
				} else {
					c.JSON(http.StatusOK, dto.BuildErrorMsg(response.Err.Error()))
				}

			}
		case <-time.After(20 * time.Second): //10秒没执行，则直接报超时
			c.JSON(http.StatusOK, dto.BuildErrorMsg("Execution timeout"))
		}
	} else {
		c.JSON(http.StatusOK, dto.BuildErrorMsg("Node communication failed"))
	}
}
