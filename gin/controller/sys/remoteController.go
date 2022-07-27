package sys

import (
	"bytes"
	"cloudCli/cfg"
	"cloudCli/ctx"
	"cloudCli/gin/security"
	"cloudCli/repository"
	"cloudCli/server"
	"cloudCli/utils/encrypt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strconv"
	"time"
)

/**
 *
 * @author jensen.chen
 * @date 2022/7/26
 */
type RemoteController struct {
	repository *repository.SysUserRepository
}

func (lc *RemoteController) Init() {
	lc.repository = &repository.SysUserRepository{}
}

func (lc *RemoteController) Accpet(c *gin.Context) {
	leaderAddr, err := cfg.GetConfig("cli.cloud.addr")
	if err != nil {
		c.Writer.WriteString("InValid Request")
		return
	}
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Header("Pragma", "no-cache")
	c.Header("Expires", "0")
	token := c.Query("token")
	var ncTicket server.NcTicket
	if err := encrypt.ParseToken(token, &ncTicket); err == nil {
		resp, err := http.Post(leaderAddr.(string)+"/node/verify?token="+token, "application/json", bytes.NewReader([]byte{}))
		if err != nil || resp.StatusCode != http.StatusOK {
			c.Writer.WriteString("InValid Request")
		} else {
			url := "http://" + ctx.APPINFO.SERVER_BIND + ":" + strconv.Itoa(int(ctx.APPINFO.SERVER_PORT)) + "/cloud/ui/index.html"

			user, err := lc.repository.FindByCode("admin")
			if err != nil {
				c.Writer.WriteString("InValid Request")
			} else {
				claims := security.Realm{
					user.Id,
					user.Code,
					jwt.StandardClaims{
						ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
						Issuer:    "admin",
					},
				}
				token, _ := encrypt.GenerateToken(claims)
				http.SetCookie(c.Writer, &http.Cookie{
					Name:     "cloudst",
					Value:    token,
					MaxAge:   3600,
					Path:     "/",
					SameSite: http.SameSiteLaxMode,
					Secure:   false,
					HttpOnly: true,
				})
				c.Redirect(http.StatusMovedPermanently, url)
			}

		}

	} else {
		c.Writer.WriteString("InValid Request:" + err.Error())
	}

}
