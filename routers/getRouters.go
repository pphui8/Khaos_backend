package routers

import (
	"khaos/api"

	"github.com/gin-gonic/gin"
)

// test
func MountGetRouters(r *gin.Engine) {
	// test
	r.GET("/", api.Test)
	// get number of users
	r.GET("/usernumber", api.GetUserNumber)
	// get users list

	// get user detail
	r.GET("/user/:id", api.GetUserDetail)
}
