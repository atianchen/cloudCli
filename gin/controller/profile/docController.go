package profile

import (
	"cloudCli/channel"
	"cloudCli/ctx"
	"cloudCli/domain"
	"cloudCli/gin/dto"
	profile2 "cloudCli/gin/dto/profile"
	"cloudCli/node/profile"
	"cloudCli/repository"
	go_beanutils "github.com/atianchen/go-beanutils"
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
		var items []profile2.DocListDto
		for _, doc := range docs {
			docDto := profile2.DocListDto{}
			go_beanutils.CopyProperties(&docDto, doc)
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
			docDto := profile2.DocDto{}
			go_beanutils.CopyProperties(&docDto, doc)
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
			/**
			删除成功，则重置扫描
			*/
			lc.Reset(c)
		} else {
			c.JSON(http.StatusOK, dto.BuildErrorMsg(err.Error()))
		}
	} else {
		c.JSON(http.StatusOK, dto.BuildErrorMsg("Miss Param"))
	}
}

func (lc *DocController) Reset(c *gin.Context) {
	nodeChan, _ := channel.GetChan(ctx.PROFILE_NODE_NAME) //获取节点channel
	if nodeChan != nil {
		nodeChan <- profile.BuildRestCommand() //发送消息
		select {                               //等待节点回复
		case res := <-nodeChan:
			{
				response := res.(*channel.AsyncResponse)
				if response.Err == nil { //如果没有错误，则返回成功执行
					c.JSON(http.StatusOK, dto.BuildEmptySuccessMsg())
				} else { //返回错误信息
					c.JSON(http.StatusOK, dto.BuildErrorMsg(response.Err.Error()))
				}

			}
		case <-time.After(20 * time.Second): //20秒没收到回复，则返回超时错误
			c.JSON(http.StatusOK, dto.BuildErrorMsg("Execution timeout"))
		}
	} else {
		c.JSON(http.StatusOK, dto.BuildErrorMsg("Node communication failed"))
	}
}
