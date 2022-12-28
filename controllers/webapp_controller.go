package controllers

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"

	// "1shikawa.com/m/model"
)

// ルートならリダイレクト
func Root(ctx *gin.Context) {
	fmt.Println("User request is redirected")
	ctx.Redirect(302, "/index")
}

func Index(ctx *gin.Context) {
	fmt.Println("User request succeed!")
	ctx.HTML(http.StatusOK, "index.html", nil)
}

func Result(ctx *gin.Context) {
	result, _ := ctx.GetPostForm("result")
	fmt.Printf("User input word is:%s\n", result)
	ctx.HTML(http.StatusOK, "result.html", result)
}


// func ShowAllArticle(c *gin.Context) {
// 	articles := model.GetAllArticle()
// 	users := model.GetAllUser()
// 	c.HTML(http.StatusOK, "index.html", gin.H{"articles": articles, "users": users})
// }
