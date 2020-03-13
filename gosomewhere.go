package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var configFile string
	var err error

	flag.StringVar(&configFile, "config", "", "config yaml")
	flag.Parse()
	fmt.Printf("%s|\n", configFile)

	if configFile == "" {
		// Config file isn't provided, check $HOME/.config/gosomewhere/config.yaml
		user, err := user.Current()
		check(err)
		configFile = filepath.Join(user.HomeDir, ".config/gosomewhere/config.yaml")
		_, err = os.Stat(configFile)
		if err != nil {
			log.Fatal("You didn't provide config.yaml and we couldn't find " + configFile)
		}
	} else {
		// Config file is provided, use it
		configFile, err = filepath.Abs(configFile)
		if err != nil {
			log.Fatal("We couldn't find " + configFile)
		}
	}

	log.Println("using " + configFile)
	var config *Config
	config, err = NewConfig(configFile)
	check(err)

	_, err = NewServer(config)
	if err != nil {
		log.Println(err)
		msg := "We couldn't run server. If this is port permission problem," +
			" run this probram with sudo privilege."
		log.Fatal(msg)
	}
}
