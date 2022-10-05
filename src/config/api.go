package config

import "log"

var handler *Manager

func init() {
	err := Init("../data/simplebank.json")
	if err != nil {
		log.Fatalln(err)
	}
}

func Init(path string) error {
	handler = New(path)
	return handler.Load()
}

func ReadConfig(c Config) error {
	return handler.ReadConfig(c)
}
