package ecrypto

import (
	"log"

	"github.com/FANGDEI/simplebank/config"
)

type Config struct {
	Secret string `json:"secret"`
}

func (c *Config) Key() string {
	return "ecrypto"
}

var C Config

func init() {
	err := config.ReadConfig(&C)
	if err != nil {
		log.Fatalln(err)
	}
}
