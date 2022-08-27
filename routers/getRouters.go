package routers

import (
	"firstApp/api"

	"github.com/gin-gonic/gin"
)

func MountGetRouters(r *gin.Engine) {
	r.GET("/", api.Test)
	
}
