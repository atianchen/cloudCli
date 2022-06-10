package test

import (
	"cloudCli/channel"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (o *OkConfig) Ok(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}

func (o *OkConfig) index(c *gin.Context) {
	ct := channel.GetChan("PluginTask")
	ct <- channel.Message{Payload: c.Get("userName")}
	select {
	case _ = <-ct:
		c.HTML(http.StatusOK, "cloud.html", "Success")
	case <-time.After(5 * time.Second):
		c.HTML(http.StatusOK, "cloud.html", "5秒超时")
	}

}
