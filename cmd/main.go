package main

import (
	"MOA/config"
	"MOA/service"
)

func main() {
	cfg := config.MustLoad()

	logger := service.SetupLogger(cfg.Environment)
}
