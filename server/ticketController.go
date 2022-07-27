package server

import (
	"cloudCli/db"
	"cloudCli/gin/dto"
	"cloudCli/utils/encrypt"
	"cloudCli/utils/timeUtils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"net/http"
	"strconv"
	"time"
)

const EXPIRES_TIME = 10 //TOKEN过期时间 单位:秒

/**
 *
 * @author jensen.chen
 * @date 2022/7/26
 */
type TicketController struct {
	deployNodeService *DeployNodeService
}

func (tc *TicketController) Init() {
	tc.deployNodeService = &DeployNodeService{}
}

/**
验证票据
*/
func (tc *TicketController) VerifyTicket(c *gin.Context) {
	token := c.Query("token")
	var ncTicket NcTicket
	if err := encrypt.ParseToken(token, &ncTicket); err == nil {
		cacheKey := ncTicket.Ip + "_" + strconv.Itoa(int(ncTicket.Port)) + "_ticket"
		_, err = db.MapDbInst.Get(cacheKey)
		if err != nil {
			c.Status(http.StatusBadRequest)
		} else {
			db.MapDbInst.Remove(cacheKey)
			c.JSON(http.StatusOK, dto.BuildEmptySuccessMsg())
		}
	} else {
		c.Status(http.StatusBadRequest)
	}

}

/**
重定向到远程登录节点
*/
func (tc *TicketController) RedirectNode(c *gin.Context) {
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Header("Pragma", "no-cache")
	c.Header("Expires", "0")
	nodeId := c.Query("nodeId")
	nodeImpl, err := tc.deployNodeService.GetByPrimary(nodeId)
	if err != nil {
		c.Status(http.StatusForbidden)
	} else {
		ticket := uuid.New().String()
		cacheKey := nodeImpl.Ip + "_" + strconv.Itoa(int(nodeImpl.Port)) + "_ticket"
		cacheVal, err := json.Marshal(ticket)
		if err == nil {
			db.MapDbInst.Set(cacheKey, string(cacheVal), EXPIRES_TIME*time.Second)
			claims := NcTicket{
				nodeImpl.Ip,
				nodeImpl.Port,
				ticket,
				jwt.StandardClaims{
					ExpiresAt: time.Now().Add(EXPIRES_TIME * time.Second).Unix(),
				},
			}
			token, _ := encrypt.GenerateToken(claims)
			var ncTicket2 NcTicket
			encrypt.ParseToken(token, &ncTicket2)
			c.Redirect(http.StatusMovedPermanently, "http://"+nodeImpl.Ip+":"+strconv.Itoa(int(nodeImpl.Port))+"/cloud/remote/accept?token="+token+"&ts="+strconv.Itoa(int(timeUtils.NowUnixTime())))
		} else {
			c.Status(http.StatusForbidden)
		}
	}
}
