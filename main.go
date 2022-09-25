// ! Enter pointer of the program
package main

import (
	"io"
	"khaos/conf"
	"khaos/database"
	"khaos/middlewares"
	"khaos/routers"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// start execute
func main() {
	// init conf
	conf.InitConf()

	r := gin.Default()

	// get log file path
	logPathYear := time.Now().Year()
	logPathMonth := time.Now().Month()
	logPathDay := time.Now().Day()

	// create log file
	f, _ := os.Create("./log/" + strconv.Itoa(logPathYear) + strconv.Itoa(int(logPathMonth)) + strconv.Itoa(logPathDay) + ".log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// set log writter
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	// set log writter in production mode
	// gin.DefaultWriter = io.MultiWriter(f)

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