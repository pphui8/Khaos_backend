package api

import (
	"khaos/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

/// /test (GET)
/// test if backend executing successfully
func Test(context *gin.Context) {
	response := gin.H{
		"status": "200",
	}
	context.JSONP(http.StatusOK, response)
}

/// /usersnumber (GET)
/// get the number of user
func GetSummary(context *gin.Context) {
	result := database.GetSummary()
	context.JSONP(http.StatusOK, result)
}

/// /usersslist (GET)
/// get the user list
func GetUsersList(context *gin.Context) {
	var result [53]database.ListUserData
	var len int = 0
	database.GetUsersList(&result, &len)
	context.JSON(200, result[:len])
}

/// /user/:id (GET)
/// get the user detail
func GetUserDetail(context *gin.Context) {
	result, err := database.GetUsersDetail(context.Param("id")); if err != nil {
		response := gin.H{
			"error": "user not found",
		}
		context.JSONP(http.StatusOK, response)
		return
	}
	context.JSON(200, result)
}

/// /deluser/:id (GET)
/// delete a user
func DelUser(context *gin.Context) {
	err := database.DelUser(context.Param("id")); if err != nil {
		response := gin.H{
			"error": "user not found",
		}
		context.JSONP(http.StatusOK, response)
		return
	}
	response := gin.H{
		"status": "succeess",
	}
	context.JSONP(http.StatusOK, response)
}

/// /productlist (GET)
/// get the product list
func GetProductList(context *gin.Context) {
	var result [53]database.ListProduct
	var len int
	database.GetProductList(&result, &len)
	context.JSON(200, result[:len])
}

/// /delproduct/:id (GET)
/// delete a product
func DelProduct(context *gin.Context) {
	err := database.DelProduct(context.Param("id")); if err != nil {
		response := gin.H{
			"error": "product not found",
		}
		context.JSONP(http.StatusOK, response)
		return
	}
	response := gin.H{
		"status": "succeess",
	}
	context.JSONP(http.StatusOK, response)
}

/// /order (GET)
/// get the order list
func GetOrderList(context *gin.Context) {
	var result [53]database.ListOrder
	var len int
	database.GetOrderList(&result, &len)
	context.JSON(200, result[:len])
}

/// /delorder/:id (GET)
/// delete an order
func DelOrder(context *gin.Context) {
	err := database.DelOrder(context.Param("id")); if err != nil {
		response := gin.H{
			"error": "order not found",
		}
		context.JSONP(http.StatusOK, response)
		return
	}
	response := gin.H{
		"status": "succeess",
	}
	context.JSONP(http.StatusOK, response)
}

/// /announcementlist (GET)
/// get the announcement list
func GetAnnouncementList(context *gin.Context) {
	var result [53]database.ListAnnouncement
	var len int
	database.GetAnnouncementList(&result, &len)
	context.JSON(200, result[:len])
}

/// /delannouncement/:id (GET)
/// delete an announcement
func DelAnnouncement(context *gin.Context) {
	err := database.DelAnnouncement(context.Param("id")); if err != nil {
		response := gin.H{
			"error": "announcement not found",
		}
		context.JSONP(http.StatusOK, response)
		return
	}
	response := gin.H{
		"status": "succeess",
	}
	context.JSONP(http.StatusOK, response)
}

/// /postlist (GET)
/// get the post list
func GetPostList(context *gin.Context) {
	var result [53]database.ListPost
	var len int
	database.GetPostsList(&result, &len)
	context.JSON(200, result[:len])
}

/// /commentlist/:id (GET)
func GetCommentListByPostId(context *gin.Context) {
	var result [53]database.ListComment
	var len int
	database.GetCommentListByPostId(&result, &len, context.Param("id"))
	context.JSON(200, result[:len])
}