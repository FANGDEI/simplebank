package gateway

import (
	"log"

	"github.com/FANGDEI/simplebank/config"
	"github.com/FANGDEI/simplebank/ecrypto"
	"github.com/FANGDEI/simplebank/store/cache"
	"github.com/FANGDEI/simplebank/store/local"
	"github.com/go-playground/validator/v10"
)

type Config struct {
	Cacher    *cache.Manager      `json:"-"`
	Localer   *local.Manager      `json:"-"`
	Cryptoer  *ecrypto.Manager    `json:"-"`
	Validator *validator.Validate `json:"-"`
	Host      string              `json:"host"`
	Port      int                 `json:"port"`
	TokenKey  string              `json:"token_key"`
}

func (c *Config) Key() string {
	return "gateway"
}

var C Config

func init() {
	err := config.ReadConfig(&C)
	if err != nil {
		log.Fatalln(err)
	}
}
