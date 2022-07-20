package sys

import (
	"cloudCli/domain"
	"cloudCli/gin/dto"
	"cloudCli/gin/dto/sys"
	"cloudCli/repository"
	go_beanutils "github.com/atianchen/go-beanutils"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * 配置参数的增删改查
 * @author jensen.chen
 * @date 2022/7/11
 */
type ParamController struct {
	repository *repository.ParamRepository
}

func (pc *ParamController) Init() {
	pc.repository = &repository.ParamRepository{}
}

/**
 * 列出所有文件
 */
func (pc *ParamController) ListParam(c *gin.Context) {

	var params []domain.Param
	err := pc.repository.GetAll(&params)
	if err != nil {
		c.JSON(http.StatusOK, dto.BuildErrorMsg(err.Error()))
	} else {
		var items []sys.ParamDto
		for _, param := range params {
			paramDto := sys.ParamDto{}
			go_beanutils.CopyProperties(&paramDto, param)
			items = append(items, paramDto)
		}
		c.JSON(http.StatusOK, dto.BuildSuccessMsg(params))
	}
}

/**
参数详情
*/
func (pc *ParamController) ParamInfo(c *gin.Context) {
	paramId := c.Query("paramId")
	if len(paramId) > 0 {
		param, err := pc.repository.GetByPrimary(paramId)
		if err == nil {
			paramDto := sys.ParamDto{}
			go_beanutils.CopyProperties(&paramDto, param)
			c.JSON(http.StatusOK, dto.BuildSuccessMsg(&paramDto))
		} else {
			c.JSON(http.StatusOK, dto.BuildErrorMsg(err.Error()))
		}
	} else {
		c.JSON(http.StatusOK, dto.BuildErrorMsg("Miss Param"))
	}
}

/**
更新参数的值
*/
func (pc *ParamController) UpdateParam(c *gin.Context) {
	var arg sys.ParamDto
	c.BindJSON(&arg)
	if len(arg.Id) > 0 {
		var err error
		param, err := pc.repository.GetByPrimary(arg.Id)
		param.Val = arg.Val
		if err == nil {

			err = pc.repository.Update(param)
		}
		if err == nil {
			c.JSON(http.StatusOK, dto.BuildEmptySuccessMsg())
		} else {
			c.JSON(http.StatusOK, dto.BuildErrorMsg(err.Error()))
		}
	} else {
		c.JSON(http.StatusOK, dto.BuildErrorMsg("Miss Param"))
	}

}
