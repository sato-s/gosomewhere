package main

import (
	"github.com/fsnotify/fsnotify"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Destinations map[string]string
type Config struct {
	Port         uint
	Listen       string
	Destinations Destinations
	Basicauth    struct {
		Username string
		Password string
	}
	filename string
	watcher  *fsnotify.Watcher
}

func NewConfig(filename string) (*Config, error) {
	config := &Config{filename: filename}

	if err := config.loadFile(); err != nil {
		return nil, err
	}
	if err := config.run(); err != nil {
		return nil, err
	}
	return config, nil
}

func (c *Config) IsBasicauthEnabled() bool {
	return c.Basicauth.Username != "" && c.Basicauth.Password != ""
}

func (c *Config) loadFile() error {
	data, err := ioutil.ReadFile(c.filename)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal([]byte(data), &c)
	return err
}

func (c *Config) run() error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	c.watcher = watcher
	if err := c.watcher.Add(c.filename); err != nil {
		return err
	}
	log.Printf("Watching %s", c.filename)
	go c.autoload()
	return nil
}

func (c *Config) autoload() {
	defer c.watcher.Close()
	log.Println("monitoring..")
	for {
		select {
		case event := <-c.watcher.Events:
			if event.Op&fsnotify.Write == fsnotify.Write {
				log.Println("modified file:", event.Name)
				err := c.loadFile()
				if err != nil {
					log.Printf("Error: $s", err)
				} else {
					log.Printf("Refreshed setting from %s", c.filename)
				}
			}
		case err := <-c.watcher.Errors:
			log.Printf("Error: $s", err)
		}
	}
}
