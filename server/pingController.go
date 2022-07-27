package server

import (
	"cloudCli/domain"
	"cloudCli/gin/dto"
	"cloudCli/utils/timeUtils"
	go_beanutils "github.com/atianchen/go-beanutils"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 *
 * @author jensen.chen
 * @date 2022/7/26
 */
type PingController struct {
	deployNodeService *DeployNodeService
	BaseController
}

func (p *PingController) Init() {
	p.deployNodeService = &DeployNodeService{}
}

/**
注册节点
*/
func (p *PingController) RegisteNode(c *gin.Context) {
	var param DeployNodeDto
	if err := p.GetPayload(c, &param); err != nil {
		c.JSON(http.StatusBadRequest, dto.BuildErrorMsg(err.Error()))
		return
	}
	var node domain.DeployNode
	if err := go_beanutils.CopyProperties(&node, &param); err != nil {
		c.JSON(http.StatusBadRequest, dto.BuildErrorMsg(err.Error()))
		return
	}
	node.Status = NODE_STATUS_VALID
	node.Ts = timeUtils.NowUnixTime()
	if err := p.deployNodeService.Save(&node); err == nil {
		c.JSON(http.StatusOK, dto.BuildEmptySuccessMsg())
	} else {
		c.JSON(http.StatusBadRequest, dto.BuildErrorMsg(err.Error()))
	}
}

/**
注册节点
*/
func (p *PingController) NodePing(c *gin.Context) {
	var param DeployNodeDto
	if err := p.GetPayload(c, &param); err != nil {
		c.JSON(http.StatusBadRequest, dto.BuildErrorMsg(err.Error()))
		return
	}
	node, _ := p.deployNodeService.GetByIpAndPort(param.Ip, param.Port)
	if node != nil {
		node.Status = NODE_STATUS_VALID
		node.Ts = timeUtils.NowUnixTime()
		if err := p.deployNodeService.UpdateTs(node); err == nil {
			c.JSON(http.StatusOK, dto.BuildEmptySuccessMsg())
		} else {
			c.JSON(http.StatusOK, dto.BuildErrorMsg(err.Error()))
		}
	} else {
		c.JSON(http.StatusOK, dto.BuildErrorMsg("No Matching Node"))
	}
}
