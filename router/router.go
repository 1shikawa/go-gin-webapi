package router

import (
	"html/template"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"1shikawa.com/m/controllers"
)

func GetRouter() *gin.Engine {

	router := gin.Default()
	// ここからCorsの設定
	router.Use(cors.New(cors.Config{
		// アクセスを許可したいアクセス元
		AllowOrigins: []string{
			"http://127.0.0.1",
		},
		// アクセスを許可したいHTTPメソッド(以下の例だとPUTやDELETEはアクセスできません)
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
		},
		// 許可したいHTTPリクエストヘッダ
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: false,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	}))

	router.SetFuncMap(template.FuncMap{
		"nl2br": nl2br,
	})

	api := router.Group("api/v1")
		{
			api.GET("/ping", controllers.Ping)
		}
		{
			api.GET("/hoge", controllers.Hoge)
		}

	router.GET("/", controllers.Root)
	router.GET("/index", controllers.Index)
	router.POST("/result", controllers.Result)
	// router.GET("/articleindex", controllers.ShowAllArticle)

	return router
}

func nl2br(text string) template.HTML {
	return template.HTML(strings.Replace(template.HTMLEscapeString(text), "\n", "<br />", -1))
}
