package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Test(context *gin.Context) {
	response := gin.H{
		"status": "200",
	}
	context.JSONP(http.StatusOK, response)
}
