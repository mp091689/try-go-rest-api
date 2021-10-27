package main

import (
	"flag"
	"log"

	"github.com/MykytaPopov/try-go-rest-api/internal/app/api"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config")
}

func main() {
	flag.Parse()

	config := api.NewConfig(configPath)

	a := api.NewApi(config)
	if err := a.Start(); err != nil {
		log.Fatal(err)
	}
}
