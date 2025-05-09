package main

import (
	"github.com/diegoliveiraa/gopportunities/config"
	"github.com/diegoliveiraa/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	// Initialize configs
	logger = config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing config: %v", err)
		return
	}

	router.Initialize()
}
