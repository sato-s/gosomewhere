package main

import (
	"log"
)

func main() {

	config, err := NewConfig("sample_config.yaml")
	log.Println("running server....")
	if err != nil {
		panic(err)
	}

	_, err = NewServer(config)
	if err != nil {
		panic(err)
	}

}
