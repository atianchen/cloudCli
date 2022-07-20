package sys

import (
	"cloudCli/cfg"
	"cloudCli/gin/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 *
 * @author jensen.chen
 * @date 2022/7/18
 */

type ConfigController struct {
}

func (lc *ConfigController) Init() {

}

func (lc *ConfigController) GetConfig(c *gin.Context) {
	val, err := cfg.GetConfig("cli")
	if err == nil {
		c.JSON(http.StatusOK, dto.BuildSuccessMsg(&val))
	} else {
		c.JSON(http.StatusOK, dto.BuildErrorMsg(err.Error()))
	}

}
