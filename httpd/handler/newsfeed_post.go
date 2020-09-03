package handler

import (
	"gin-gonic-framework/platform/newsfeed"
	"github.com/gin-gonic/gin"
	"net/http"
)

type newsfeedRequest struct {
	Title string `json:"title"`
	Post string `json:"post"`
}

func NewsfeedPost(feed  newsfeed.Adder/* *newsfeed.Repo*/) gin.HandlerFunc {
	return func(context *gin.Context) {
		requestBody := newsfeedRequest{}
		context.Bind(&requestBody)

		feed.Add(newsfeed.Item{
			Title: requestBody.Title,
			Post:  requestBody.Post,
		})

		context.JSON(http.StatusOK, &requestBody)
	}
}
