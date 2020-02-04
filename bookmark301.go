package main

import (
	"os"
	"os/user"
	"log"
	"path/filepath"
)

func check(err error){
	if err != nil{
		log.Fatal(err)
	}
}

func main() {
	var configFile string
	var err error

	switch len(os.Args) {
		case 1:
			// Config file isn't provided, check $HOME/.config/bookmark301/config.yaml
			user, err := user.Current()
			check(err)
			configFile = filepath.Join(user.HomeDir, ".config/bookmark301/config.yaml")
		case 2:
			// Config file is provided, use it
			configFile, err = filepath.Abs(os.Args[1])
			check(err)
		case 3:
			log.Fatal("Invalid argument")
	}

	log.Println("using " + configFile)
	var config *Config
	config, err = NewConfig(configFile)
	check(err)

	_, err = NewServer(config)
	check(err)
}
