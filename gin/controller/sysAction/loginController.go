package sysAction

import (
	"cloudCli/gin/dto"
	"cloudCli/gin/security"
	"cloudCli/repository"
	"cloudCli/utils/encrypt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"time"
)

type LoginController struct {
	repository *repository.SysUserRepository
}

func (lc *LoginController) Init() {

	lc.repository = &repository.SysUserRepository{}
}

/**
 * 用户登录
 * @author jensen.chen
 * @date 2022/7/7
 */
func (lc *LoginController) Login(c *gin.Context) {
	var param dto.LoginDto
	c.BindJSON(&param)
	user, err := lc.repository.FindByCode(param.Name)
	if err != nil {
		c.JSON(http.StatusOK, dto.BuildErrorMsg(err.Error()))
	} else {
		if user.Pwd == encrypt.MD5(param.Pwd) {
			claims := security.Realm{
				user.Id,
				user.Code,
				jwt.StandardClaims{
					ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
					Issuer:    "admin",
				},
			}
			token, _ := encrypt.GenerateToken(claims)
			c.JSON(http.StatusOK, dto.BuildSuccessMsg(token))
		} else {
			c.JSON(http.StatusOK, dto.BuildEmptySuccessMsg())
		}
	}

}
