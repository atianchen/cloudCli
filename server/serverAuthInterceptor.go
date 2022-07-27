package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * 服务器通信的拦截器
 * @author jensen.chen
 * @date 2022/7/7
 */
func ServerAuthInterceptor() gin.HandlerFunc {
	return func(context *gin.Context) {
		var param NodePayload
		if err := context.BindJSON(&param); err != nil {
			context.Abort()
			context.Status(http.StatusBadRequest)
		} else {
			content, err := Decrypt(param.Content)
			if err == nil {
				context.Set("payload", content)
				context.Next()
			} else {
				context.Abort()
				context.Status(http.StatusBadRequest)
			}
		}
	}
}
