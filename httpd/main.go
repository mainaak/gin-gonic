package main

import (
	"gin-gonic-framework/httpd/handler"
	"gin-gonic-framework/platform/newsfeed"
	"github.com/gin-gonic/gin"
)
//go mod init -> initiates go modules
//go get -u github.com/gin-gonic/gin -> gets the dependency and adds it to the src bin and pkg accordingly where the go path points
func main() {
	//gin.SetMode(gin.ReleaseMode)
	news := newsfeed.New()

	r := gin.Default()

	r.GET("/ping", handler.PingGet())
	r.GET("/newsfeed", handler.NewsfeedGet(news))
	r.POST("/newsfeed", handler.NewsfeedPost(news))

	r.Run(":8081")
}
