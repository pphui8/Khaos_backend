// ! Enter pointer of the program
package main

import (
	"khaos/conf"
	"khaos/database"
	"khaos/middlewares"
	"khaos/routers"
	"strconv"

	"github.com/gin-gonic/gin"
)

// start execute
func main() {
	// init conf
	conf.InitConf()

	r := gin.Default()

	// set middlewares
	r.Use(middlewares.Cors())

	// init database
	database.InitDB()

	// register routers
	routers.MountGetRouters(r)
	routers.MountPostRouters(r)

	// start
	r.Run(":" + strconv.Itoa(conf.MyConf.Server.Port))
}