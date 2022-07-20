package security

import (
	"cloudCli/gin/webConst"
	"cloudCli/repository"
	"cloudCli/utils/encrypt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var userRepository = repository.SysUserRepository{}

/**
 * JWT拦截器
 * @author jensen.chen
 * @date 2022/7/7
 */
func JwtAuthInterceptor() gin.HandlerFunc {
	return func(context *gin.Context) {
		auth := context.Request.Header.Get("x-access-token")
		if len(auth) < 1 {
			auth, _ = context.Cookie("cloudst")
		}
		var realm Realm
		err := encrypt.ParseToken(auth, &realm)
		if err != nil {
			context.Abort()
			context.Status(http.StatusForbidden)
		} else {
			user, err := userRepository.GetByPrimary(realm.Id)
			if err != nil {
				context.Abort()
				context.Status(http.StatusForbidden)
			} else {
				context.Set(webConst.KEY_LOGINUSER, user)
			}
		}
		context.Next()
	}
}
