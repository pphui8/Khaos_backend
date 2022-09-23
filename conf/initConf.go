// / init conf file
// Path: conf\app.toml
package conf

import (
	"log"

	"github.com/BurntSushi/toml"
)

type myConf struct {
	AppName string `default:"firstApp"`
	Server  struct {
		IP   string `default:"0.0.0.0"`
		Port int    `default:"7234"`
	} `toml:"server"`
	Database struct {
		Database     string `default:"khaos"`
		User     string `default:"root"`
		Password string `default:"123212321"`
		Ip	   string `default:"127.0.0.1"`
		Port     string    `default:"3306"`
	} `toml:"database"`
}

// global variable
var MyConf myConf

func InitConf() {
	// read conf file
	_, err := toml.DecodeFile("./conf/app.toml", &MyConf)
	if err != nil {
		log.Println("no config file found, use default settings")
	}
}