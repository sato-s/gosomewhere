package main

func main() {

	config, err := NewConfig("sample_config.yaml")
	if err != nil {
		panic(err)
	}

	_, err = NewServer(config)
	if err != nil {
		panic(err)
	}

}
