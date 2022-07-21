package sys

import (
	"cloudCli/cfg"
	"cloudCli/channel"
	"cloudCli/gin/dto"
	"cloudCli/gin/dto/profile"
	_profile2 "cloudCli/node/profile"
	"cloudCli/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
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

/**
获取文件巡检的配置
*/
func (lc *ConfigController) GetProfileConfig(c *gin.Context) {
	configFile, err := cfg.GetConfig("cli.profile-inspect.config")
	if err == nil {
		filePath, err := utils.GetFilePath(configFile.(string))
		if err == nil {
			content, err := utils.GetFileStringContent(filePath)
			if err == nil {
				c.JSON(http.StatusOK, dto.BuildSuccessMsg(&profile.InspectConfigDto{Content: content}))
			} else {
				c.JSON(http.StatusOK, dto.BuildErrorMsg(err.Error()))
			}
		} else {
			c.JSON(http.StatusOK, dto.BuildErrorMsg(err.Error()))
		}
	} else {
		c.JSON(http.StatusOK, dto.BuildErrorMsg(err.Error()))
	}
}

/**
保存配置文件
*/
func (lc *ConfigController) SaveProfileConfig(c *gin.Context) {
	var param profile.InspectConfigDto
	c.BindJSON(&param)
	var configs []_profile2.ProfileConfig
	if jsonError := json.Unmarshal([]byte(param.Content), &configs); jsonError != nil {
		c.JSON(http.StatusOK, dto.BuildErrorMsg(jsonError.Error()))
		return
	}
	configFile, err := cfg.GetConfig("cli.profile-inspect.config")
	if err == nil {
		filePath, err := utils.GetFilePath(configFile.(string))
		if err == nil {
			if err := utils.WriteStringToFile(filePath, strings.Replace(param.Content, " ", "", -1)); err == nil {
				nodeChan, _ := channel.GetChan(_profile2.PROFILE_NODE_NAME) //获取节点channel
				if nodeChan != nil {
					nodeChan <- _profile2.BuildRestCommand() //发送消息
					select {                                 //等待节点回复
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
			} else {
				c.JSON(http.StatusOK, dto.BuildErrorMsg(err.Error()))
			}
		} else {
			c.JSON(http.StatusOK, dto.BuildErrorMsg(err.Error()))
		}
	}

}
