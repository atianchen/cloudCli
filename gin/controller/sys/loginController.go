package sys

import (
	"cloudCli/gin/dto"
	"cloudCli/gin/security"
	"cloudCli/repository"
	"cloudCli/utils/encrypt"
	go_beanutils "github.com/atianchen/go-beanutils"
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

func (lc *LoginController) CurrentUser(c *gin.Context) {
	auth, err := c.Cookie("cloudst")
	if err != nil {
		c.JSON(http.StatusOK, dto.BuildErrorMsg("Not Login"))
		return
	}
	var realm security.Realm
	if err = encrypt.ParseToken(auth, &realm); err == nil {
		if u, err := lc.repository.GetByPrimary(realm.Id); err == nil {
			var userDto UserDto
			go_beanutils.CopyProperties(&userDto, u)
			c.JSON(http.StatusOK, dto.BuildSuccessMsg(&userDto))
		} else {
			c.JSON(http.StatusOK, dto.BuildErrorMsg("Not Login"))
		}
	} else {
		c.JSON(http.StatusOK, dto.BuildErrorMsg("Not Login"))
	}
}

/**
 * 用户登录
 * @author jensen.chen
 * @date 2022/7/7
 */
func (lc *LoginController) Login(c *gin.Context) {
	var param LoginDto
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
			c.SetCookie("cloudst", token, 3600, "/", "/cloudCli", true, true)
			var userDto UserDto
			go_beanutils.CopyProperties(&userDto, user)
			c.JSON(http.StatusOK, dto.BuildSuccessMsg(&userDto))
		} else {
			c.JSON(http.StatusOK, dto.BuildEmptySuccessMsg())
		}
	}

}
