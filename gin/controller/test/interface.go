package test

import "github.com/gin-gonic/gin"

type OkInter interface {
	Ok(c *gin.Context)
}
