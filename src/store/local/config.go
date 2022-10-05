package local

import (
	"log"

	"github.com/FANGDEI/simplebank/config"
)

type Config struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	DB       string `json:"db"`
	Password string `json:"password"`
	TimeZone string `json:"time_zone"`
}

func (c *Config) Key() string {
	return "local"
}

var C Config

func init() {
	err := config.ReadConfig(&C)
	if err != nil {
		log.Fatalln(err)
	}
}
