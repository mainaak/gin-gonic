package handler

import (
	"gin-gonic-framework/platform/newsfeed"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewsfeedGet(feed newsfeed.Getter) gin.HandlerFunc{
	return func (ctx *gin.Context) {
		results := feed.GetAll()
		ctx.JSON(http.StatusOK, results)
	}
}
