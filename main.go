//! Enter pointer of the program

package main

import (
	"khaos/routers"
	"log"
	"strconv"

	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
)

// global variable
var MyConf myConf

type myConf struct {
	AppName string `default:"firstApp"`
	Server  struct {
		IP   string `default:"0.0.0.0"`
		Port int    `default:"7234"`
	} `toml:"server"`
	Database struct {
		Type     string `default:"mysql"`
		User     string `default:"root"`
		Password string `default:"123212321"`
	} `toml:"database"`
}

// start execute
func main() {
	// read conf file
	_, err := toml.DecodeFile("./conf/app.toml", &MyConf)
	if err != nil {
		log.Println("no config file found, use default settings")
	}

	r := gin.Default()

	// register routers
	routers.MountGetRouters(r)
	

	// start
	r.Run(":" + strconv.Itoa(MyConf.Server.Port))
}