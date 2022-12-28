package controllers

import(
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context){
			c.JSON(200, gin.H{
			"message":     "pong",
			"description": "This is Api Test",
		})
}

func Hoge(c *gin.Context){
			c.JSON(200, gin.H{
			"massage": "Hello Fuga",
		})
}
