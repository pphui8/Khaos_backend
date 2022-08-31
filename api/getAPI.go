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
func GetUserNumber(context *gin.Context) {
	result := database.GetUsersNumber()
	response := gin.H{
		"usernumber": result,
	}
	context.JSONP(http.StatusOK, response)
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

/// /productlist (GET)
/// get the product list
func GetProductList(context *gin.Context) {
	var result [53]database.ListProduct
	var len int
	database.GetProductList(&result, &len)
	context.JSON(200, result[:len])
}