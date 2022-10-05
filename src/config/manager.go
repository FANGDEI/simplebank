package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Manager struct {
	path  string
	store map[string]any
}

func New(path string) *Manager {
	return &Manager{
		path:  path,
		store: make(map[string]any),
	}
}

func (m *Manager) Load() error {
	buf, err := ioutil.ReadFile(m.path)
	if err != nil {
		return err
	}
	return json.Unmarshal(buf, &m.store)
}

func (m *Manager) ReadConfig(c Config) error {
	log.Println("[CONFIG] READING", c.Key())
	obj, ok := m.store[c.Key()]
	if !ok {
		return fmt.Errorf("%s is not found in config", c.Key())
	}
	buf, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	return json.Unmarshal(buf, c)
}
