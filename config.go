package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Port         uint
	Listen       string
	Destinations map[string]string
	filename     string
}

func NewConfig(filename string) (*Config, error) {
	config := &Config{filename: filename}
	return config, config.loadFile()
}

func (c *Config) loadFile() error {
	data, err := ioutil.ReadFile(c.filename)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal([]byte(data), &c)
	return err
}
