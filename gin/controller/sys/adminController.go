package sys

import (
	"cloudCli/domain"
	"cloudCli/gin/dto"
	"cloudCli/gin/dto/sys"
	"cloudCli/gin/webConst"
	"cloudCli/repository"
	"cloudCli/utils/encrypt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * admin相关操作
 * @author jensen.chen
 * @date 2022/7/20
 */
type AdminController struct {
	repository *repository.SysUserRepository
}

func (lc *AdminController) Init() {
	lc.repository = &repository.SysUserRepository{}
}

func (lc *AdminController) UpdateAdminPwd(c *gin.Context) {
	var param sys.UpdatePwdDto
	c.BindJSON(&param)
	u, exists := c.Get(webConst.KEY_LOGINUSER)
	if exists {
		if err := lc.repository.UpdatePwd((u.(*domain.SysUser)).Id, encrypt.MD5(param.Pwd)); err == nil {
			c.JSON(http.StatusOK, dto.BuildEmptySuccessMsg())
		} else {
			c.JSON(http.StatusOK, dto.BuildErrorMsg(err.Error()))
		}
	} else {
		c.JSON(http.StatusOK, dto.BuildErrorMsg("Update Error,Not Login"))
	}
}
