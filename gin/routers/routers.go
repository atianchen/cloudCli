package routers

import (
	"cloudCli/utils/log"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

type Option func(*gin.RouterGroup)

var options = []Option{}

// 设置gin的格式，生产环境设置为gin.ReleaseMode
const mode = ""

// 注册app的路由配置
func Include(opts ...Option) {
	options = append(options, opts...)
}

// 初始化
func Init() *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}
	r := gin.Default()
	//r.LoadHTMLGlob("views/*.html")
	r.Use(Cors(), log.GinLogger(), log.GinRecovery(true))
	//r.StaticFS("/static", http.Dir("./views/static"))
	group := r.Group("/api")

	for _, opt := range options {
		opt(group)
	}
	return r
}

func Cors() gin.HandlerFunc {
	return cors.New(
		cors.Config{
			AllowAllOrigins:  true,
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"*"},
			ExposeHeaders:    []string{"Content-Length", "text/plain", "Authorization", "Content-Type"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		},
	)
}

/**
 * 如果用到JWT，可以开启
 */
//func JWTAuthMiddleware(c *gin.Context) {
//	if c.Request.RequestURI == "/api/login/account" {
//		c.Next()
//		return
//	}
//	// 从请求头中取出
//	signToken := c.Request.Header.Get("x-access-token")
//	if signToken == "" {
//		c.JSON(http.StatusOK, gin.H{
//			"code":    1002,
//			"data":    "",
//			"message": "token为空",
//		})
//		c.Abort()
//		return
//	}
//	// 校验token
//	myclaims, err := common.ParserToken(signToken)
//	if err != nil {
//		log.Error(err.Error())
//		c.JSON(http.StatusOK, gin.H{
//			"code":    1003,
//			"data":    err.Error(),
//			"message": "token校验失败",
//		})
//		c.Abort()
//		return
//	}
//	// 将用户的id放在到请求的上下文c上
//	c.Set("id", myclaims.Id)
//	c.Next() // 后续的处理函数可以用过c.Get("userid")来获取当前请求的id
//
//}
