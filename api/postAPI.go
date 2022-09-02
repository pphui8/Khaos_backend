package api

import (
	"khaos/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login
func Login(context *gin.Context) {
	// json body
	var body struct {
		Publickey string `json:"publickey"`
	}
	// bind json with the struct
	if err := context.BindJSON(&body); err != nil {
		return
	}
	result, err := database.Login(body.Publickey); if err != nil {
		response := gin.H{
			"error": "user not found",
		}
		context.JSONP(http.StatusOK, response)
		return
	}
	// query from database
	context.JSON(200, result)
}

// Add product
func AddProduct(context *gin.Context) {
	// json body
	var body database.ProductDetail
	// bind json with the struct
	if err := context.BindJSON(&body); err != nil {
		return
	}
	err := database.AddProduct(body); if err != nil {
		response := gin.H{
			"error": "add product failed",
		}
		context.JSONP(http.StatusOK, response)
		return
	}
	// query from database
	context.JSON(200, gin.H{
		"status": "success",
	})
}