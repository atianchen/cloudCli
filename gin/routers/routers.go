package routers

import (
	"cloudCli/utils/log"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr/v2"
	"os"
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
	getwd, _ := os.Getwd()
	//r.LoadHTMLGlob(getwd + "/gin/views/*.html")
	r.Use(Cors(), log.GinLogger(), log.GinRecovery(true))
	staticPack := packr.New("cloudCli", getwd+"/gin/ui")
	r.StaticFS("/cloud/ui", staticPack) // http.Dir(getwd+"/gin/views/ui"))
	group := r.Group("/cloud")

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
