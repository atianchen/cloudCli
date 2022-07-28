package server

import (
	"cloudCli/domain"
	"cloudCli/gin/dto"
	go_beanutils "github.com/atianchen/go-beanutils"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * 节点管理
 * @author jensen.chen
 * @date 2022/7/28
 */
type NodeManageController struct {
	deployNodeService *DeployNodeService
}

func (lc *NodeManageController) Init() {
	lc.deployNodeService = &DeployNodeService{}
}

/**
 *
 * @author jensen.chen
 * @date 2022/7/15
 */
func (lc *NodeManageController) ListNode(c *gin.Context) {
	var param dto.PageRequestDto
	c.BindJSON(&param)
	var nodes []domain.DeployNode
	err := lc.deployNodeService.PageQuery(&nodes, param.Page*param.Limit, param.Limit, param.Keyword)
	if err != nil {
		c.JSON(http.StatusOK, dto.BuildErrorMsg(err.Error()))
	} else {
		var items []DeployNodeDto
		for _, node := range nodes {
			nodeDto := DeployNodeDto{}
			go_beanutils.CopyProperties(&nodeDto, node)
			items = append(items, nodeDto)
		}
		c.JSON(http.StatusOK, dto.BuildSuccessMsg(dto.PageResponse{Page: param.Page, Limit: param.Limit, Items: &items}))
	}
}
