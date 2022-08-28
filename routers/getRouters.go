package routers

import (
	"khaos/api"

	"github.com/gin-gonic/gin"
)

func MountGetRouters(r *gin.Engine) {
	r.GET("/", api.Test)
	
}
