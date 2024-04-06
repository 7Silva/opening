package main

import (
	"github.com/7Silva/openings/config"
	"github.com/7Silva/openings/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}

	router.Start()
}
