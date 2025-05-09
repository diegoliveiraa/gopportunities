package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	// Initialize the Gin router
	router := gin.Default()

	//Initialize the routes
	initializeRoutes(router)

	// Run the server on port 8080
	router.Run(":8080")
}
