package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Port         uint
	Listen       string
	Destinations map[string]string
}

func NewConfig(filename string) (*Config, error) {
	data, err := ioutil.ReadFile("sample_config.yaml")
	if err != nil {
		return nil, err
	}
	config := Config{}
	err = yaml.Unmarshal([]byte(data), &config)
	return &config, err
}
