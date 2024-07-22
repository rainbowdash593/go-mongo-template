package main

import (
	"example/template/config"
	"example/template/internal/app"
	logging "example/template/pkg/logger"
	"log"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	logging.ConfigureLogger(logging.Level(cfg.Level))

	// Run
	app.Run(cfg)
}
