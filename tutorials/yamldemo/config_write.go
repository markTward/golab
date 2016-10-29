package main

import (
	"fmt"
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

func main() {
	cfg := Config{}

	// Write a yaml config file from struct
	cfg.Config.JobQueue.Number = 20
	cfg.Config.WorkerQueue.Number = 100
	cfg.Config.Worker.Number = 100
	cfg.Config.StatChan.Number = 100
	cfg.Config.JobQueue.Number = 1000
	cfg.Config.Testing.Duration = 15

	data, _ := yaml.Marshal(&cfg)
	fmt.Println(string(data))
	err := ioutil.WriteFile(configFile, data, 0644)
	if err != nil {
		log.Fatalln("error", err)
	}

}
