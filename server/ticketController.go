package server

import (
	"cloudCli/db"
	"cloudCli/gin/dto"
	"cloudCli/utils/encrypt"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"net/http"
	"strconv"
	"time"
)

const EXPIRES_TIME = 6 //TOKEN过期时间 单位:秒

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
	var param TicketVerifyDto
	if err := c.BindJSON(&param); err != nil {
		c.Status(http.StatusBadRequest)
	} else {
		var ncTicket NcTicketToken
		if err := encrypt.ParseToken(param.Ticket, &ncTicket); err == nil {
			cacheKey := param.Ip + "_" + strconv.Itoa(param.Port) + "_ticket"
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

}

/**
重定向到远程登录节点
*/
func (tc *TicketController) RedirectNode(c *gin.Context) {
	nodeId := c.Query("nodeId")
	nodeImpl, err := tc.deployNodeService.GetByPrimary(nodeId)
	if err != nil {
		c.Status(http.StatusForbidden)
	} else {
		ticket := NcTicket{
			nodeImpl.Ip,
			nodeImpl.Port,
			uuid.New().String(),
		}
		cacheKey := nodeImpl.Ip + "_" + strconv.Itoa(nodeImpl.Port) + "_ticket"
		cacheVal, err := json.Marshal(ticket)
		if err == nil {
			db.MapDbInst.Set(cacheKey, string(cacheVal), EXPIRES_TIME*time.Second)
			claims := NcTicketToken{
				ticket,
				jwt.StandardClaims{
					ExpiresAt: time.Now().Add(EXPIRES_TIME * time.Second).Unix(),
					Issuer:    "admin",
				},
			}
			token, _ := encrypt.GenerateToken(claims)
			c.Redirect(http.StatusMovedPermanently, "http://"+nodeImpl.Ip+":"+strconv.Itoa(nodeImpl.Port)+"/cloud/remote/accept?token="+token)
		} else {
			c.Status(http.StatusForbidden)
		}
	}
}
