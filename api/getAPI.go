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
func GetUserList(context *gin.Context) {
	response := gin.H{
		"status": "200",
	}
	context.JSONP(http.StatusOK, response)
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