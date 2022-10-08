package main

import (
	"log"

	"github.com/FANGDEI/simplebank/ecrypto"
	"github.com/FANGDEI/simplebank/gateway"
	"github.com/FANGDEI/simplebank/store/cache"
	"github.com/FANGDEI/simplebank/store/local"
)

func init() {
	log.SetFlags(log.Llongfile)
}

func main() {
	var err error

	gateway.C.Cryptoer = ecrypto.New()

	gateway.C.Localer, err = local.New()
	if err != nil {
		log.Fatalln(err)
	}

	gateway.C.Cacher, err = cache.New()
	if err != nil {
		log.Fatalln(err)
	}

	m := gateway.New()
	err = m.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
