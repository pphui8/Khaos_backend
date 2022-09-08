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
	r.GET("/summary", api.GetSummary)
	// get users list
	r.GET("/userslist", api.GetUsersList)
	// get user detail
	r.GET("/user/:id", api.GetUserDetail)
	// delete user
	r.GET("/deluser/:id", api.DelUser)
	// get product list
	r.GET("/productlist", api.GetProductList)
	// delete product
	r.GET("/delproduct/:id", api.DelProduct)
	// get order list
	r.GET("/orderlist", api.GetOrderList)
	// delete order
	r.GET("/delorder/:id", api.DelOrder)
	// get announcement list
	r.GET("/announcementlist", api.GetAnnouncementList)
	// delete a announcement
	r.GET("/delannouncement/:id", api.DelAnnouncement)
}
