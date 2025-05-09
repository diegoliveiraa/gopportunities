// @title Gopportunities API
// @version 1.0
// @description API para gerenciamento de vagas
// @termsOfService http://swagger.io/terms/

// @contact.name Diego Oliveira
// @contact.url https://github.com/diegoliveiraa
// @contact.email seuemail@exemplo.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1

package main

import (
	"github.com/diegoliveiraa/gopportunities/config"
	_ "github.com/diegoliveiraa/gopportunities/docs"
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
