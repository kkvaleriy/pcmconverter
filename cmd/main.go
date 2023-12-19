package main

import (
	"converter/internal/app/apiserver"
	"encoding/json"
	"flag"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "../configs/apiserver.json", "path to config json file")
}

func main() {
	flag.Parse()
	log.Printf("Start server with config path %s", configPath)
	config := apiserver.NewConfig()
	configString := config.ReadConfigFile(configPath)
	log.Println(configString)
	if configString != "" {
		err := json.Unmarshal([]byte(configString), config)
		if err != nil {
			log.Fatal(err)
		}
	}
	swFS := apiserver.NewSeaweedFS(config)
	log.Printf("Conf %v", config)
	server := apiserver.New(config, swFS)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}

}
