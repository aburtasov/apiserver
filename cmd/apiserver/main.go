package main

import (
	"encoding/json"
	"flag"
	"log"

	"github.com/aburtasov/apiserver/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "config/apiserver.json", "path to config file")
}

func main() {

	flag.Parse()

	config := apiserver.NewConfig()

	err := json.Unmarshal([]byte(configPath), &config)
	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
