package sys

import (
	"cloudCli/gin/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LogoutController struct {
}

func (lc *LogoutController) Init() {
}

/**
 *
 * @author jensen.chen
 * @date 2022/7/15
 */
func (lc *LogoutController) Logout(c *gin.Context) {
	c.SetCookie("cloudst", "", -1, "/", "/cloudCli", true, true)
	c.JSON(http.StatusOK, dto.BuildEmptySuccessMsg())
}
