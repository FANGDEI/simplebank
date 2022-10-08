package cache

import (
	"log"

	"github.com/FANGDEI/simplebank/config"
)

type Config struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
}

func (c *Config) Key() string {
	return "cache"
}

var C Config

func init() {
	err := config.ReadConfig(&C)
	if err != nil {
		log.Fatalln(err)
	}
}
