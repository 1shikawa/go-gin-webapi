package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	// "github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"1shikawa.com/m/controllers"
	"1shikawa.com/m/router"
)

func TestHello(t *testing.T) {
	want := "Hello Golang"
	got := controllers.Hello()
	assert.Equal(t, want, got)
}

func TestGoodbye(t *testing.T) {
	want := "Good Bye Golang"
	got := controllers.Goodbye()
	assert.Equal(t, want, got)
}

func TestPingRouter(t *testing.T) {
	router := router.GetRouter()
	w := httptest.NewRecorder()
	//c, _ := gin.CreateTestContext(w)
	req, _ := http.NewRequest("GET", "/api/v1/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	assert.EqualValues(t, "{\"description\":\"This is Api Test\",\"message\":\"pong\"}", w.Body.String())
}

func TestHogeRouter(t *testing.T) {
	router := router.GetRouter()
	w := httptest.NewRecorder()
	//c, _ := gin.CreateTestContext(w)
	req, _ := http.NewRequest("GET", "/api/v1/hoge", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	assert.EqualValues(t, "{\"massage\":\"Hello Fuga\"}", w.Body.String())
}

func TestRootRouter(t *testing.T) {
	router := router.GetRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 302, w.Code)
}

// func TestIndexRouter(t *testing.T) {
// 	router := router.GetRouter()
	//レスポンス ここに返ってくる
	// w := httptest.NewRecorder()
	// コンテキストを生成
	// c, _ := gin.CreateTestContext(w)
	// リクエストを格納
	// c.Request, _ =http.NewRequest("GET", "/index", nil)
	// req, _ := http.NewRequest("GET", "/index", nil)
	// テストのコンテキストを持って実行
  // router.ServeHTTP(w, req)
  // router.ServeHTTP(w, c.Request)
	// controllers.Index(c)

	// assert.Equal(t, 200, w.Code)
// }

// func TestResultRouter(t *testing.T) {
// 	router := router.GetRouter()
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("POST", "/result", nil)
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, 200, w.Code)
// }
