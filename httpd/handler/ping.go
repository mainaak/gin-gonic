package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func PingGet() gin.HandlerFunc{
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, map[string]interface{}{
			"timestamp": time.Now(),
		})
	}
}
