package main

import (
	"github/elliot9/ginExample/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	config.Load()
}

func main() {
	gin.SetMode(config.AppSetting.Env)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
