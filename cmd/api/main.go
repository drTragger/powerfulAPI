package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/drTragger/powerfulAPI/internal/app/api"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "path", "configs/api.toml", "Path to config file in .toml format")
}

func main() {
	flag.Parse()
	log.Println("It works!")
	config := api.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Println("Could not find configs file. Using default values:", err)
	}
	server := api.New(config)
	// API server start
	log.Fatal(server.Start())
}
