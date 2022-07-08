package controller

import "github.com/gin-gonic/gin"

/**
 *
 * @author jensen.chen
 * @date 2022/7/7
 */
type WebAction interface {
	InitAction()                  //初始化
	AddRouter(g *gin.RouterGroup) //添加路由
}
