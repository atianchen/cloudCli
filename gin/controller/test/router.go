package test

import "github.com/gin-gonic/gin"

func Routers(g *gin.RouterGroup) {

	upload := g.Group("/test")
	{
		upload.GET("/ok", NewOkConfig().Ok)
	}

}
