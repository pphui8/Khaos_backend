package api

import (
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

/// /Login (POST)
/// login for manager account
func Login(context *gin.Context) {
	// get post body

	// destructer values
	
	response := gin.H{
		"status": "200",
	}
	context.JSONP(http.StatusOK, response)
}

/// /users (GET)
/// get the user list
func GetUsers(context *gin.Context) {
	// get post body

	// destructer values
	
	response := gin.H{
		"status": "200",
	}
	context.JSONP(http.StatusOK, response)
}

/// /users/:id (GET)
/// get the detail message of a user
func GetUserDetail(context *gin.Context) {
	// get post body

	// destructer values
	
	response := gin.H{
		"status": "200",
	}
	context.JSONP(http.StatusOK, response)
}