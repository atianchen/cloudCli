package test

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (o *OkConfig) Ok(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}
