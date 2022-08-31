package routers

import (
	"khaos/api"

	"github.com/gin-gonic/gin"
)

func MountPostRouters(r *gin.Engine) {
	// login
	r.POST("/login", api.Login)
}
