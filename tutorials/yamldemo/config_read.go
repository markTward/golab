package main

import (
	"fmt"
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

func main() {
	cfg := Config{}

	// Read a yaml config file into struct
	data, _ := ioutil.ReadFile(configFile)

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		log.Fatalln("error", err)
	} else {
		fmt.Printf("Config: %+v\n", &cfg)
	}

}
