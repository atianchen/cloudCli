package test

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (o *OkConfig) Ok(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}

func (o *OkConfig) index(c *gin.Context) {
	c.HTML(http.StatusOK, "cloud.html", "aaa")
}
